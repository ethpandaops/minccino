package coordinator

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/ethpandaops/assertoor/pkg/coordinator/test"
	"github.com/ethpandaops/assertoor/pkg/coordinator/types"
	"github.com/gorhill/cronexpr"
)

type TestRunner struct {
	coordinator types.Coordinator

	runIDCounter       uint64
	lastExecutedRunID  uint64
	testSchedulerMutex sync.Mutex

	testRunMap           map[uint64]types.Test
	testQueue            []types.TestRunner
	testRegistryMutex    sync.RWMutex
	testNotificationChan chan bool
}

func NewTestRunner(coordinator types.Coordinator, lastRunID uint64) *TestRunner {
	return &TestRunner{
		coordinator:       coordinator,
		runIDCounter:      lastRunID,
		lastExecutedRunID: lastRunID,

		testRunMap:           map[uint64]types.Test{},
		testQueue:            []types.TestRunner{},
		testNotificationChan: make(chan bool, 1),
	}
}

func (c *TestRunner) GetTestByRunID(runID uint64) types.Test {
	c.testRegistryMutex.RLock()
	defer c.testRegistryMutex.RUnlock()

	return c.testRunMap[runID]
}

func (c *TestRunner) GetTestQueue() []types.Test {
	c.testRegistryMutex.RLock()
	defer c.testRegistryMutex.RUnlock()

	tests := make([]types.Test, len(c.testQueue))
	for idx, test := range c.testQueue {
		tests[idx] = test
	}

	return tests
}

func (c *TestRunner) ScheduleTest(descriptor types.TestDescriptor, configOverrides map[string]any, allowDuplicate bool) (types.TestRunner, error) {
	if descriptor.Err() != nil {
		return nil, fmt.Errorf("cannot create test from failed test descriptor: %w", descriptor.Err())
	}

	testRef, err := c.createTestRun(descriptor, configOverrides, allowDuplicate)
	if err != nil {
		return nil, err
	}

	select {
	case c.testNotificationChan <- true:
	default:
	}

	return testRef, nil
}

func (c *TestRunner) createTestRun(descriptor types.TestDescriptor, configOverrides map[string]any, allowDuplicate bool) (types.TestRunner, error) {
	c.testSchedulerMutex.Lock()
	defer c.testSchedulerMutex.Unlock()

	if !allowDuplicate {
		for _, queuedTest := range c.GetTestQueue() {
			if queuedTest.TestID() == descriptor.ID() {
				return nil, fmt.Errorf("test already in queue")
			}
		}
	}

	c.runIDCounter++
	runID := c.runIDCounter

	testRef, err := test.CreateTest(runID, descriptor, c.coordinator.Logger().WithField("module", "test"), c.coordinator, configOverrides)
	if err != nil {
		return nil, fmt.Errorf("failed initializing test run #%v '%v': %w", runID, descriptor.Config().Name, err)
	}

	c.testRegistryMutex.Lock()
	c.testQueue = append(c.testQueue, testRef)
	c.testRunMap[runID] = testRef
	c.testRegistryMutex.Unlock()

	return testRef, nil
}

func (c *TestRunner) RunTestExecutionLoop(ctx context.Context, concurrencyLimit uint64) {
	if concurrencyLimit < 1 {
		concurrencyLimit = 1
	}

	semaphore := make(chan bool, concurrencyLimit)
	waitGroup := sync.WaitGroup{}

runLoop:
	for {
		var nextTest types.TestRunner

		c.testRegistryMutex.Lock()
		if len(c.testQueue) > 0 {
			nextTest = c.testQueue[0]
			c.testQueue = c.testQueue[1:]
		}
		c.testRegistryMutex.Unlock()

		if nextTest != nil {
			// run next test
			testFunc := func(nextTest types.TestRunner) {
				defer func() {
					<-semaphore
					waitGroup.Done()
				}()

				c.runTest(ctx, nextTest)
			}
			semaphore <- true
			if ctx.Err() != nil {
				break runLoop
			}

			waitGroup.Add(1)

			go testFunc(nextTest)
		} else {
			// sleep and wait for queue notification
			select {
			case <-ctx.Done():
				break runLoop
			case <-c.testNotificationChan:
			case <-time.After(60 * time.Second):
			}
		}
	}

	waitGroup.Wait()
}

