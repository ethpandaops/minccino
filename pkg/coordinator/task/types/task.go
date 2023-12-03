package types

import (
	"context"
	"time"

	"github.com/ethpandaops/minccino/pkg/coordinator/helper"
	"github.com/ethpandaops/minccino/pkg/coordinator/human-duration"
	"github.com/sirupsen/logrus"
)

type TaskDescriptor struct {
	Name        string
	Description string
	Config      interface{}
	NewTask     func(ctx *TaskContext, options *TaskOptions) (Task, error)
}

type TaskOptions struct {
	// The name of the task to run.
	Name string `yaml:"name" json:"name"`
	// The configuration object of the task.
	Config *helper.RawMessage `yaml:"config" json:"config"`
	// The title of the task - this is used to describe the task to the user.
	Title string `yaml:"title" json:"title"`
	// Timeout defines the max time waiting for the condition to be met.
	Timeout human.Duration `yaml:"timeout" json:"timeout"`
}

type TaskResult uint8

const (
	TaskResultNone    TaskResult = 0
	TaskResultSuccess TaskResult = 1
	TaskResultFailure TaskResult = 2
)

type Task interface {
	Name() string
	Title() string
	Description() string

	Config() interface{}
	Logger() logrus.FieldLogger
	Timeout() time.Duration

	ValidateConfig() error
	Execute(ctx context.Context) error
}
