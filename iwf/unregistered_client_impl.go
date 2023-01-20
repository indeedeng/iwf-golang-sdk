package iwf

import (
	"context"
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"net/http"
)

type unregisteredClientImpl struct {
	options   *ClientOptions
	apiClient *iwfidl.APIClient
}

func (u *unregisteredClientImpl) StartWorkflow(ctx context.Context, workflowType string, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *UnregisteredWorkflowOptions) (string, error) {
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
		for _, sa := range options.InitialSearchAttributes {
			val, _ := getSearchAttributeValue(sa)
			if val == nil {
				return "", fmt.Errorf("search attribute value is not set correctly for key %s with value type %s", sa.GetKey(), sa.GetValueType())
			}
		}
		stateOptions = options.StartStateOptions
		startOptions = &iwfidl.WorkflowStartOptions{
			WorkflowIDReusePolicy: options.WorkflowIdReusePolicy,
			CronSchedule:          options.WorkflowCronSchedule,
			RetryPolicy:           options.WorkflowRetryPolicy,
			SearchAttributes:      options.InitialSearchAttributes,
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

func (u *unregisteredClientImpl) GetWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string, keys []string) (map[string]Object, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("must specify keys to return, use GetAllWorkflowDataObjects if intended to get all keys")
	}
	return u.doGetWorkflowDataObjects(ctx, workflowId, workflowRunId, keys)
}

func (u *unregisteredClientImpl) GetAllWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string) (map[string]Object, error) {
	return u.doGetWorkflowDataObjects(ctx, workflowId, workflowRunId, nil)
}

func (u *unregisteredClientImpl) doGetWorkflowDataObjects(ctx context.Context, workflowId, workflowRunId string, keys []string) (map[string]Object, error) {
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowDataobjectsGetPost(ctx)
	resp, httpResp, err := reqPost.WorkflowGetDataObjectsRequest(iwfidl.WorkflowGetDataObjectsRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: iwfidl.PtrString(workflowRunId),
		Keys:          keys,
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return nil, err
	}
	out := make(map[string]Object, len(resp.Objects))
	for _, kv := range resp.Objects {
		out[kv.GetKey()] = NewObject(ptr.Any(kv.GetValue()), u.options.ObjectEncoder)
	}
	return out, nil
}

func (u *unregisteredClientImpl) GetWorkflowSearchAttributes(ctx context.Context, workflowId, workflowRunId string, keys []iwfidl.SearchAttributeKeyAndType) (map[string]iwfidl.SearchAttribute, error) {
	if len(keys) == 0 {
		return nil, fmt.Errorf("must specify keys to return, use GetAllWorkflowSearchAttributes if intended to get all keys")
	}
	return u.doGetWorkflowSearchAttributes(ctx, workflowId, workflowRunId, keys)
}

func (u *unregisteredClientImpl) doGetWorkflowSearchAttributes(ctx context.Context, workflowId, workflowRunId string, keys []iwfidl.SearchAttributeKeyAndType) (map[string]iwfidl.SearchAttribute, error) {
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowSearchattributesGetPost(ctx)
	resp, httpResp, err := reqPost.WorkflowGetSearchAttributesRequest(iwfidl.WorkflowGetSearchAttributesRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: iwfidl.PtrString(workflowRunId),
		Keys:          keys,
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return nil, err
	}
	out := make(map[string]iwfidl.SearchAttribute, len(resp.SearchAttributes))
	for _, kv := range resp.SearchAttributes {
		out[kv.GetKey()] = kv
	}
	return out, nil
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
	resetType := iwfidl.BEGINNING
	reason := ptr.Any("")
	if options != nil {
		resetType = options.ResetType
		reason = &options.Reason
	}

	req := iwfidl.WorkflowResetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: iwfidl.PtrString(workflowRunId),
		ResetType:     resetType,
		Reason:        reason,
	}
	if options != nil {
		req.HistoryEventId = options.HistoryEventId
		req.HistoryEventTime = ptr.Any(options.HistoryEventTime.Format(DateTimeFormat))
		req.SkipSignalReapply = options.SkipSignalReapply
		req.StateId = options.StateId
		req.StateExecutionId = options.StateExecutionId
	}
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowResetPost(ctx)
	resp, httpResp, err := reqPost.WorkflowResetRequest(req).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return "", err
	}
	return resp.WorkflowRunId, nil
}

func (u *unregisteredClientImpl) DescribeWorkflow(ctx context.Context, workflowId, workflowRunId string) (*WorkflowInfo, error) {
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowGetPost(ctx)
	resp, httpResp, err := reqPost.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId:    workflowId,
		WorkflowRunId: iwfidl.PtrString(workflowRunId),
		NeedsResults:  ptr.Any(false),
	}).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return nil, err
	}
	return &WorkflowInfo{
		Status:       resp.WorkflowStatus,
		CurrentRunId: resp.WorkflowRunId,
	}, nil
}

func (u *unregisteredClientImpl) SearchWorkflow(ctx context.Context, request iwfidl.WorkflowSearchRequest) (*iwfidl.WorkflowSearchResponse, error) {
	reqPost := u.apiClient.DefaultApi.ApiV1WorkflowSearchPost(ctx)
	resp, httpResp, err := reqPost.WorkflowSearchRequest(request).Execute()
	if err := u.processError(err, httpResp); err != nil {
		return nil, err
	}
	return resp, nil
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