func (c *TestRunner) runTest(ctx context.Context, testRef types.TestRunner) {
	c.lastExecutedRunID = testRef.RunID()

	if err := testRef.Validate(); err != nil {
		testRef.Logger().Errorf("test validation failed: %v", err)
		return
	}

	if err := testRef.Run(ctx); err != nil {
		testRef.Logger().Errorf("test execution failed: %v", err)
	}
}

func (c *TestRunner) RunTestScheduler(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.coordinator.Logger().WithError(err.(error)).Panicf("uncaught panic in TestRunner.RunTestScheduler: %v, stack: %v", err, string(debug.Stack()))
		}
	}()

	// startup scheduler
	for _, testDescr := range c.getStartupTests() {
		_, err := c.ScheduleTest(testDescr, nil, false)
		if err != nil {
			c.coordinator.Logger().Errorf("could not schedule startup test execution for %v (%v): %v", testDescr.ID(), testDescr.Config().Name, err)
		}
	}

	// cron scheduler
	cronTime := time.Unix((time.Now().Unix()/60)*60, 0)

	for {
		cronTime = cronTime.Add(1 * time.Minute)
		cronTimeDiff := time.Since(cronTime)

		if cronTimeDiff < 0 {
			select {
			case <-ctx.Done():
				return
			case <-time.After(cronTimeDiff.Abs()):
			}
		}

		for _, testDescr := range c.getCronTests(cronTime) {
			_, err := c.ScheduleTest(testDescr, nil, false)
			if err != nil {
				c.coordinator.Logger().Errorf("could not schedule cron test execution for %v (%v): %v", testDescr.ID(), testDescr.Config().Name, err)
			}
		}
	}
}

func (c *TestRunner) getStartupTests() []types.TestDescriptor {
	descriptors := []types.TestDescriptor{}

	for _, testDescr := range c.coordinator.TestRegistry().GetTestDescriptors() {
		if testDescr.Err() != nil {
			continue
		}

		testConfig := testDescr.Config()
		if testConfig.Schedule == nil || testConfig.Schedule.Startup {
			descriptors = append(descriptors, testDescr)
		}
	}

	return descriptors
}

func (c *TestRunner) getCronTests(cronTime time.Time) []types.TestDescriptor {
	descriptors := []types.TestDescriptor{}

	for _, testDescr := range c.coordinator.TestRegistry().GetTestDescriptors() {
		if testDescr.Err() != nil {
			continue
		}

		testConfig := testDescr.Config()
		if testConfig.Schedule == nil || len(testConfig.Schedule.Cron) == 0 {
			continue
		}

		triggerTest := false

		for _, cronExprStr := range testConfig.Schedule.Cron {
			cronExpr, err := cronexpr.Parse(cronExprStr)
			if err != nil {
				c.coordinator.Logger().Errorf("invalid cron expression for test %v (%v): %v", testDescr.ID(), testConfig.Name, err)
				break
			}

			next := cronExpr.Next(cronTime.Add(-1 * time.Second))
			if next.Compare(cronTime) == 0 {
				triggerTest = true
				break
			}
		}

		if !triggerTest {
			continue
		}

		descriptors = append(descriptors, testDescr)
	}

	return descriptors
}

func (c *TestRunner) RunTestCleanup(ctx context.Context, retentionTime time.Duration) {
	defer func() {
		if err := recover(); err != nil {
			c.coordinator.Logger().WithError(err.(error)).Panicf("uncaught panic in TestRunner.runTestCleanup: %v, stack: %v", err, string(debug.Stack()))
		}
	}()

	if retentionTime <= 0 {
		retentionTime = 14 * 24 * time.Hour
	}

	cleanupInterval := 1 * time.Hour
	if retentionTime <= 4*time.Hour {
		cleanupInterval = 10 * time.Minute
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(cleanupInterval):
		}

		c.cleanupTestHistory(retentionTime)
	}
}

func (c *TestRunner) cleanupTestHistory(_ time.Duration) {
	// TODO: clean db
}
