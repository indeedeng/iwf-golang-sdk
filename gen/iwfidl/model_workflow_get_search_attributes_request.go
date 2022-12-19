/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// WorkflowGetSearchAttributesRequest struct for WorkflowGetSearchAttributesRequest
type WorkflowGetSearchAttributesRequest struct {
	WorkflowId string `json:"workflowId"`
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
	Keys []SearchAttributeKeyAndType `json:"keys,omitempty"`
}

// NewWorkflowGetSearchAttributesRequest instantiates a new WorkflowGetSearchAttributesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowGetSearchAttributesRequest(workflowId string) *WorkflowGetSearchAttributesRequest {
	this := WorkflowGetSearchAttributesRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowGetSearchAttributesRequestWithDefaults instantiates a new WorkflowGetSearchAttributesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowGetSearchAttributesRequestWithDefaults() *WorkflowGetSearchAttributesRequest {
	this := WorkflowGetSearchAttributesRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowGetSearchAttributesRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetSearchAttributesRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowGetSearchAttributesRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowGetSearchAttributesRequest) GetWorkflowRunId() string {
	if o == nil || o.WorkflowRunId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetSearchAttributesRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || o.WorkflowRunId == nil {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowGetSearchAttributesRequest) HasWorkflowRunId() bool {
	if o != nil && o.WorkflowRunId != nil {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowGetSearchAttributesRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *WorkflowGetSearchAttributesRequest) GetKeys() []SearchAttributeKeyAndType {
	if o == nil || o.Keys == nil {
		var ret []SearchAttributeKeyAndType
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetSearchAttributesRequest) GetKeysOk() ([]SearchAttributeKeyAndType, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *WorkflowGetSearchAttributesRequest) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []SearchAttributeKeyAndType and assigns it to the Keys field.
func (o *WorkflowGetSearchAttributesRequest) SetKeys(v []SearchAttributeKeyAndType) {
	o.Keys = v
}

func (o WorkflowGetSearchAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.WorkflowRunId != nil {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowGetSearchAttributesRequest struct {
	value *WorkflowGetSearchAttributesRequest
	isSet bool
}

func (v NullableWorkflowGetSearchAttributesRequest) Get() *WorkflowGetSearchAttributesRequest {
	return v.value
}

func (v *NullableWorkflowGetSearchAttributesRequest) Set(val *WorkflowGetSearchAttributesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowGetSearchAttributesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowGetSearchAttributesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowGetSearchAttributesRequest(val *WorkflowGetSearchAttributesRequest) *NullableWorkflowGetSearchAttributesRequest {
	return &NullableWorkflowGetSearchAttributesRequest{value: val, isSet: true}
}

func (v NullableWorkflowGetSearchAttributesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowGetSearchAttributesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


