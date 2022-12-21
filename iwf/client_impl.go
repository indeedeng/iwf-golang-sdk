package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"net/http"
)

type client struct {
	registry  Registry
	options   *ClientOptions
	apiClient *iwfidl.APIClient
}

func (c *client) StartWorkflow(ctx context.Context, workflow interface{}, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error) {
	var wfType string
	wfType, ok := workflow.(string)
	if !ok {
		wf, ok := workflow.(Workflow)
		if !ok {
			return "", NewWorkflowDefinitionError("workflow parameter must be either iwf.Workflow instance, or a string format of workflow type")
		}
		wfType = c.registry.GetWorkflowType(wf)
	}
	var encodedInput *iwfidl.EncodedObject
	var err error
	if input != nil {
		encodedInput, err = c.options.ObjectEncoder.Encode(input)
		if err != nil {
			return "", err
		}
	}

	var stateOptions *iwfidl.WorkflowStateOptions
	var startOptions *iwfidl.WorkflowStartOptions
	if options != nil {
		stateOptions = options.StartStateOptions
		startOptions = &iwfidl.WorkflowStartOptions{
			WorkflowIDReusePolicy: options.WorkflowIdReusePolicy,
			CronSchedule:          options.WorkflowCronSchedule,
			RetryPolicy:           options.WorkflowRetryPolicy,
		}
	}

	req := c.apiClient.DefaultApi.ApiV1WorkflowStartPost(ctx)
	resp, httpResp, err := req.WorkflowStartRequest(iwfidl.WorkflowStartRequest{
		WorkflowId:             workflowId,
		IwfWorkflowType:        wfType,
		WorkflowTimeoutSeconds: timeoutSecs,
		IwfWorkerUrl:           c.options.WorkerUrl,
		StartStateId:           startStateId,
		StateInput:             encodedInput,
		StateOptions:           stateOptions,
		WorkflowStartOptions:   startOptions,
	}).Execute()
	if err := processError(err, httpResp); err != nil {
		return "", err
	}
	return resp.GetWorkflowRunId(), nil
}

func (c *client) StopWorkflow(ctx context.Context, workflowId, workflowRunId string, options *WorkflowStopOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetSimpleWorkflowResult(ctx context.Context, workflowId, workflowRunId string, resultPtr interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetComplexWorkflowResults(ctx context.Context, workflowId, workflowRunId string) ([]iwfidl.StateCompletionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SignalWorkflow(ctx context.Context, workflow interface{}, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *client) ResetWorkflow(ctx context.Context, workflowId, workflowRunId string, options *ResetWorkflowTypeAndOptions) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string, keys []string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetAllWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetAllWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) SearchWorkflow(ctx context.Context, query string, pageSize int) (*iwfidl.WorkflowSearchResponse, error) {
	//TODO implement me
	panic("implement me")
}

func processError(err error, httpResp *http.Response) error {
	if err != nil {
		return err
	}
	if httpResp.StatusCode != http.StatusOK {
		return NewInternalServiceError("HTTP request failed", *httpResp)
	}
	return nil
}
