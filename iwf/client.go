package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
)

type Client interface {
	// StartWorkflow starts a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// startStateId is the first stateId to start
	// workflowId is the required identifier for the workflow execution(see Cadence/Temporal for more details about WorkflowId uniqueness)
	// timeoutSecs is required as the workflow execution timeout in seconds
	// input can be optional, it's the input for the startState
	// options is optional includes like IdReusePolicy, RetryPolicy, CronSchedule and also WorkflowStateOptions for the startState. Empty by default(when nil).
	// return the workflowRunId
	StartWorkflow(ctx context.Context, workflow interface{}, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error)
	// StopWorkflow stops a workflow execution.
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// options is optional, default (when nil)to use Cancel as stopType
	StopWorkflow(ctx context.Context, workflowId, workflowRunId string, options *WorkflowStopOptions) error
	// GetSimpleWorkflowResult returns the result of a workflow execution, for simple case that only one WorkflowState completes with result
	// If there are more than one WorkflowStates complete with result, GetComplexWorkflowResults must be used instead
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// resultPtr is the pointer to retrieve the result
	GetSimpleWorkflowResult(ctx context.Context, workflowId, workflowRunId string, resultPtr interface{}) error
	// GetComplexWorkflowResults returns the results of a workflow execution
	// It returns a list of iwfidl.StateCompletionOutput and user code will have to use ObjectEncoder to deserialize
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	GetComplexWorkflowResults(ctx context.Context, workflowId, workflowRunId string) ([]iwfidl.StateCompletionOutput, error)
	// SignalWorkflow signals a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// signalChannelName is required, signalValue is optional(for case of empty value)
	SignalWorkflow(ctx context.Context, workflow interface{}, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error
	// ResetWorkflow resets a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// resetWorkflowTypeAndOptions is optional, it provides combination parameter for reset. Default (when nil) will reset to iwfidl.BEGINNING resetType
	// return the workflowRunId
	ResetWorkflow(ctx context.Context, workflowId, workflowRunId string, options *ResetWorkflowTypeAndOptions) (string, error)
	// DescribeWorkflow describes the basic info of a workflow execution
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error)
	// GetWorkflowDataObjects returns the data objects of a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowDataObjects API instead
	// It returns data objects in format of iwfidl.EncodedObject and user code have to use ObjectEncoder to deserialize
	GetWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string, keys []string) (map[string]iwfidl.EncodedObject, error)
	// GetAllWorkflowDataObjects returns all the data objects of a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// It returns data objects in format of iwfidl.EncodedObject and user code have to use ObjectEncoder to deserialize
	GetAllWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]iwfidl.EncodedObject, error)
	// GetWorkflowSearchAttributes returns search attributes of a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	// keys is required to be non-empty. If you intend to return all data objects, use GetAllWorkflowDataObjects API instead
	GetWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error)
	// GetAllWorkflowSearchAttributes returns all search attributes of a workflow execution
	// workflow must be either an instance of Workflow interface or the workflowType in string format
	// workflowId is required, workflowRunId is optional and default to current runId of the workflowId
	GetAllWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error)
	// SearchWorkflow searches for workflow executions given a query (see SearchAttribute query in Cadence/Temporal)
	SearchWorkflow(ctx context.Context, query string, pageSize int) (*iwfidl.WorkflowSearchResponse, error)
}

// NewUntypedClient returns a UntypedClient
// It will let you invoke the APIs to iWF server without much type validation checks(workflow type, signalChannelName, etc).
// It's useful for :
// 1) calling Client APIs without workflow registry(which may require to have all the workflow dependencies)
// 2) some use cases like dynamic workflow instance that workflowType is not from class simple name.
func NewUntypedClient(options *ClientOptions) Client {
	return doNewClient(nil, options)
}

// NewClient returns a Client
func NewClient(registry Registry, options *ClientOptions) Client {
	if registry == nil {
		panic("cannot have nil registry")
	}
	return doNewClient(registry, options)
}

func doNewClient(registry Registry, options *ClientOptions) Client {
	if options == nil {
		options = GetLocalDefaultClientOptions()
	}

	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: options.ServerUrl,
			},
		},
	})

	return &clientImpl{
		registry:  registry.(*registryImpl),
		options:   options,
		apiClient: apiClient,
	}
}
