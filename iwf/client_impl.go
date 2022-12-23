package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
	"net/http"
)

type clientImpl struct {
	registry  Registry
	options   *ClientOptions
	apiClient *iwfidl.APIClient
}

func (c *clientImpl) StartWorkflow(ctx context.Context, workflow interface{}, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error) {
	wfType, err := getWorkflowType(workflow)
	if err != nil {
		return "", err
	}
	if c.registry != nil {
		stateDef := c.registry.getWorkflowStateDef(wfType, startStateId)
		if !stateDef.CanStartWorkflow {
			return "", NewWorkflowDefinitionFmtError("cannot start workflow %v with start state %v", wfType, startStateId)
		}
	}

	var encodedInput *iwfidl.EncodedObject
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

func getWorkflowType(workflow interface{}) (string, error) {
	var wfType string
	wfType, ok := workflow.(string)
	if ok {
		return wfType, nil
	}

	wf, ok := workflow.(Workflow)
	if !ok {
		return "", NewWorkflowDefinitionError("workflow parameter must be either iwf.Workflow instance, or a string format of workflow type")
	}
	return GetDefaultWorkflowType(wf), nil

}

func (c *clientImpl) StopWorkflow(ctx context.Context, workflowId, workflowRunId string, options *WorkflowStopOptions) error {
	reqPost := c.apiClient.DefaultApi.ApiV1WorkflowStopPost(ctx)
	req := &iwfidl.WorkflowStopRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
	}
	if options != nil {
		req.StopType = &options.StopType
	}
	httpResp, err := reqPost.WorkflowStopRequest(*req).Execute()
	return processError(err, httpResp)
}

func (c *clientImpl) GetSimpleWorkflowResult(ctx context.Context, workflowId, workflowRunId string, resultPtr interface{}) error {
	req := c.apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(ctx)
	resp, httpResp, err := req.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
		NeedsResults:  ptr.Any(true),
	}).Execute()
	if err := processError(err, httpResp); err != nil {
		return err
	}
	if len(resp.Results) != 1 {
		return NewWorkflowDefinitionError("this workflow should have one or zero state output for using this API")
	}
	output := resp.Results[0].CompletedStateOutput
	return c.options.ObjectEncoder.Decode(output, resultPtr)
}

func (c *clientImpl) GetComplexWorkflowResults(ctx context.Context, workflowId, workflowRunId string) ([]iwfidl.StateCompletionOutput, error) {
	req := c.apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(ctx)
	resp, httpResp, err := req.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
		NeedsResults:  ptr.Any(true),
	}).Execute()
	if err := processError(err, httpResp); err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func (c *clientImpl) SignalWorkflow(ctx context.Context, workflow interface{}, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error {
	if c.registry != nil {
		wfType, err := getWorkflowType(workflow)
		if err != nil {
			return err
		}
		signalNameStore := c.registry.getWorkflowSignalNameStore(wfType)
		if !signalNameStore[signalChannelName] {
			return NewWorkflowDefinitionFmtError("signal channel %v is not defined in workflow type %v", signalChannelName, wfType)
		}
	}
	value, err := c.options.ObjectEncoder.Encode(signalValue)
	if err != nil {
		return err
	}
	req := c.apiClient.DefaultApi.ApiV1WorkflowSignalPost(ctx)
	httpResp, err := req.WorkflowSignalRequest(iwfidl.WorkflowSignalRequest{
		WorkflowId:        workflowId,
		WorkflowRunId:     &workflowRunId,
		SignalChannelName: signalChannelName,
		SignalValue:       value,
	}).Execute()
	return processError(err, httpResp)
}

func (c *clientImpl) ResetWorkflow(ctx context.Context, workflowId, workflowRunId string, options *ResetWorkflowTypeAndOptions) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) GetWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string, keys []string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) GetAllWorkflowDataObjects(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) GetWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) GetAllWorkflowSearchAttributes(ctx context.Context, workflow interface{}, workflowId, workflowRunId string) (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *clientImpl) SearchWorkflow(ctx context.Context, query string, pageSize int) (*iwfidl.WorkflowSearchResponse, error) {
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
