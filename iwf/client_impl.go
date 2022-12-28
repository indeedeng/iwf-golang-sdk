package iwf

import (
	"context"
	"fmt"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
)

type clientImpl struct {
	UnregisteredClient
	registry Registry
	options  *ClientOptions
}

func (c *clientImpl) StartWorkflow(ctx context.Context, workflow Workflow, startStateId, workflowId string, timeoutSecs int32, input interface{}, options *WorkflowOptions) (string, error) {
	wfType := GetDefaultWorkflowType(workflow)
	stateDef := c.registry.getWorkflowStateDef(wfType, startStateId)
	if !stateDef.CanStartWorkflow {
		return "", NewWorkflowDefinitionFmtError("cannot start workflow %v with start state %v", wfType, startStateId)
	}
	if options != nil {
		for _, sa := range options.InitialSearchAttributes {
			typeMap := c.registry.getSearchAttributeTypeStore(wfType)
			registeredType, ok := typeMap[sa.GetKey()]
			if !ok || registeredType != sa.GetValueType() {
				return "", fmt.Errorf("key %s is not defined as search attribute value type %s", sa.GetKey(), registeredType)
			}
			v, _ := getSearchAttributeValue(sa)
			if v == nil {
				return "", fmt.Errorf("search attribute value is not set correctly for key %s with value type %s", sa.GetKey(), sa.GetValueType())
			}
		}
	}
	return c.UnregisteredClient.StartWorkflow(ctx, wfType, startStateId, workflowId, timeoutSecs, input, options)
}

func (c *clientImpl) SignalWorkflow(ctx context.Context, workflow Workflow, workflowId, workflowRunId, signalChannelName string, signalValue interface{}) error {
	wfType := GetDefaultWorkflowType(workflow)
	signalNameStore := c.registry.getWorkflowSignalNameStore(wfType)
	if !signalNameStore[signalChannelName] {
		return NewWorkflowDefinitionFmtError("signal channel %v is not defined in workflow type %v", signalChannelName, wfType)
	}
	return c.UnregisteredClient.SignalWorkflow(ctx, workflowId, workflowRunId, signalChannelName, signalValue)
}

func (c *clientImpl) GetWorkflowDataObjects(ctx context.Context, workflow Workflow, workflowId, workflowRunId string, keys []string) (map[string]Object, error) {
	wfType := GetDefaultWorkflowType(workflow)
	doTypeMap := c.registry.getWorkflowDataObjectKeyStore(wfType)
	for _, k := range keys {
		_, ok := doTypeMap[k]
		if !ok {
			return nil, fmt.Errorf("data object type %v is not registered", k)
		}
	}
	return c.UnregisteredClient.GetWorkflowDataObjects(ctx, workflowId, workflowRunId, keys)
}

func (c *clientImpl) GetWorkflowSearchAttributes(ctx context.Context, workflow Workflow, workflowId, workflowRunId string, keys []string) (map[string]interface{}, error) {
	wfType := GetDefaultWorkflowType(workflow)
	allTypes := c.registry.getSearchAttributeTypeStore(wfType)
	var keyAndTypes []iwfidl.SearchAttributeKeyAndType
	for _, k := range keys {
		keyAndTypes = append(keyAndTypes, iwfidl.SearchAttributeKeyAndType{
			Key:       &k,
			ValueType: ptr.Any(allTypes[k]),
		})
	}
	vals, err := c.UnregisteredClient.GetWorkflowSearchAttributes(ctx, workflowId, workflowRunId, keyAndTypes)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{}, len(vals))
	for _, val := range vals {
		v, err := getSearchAttributeValue(val)
		if err != nil {
			return nil, err
		}
		out[val.GetKey()] = v
	}
	return out, nil
}

func (c *clientImpl) GetAllWorkflowSearchAttributes(ctx context.Context, workflow Workflow, workflowId, workflowRunId string) (map[string]interface{}, error) {
	wfType := GetDefaultWorkflowType(workflow)
	allTypes := c.registry.getSearchAttributeTypeStore(wfType)
	var keys []string
	for k := range allTypes {
		keys = append(keys, k)
	}
	return c.GetWorkflowSearchAttributes(ctx, workflow, workflowId, workflowRunId, keys)
}

