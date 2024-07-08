// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/test/{testId}": {
            "get": {
                "description": "Returns the test definition with given ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Get test definition by test ID",
                "operationId": "getTest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the test definition to get details for",
                        "name": "testId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.GetTestResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_run/{runId}": {
            "get": {
                "description": "Returns the run details with given ID. Includes a summary and a list of task with limited details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Get test run by run ID",
                "operationId": "getTestRun",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the test run to get details for",
                        "name": "runId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.GetTestRunResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_run/{runId}/cancel": {
            "post": {
                "description": "Returns the test/run id \u0026 status of the cancelled test.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Cancel test run by test ID",
                "operationId": "postTestRunCancel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the test run to cancel",
                        "name": "runId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Test cancellation options",
                        "name": "cancelOptions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PostTestRunCancelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.PostTestRunCancelResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_run/{runId}/details": {
            "get": {
                "description": "Returns the run details with given ID. Includes a summary and a list of task with all details (incl. logs \u0026 task configurations)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Get detailed test run by run ID",
                "operationId": "getTestRunDetails",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the test run to get details for",
                        "name": "runId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.GetTestRunDetailsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_run/{runId}/status": {
            "get": {
                "description": "Returns the run status with given ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Get test run status by run ID",
                "operationId": "getTestRunStatus",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the test run to get the status for",
                        "name": "runId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_runs": {
            "get": {
                "description": "Returns a list of all test runs.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Get list of test runs",
                "operationId": "getTestRuns",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Return test runs for this test ID only",
                        "name": "test_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/api.GetTestRunsResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/test_runs/schedule": {
            "post": {
                "description": "Returns the test \u0026 run id of the scheduled test execution.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TestRun"
                ],
                "summary": "Schedule new test run by test ID",
                "operationId": "postTestRunsSchedule",
                "parameters": [
                    {
                        "description": "Rest run options",
                        "name": "runOptions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PostTestRunsScheduleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.PostTestRunsScheduleResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tests": {
            "get": {
                "description": "Returns the list of test definitions. These test definitions can be used to create new test runs and are supplied via the assertoor configuration.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Get list of test definitions",
                "operationId": "getTests",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/api.GetTestsResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tests/register": {
            "post": {
                "description": "Returns the test id and name of the added test.",
                "consumes": [
                    "application/json",
                    "application/yaml"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Register new test via yaml configuration",
                "operationId": "postTestsRegister",
                "parameters": [
                    {
                        "description": "Test configuration (json or yaml)",
                        "name": "testConfig",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PostTestsRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.PostTestsRegisterResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tests/register_external": {
            "post": {
                "description": "Returns the test id and name of the added test.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Register new test via external test configuration",
                "operationId": "postTestsRegisterExternal",
                "parameters": [
                    {
                        "description": "Test configuration (json or yaml)",
                        "name": "externalTestConfig",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PostTestsRegisterExternalRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api.PostTestsRegisterExternalResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.GetTestResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "configVars": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "schedule": {
                    "$ref": "#/definitions/types.TestSchedule"
                },
                "source": {
                    "type": "string"
                },
                "timeout": {
                    "type": "integer"
                }
            }
        },
        "api.GetTestRunDetailedTask": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "config_yaml": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "log": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.GetTestRunDetailedTaskLog"
                    }
                },
                "name": {
                    "type": "string"
                },
                "parent_index": {
                    "type": "integer"
                },
                "result": {
                    "type": "string"
                },
                "result_error": {
                    "type": "string"
                },
                "result_yaml": {
                    "type": "string"
                },
                "runtime": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "started": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string"
                },
                "stop_time": {
                    "type": "integer"
                },
                "timeout": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.GetTestRunDetailedTaskLog": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "datalen": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "api.GetTestRunDetailsResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "run_id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/types.TestStatus"
                },
                "stop_time": {
                    "type": "integer"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.GetTestRunDetailedTask"
                    }
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.GetTestRunResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "run_id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/types.TestStatus"
                },
                "stop_time": {
                    "type": "integer"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.GetTestRunTask"
                    }
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.GetTestRunTask": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "index": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parent_index": {
                    "type": "integer"
                },
                "result": {
                    "type": "string"
                },
                "result_error": {
                    "type": "string"
                },
                "runtime": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "started": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string"
                },
                "stop_time": {
                    "type": "integer"
                },
                "timeout": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.GetTestRunsResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "run_id": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/types.TestStatus"
                },
                "stop_time": {
                    "type": "integer"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.GetTestsResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "api.PostTestRunCancelRequest": {
            "type": "object",
            "properties": {
                "skip_cleanup": {
                    "type": "boolean"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.PostTestRunCancelResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "run_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.PostTestRunsScheduleRequest": {
            "type": "object",
            "properties": {
                "allow_duplicate": {
                    "type": "boolean"
                },
                "config": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.PostTestRunsScheduleResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "name": {
                    "type": "string"
                },
                "run_id": {
                    "type": "integer"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.PostTestsRegisterExternalRequest": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": true
                },
                "configVars": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "file": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "schedule": {
                    "$ref": "#/definitions/types.TestSchedule"
                },
                "timeout": {
                    "type": "integer"
                }
            }
        },
        "api.PostTestsRegisterExternalResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "name": {
                    "type": "string"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "api.PostTestsRegisterRequest": {
            "type": "object",
            "properties": {
                "cleanupTasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/helper.RawMessage"
                    }
                },
                "config": {
                    "type": "object",
                    "additionalProperties": true
                },
                "configVars": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "schedule": {
                    "$ref": "#/definitions/types.TestSchedule"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/helper.RawMessage"
                    }
                },
                "timeout": {
                    "type": "string"
                }
            }
        },
        "api.PostTestsRegisterResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "name": {
                    "type": "string"
                },
                "test_id": {
                    "type": "string"
                }
            }
        },
        "github_com_ethpandaops_assertoor_pkg_coordinator_web_api.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "helper.RawMessage": {
            "type": "object"
        },
        "types.TestSchedule": {
            "type": "object",
            "properties": {
                "cron": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "startup": {
                    "type": "boolean"
                }
            }
        },
        "types.TestStatus": {
            "type": "string",
            "enum": [
                "pending",
                "running",
                "success",
                "failure",
                "skipped",
                "aborted"
            ],
            "x-enum-varnames": [
                "TestStatusPending",
                "TestStatusRunning",
                "TestStatusSuccess",
                "TestStatusFailure",
                "TestStatusSkipped",
                "TestStatusAborted"
            ]
        }
    },
    "tags": [
        {
            "description": "All endpoints related to test definitions",
            "name": "Test"
        },
        {
            "description": "All endpoints related to test runs",
            "name": "TestRun"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Assertoor API Documentation",
	Description:      "API for querying information about Assertoor tests",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
