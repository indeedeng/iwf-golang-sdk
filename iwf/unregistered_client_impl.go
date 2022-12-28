package iwf

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
	"net/http"
)

type unregisteredClientImpl struct {
	options   *ClientOptions
	apiClient *iwfidl.APIClient
}

func (u *unregisteredClientImpl) StartWorkflow(ctx context.Context, workflowType string, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error) {
	var encodedInput *iwfidl.EncodedObject
	var err error
	if input != nil {
		encodedInput, err = u.options.ObjectEncoder.Encode(input)
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

	req := u.apiClient.DefaultApi.ApiV1WorkflowStartPost(ctx)
	resp, httpResp, err := req.WorkflowStartRequest(iwfidl.WorkflowStartRequest{
		WorkflowId:             workflowId,
		IwfWorkflowType:        workflowType,
		WorkflowTimeoutSeconds: timeoutSecs,
		IwfWorkerUrl:           u.options.WorkerUrl,
		StartStateId:           startStateId,
		StateInput:             encodedInput,
		StateOptions:           stateOptions,
		WorkflowStartOptions:   startOptions,
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return "", err
	}
	return resp.GetWorkflowRunId(), nil
}

func (u *unregisteredClientImpl) SignalWorkflow(ctx context.Context, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error {
	value, err := u.options.ObjectEncoder.Encode(signalValue)
	if err != nil {
		return err
	}
	req := u.apiClient.DefaultApi.ApiV1WorkflowSignalPost(ctx)
	httpResp, err := req.WorkflowSignalRequest(iwfidl.WorkflowSignalRequest{
		WorkflowId:        workflowId,
		WorkflowRunId:     &workflowRunId,
		SignalChannelName: signalChannelName,
		SignalValue:       value,
	}).Execute()
	return u.processError(err, httpResp)
}

func (u *unregisteredClientImpl) GetWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string, keys []string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) GetAllWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string) (map[string]iwfidl.EncodedObject, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) GetWorkflowSearchAttributes(ctx context.Context, workflowId, workflowRunId string) (map[string]iwfidl.SearchAttribute, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) GetAllWorkflowSearchAttributes(ctx context.Context, workflowId, workflowRunId string) (map[string]iwfidl.SearchAttribute, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) StopWorkflow(ctx context.Context, workflowId, workflowRunId string, options *WorkflowStopOptions) error {
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowStopPost(ctx)
	req := &iwfidl.WorkflowStopRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
	}
	if options != nil {
		req.StopType = &options.StopType
	}
	httpResp, err := reqPost.WorkflowStopRequest(*req).Execute()
	return u.processError(err, httpResp)
}

func (u *unregisteredClientImpl) GetSimpleWorkflowResult(ctx context.Context, workflowId, workflowRunId string, resultPtr interface{}) error {
	req := u.apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(ctx)
	resp, httpResp, err := req.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
		NeedsResults:  ptr.Any(true),
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return err
	}
	if len(resp.Results) != 1 {
		return NewWorkflowDefinitionError("this workflow should have one or zero state output for using this API")
	}
	output := resp.Results[0].CompletedStateOutput
	return u.options.ObjectEncoder.Decode(output, resultPtr)
}

func (u *unregisteredClientImpl) GetComplexWorkflowResults(ctx context.Context, workflowId, workflowRunId string) ([]iwfidl.StateCompletionOutput, error) {
	req := u.apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(ctx)
	resp, httpResp, err := req.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: &workflowRunId,
		NeedsResults:  ptr.Any(true),
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func (u *unregisteredClientImpl) ResetWorkflow(ctx context.Context, workflowId, workflowRunId string, options *ResetWorkflowTypeAndOptions) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) SearchWorkflow(ctx context.Context, request iwfidl.WorkflowSearchRequest) (*iwfidl.WorkflowSearchResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *unregisteredClientImpl) processError(err error, httpResp *http.Response) error {
	if err != nil {
		return err
	}
	if httpResp.StatusCode != http.StatusOK {
		return NewInternalServiceError("HTTP request failed", *httpResp)
	}
	return nil
}
