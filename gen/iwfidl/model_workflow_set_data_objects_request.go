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

// checks if the WorkflowSetDataObjectsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowSetDataObjectsRequest{}

// WorkflowSetDataObjectsRequest struct for WorkflowSetDataObjectsRequest
type WorkflowSetDataObjectsRequest struct {
	WorkflowId    string     `json:"workflowId"`
	WorkflowRunId *string    `json:"workflowRunId,omitempty"`
	Objects       []KeyValue `json:"objects,omitempty"`
}

// NewWorkflowSetDataObjectsRequest instantiates a new WorkflowSetDataObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSetDataObjectsRequest(workflowId string) *WorkflowSetDataObjectsRequest {
	this := WorkflowSetDataObjectsRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowSetDataObjectsRequestWithDefaults instantiates a new WorkflowSetDataObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSetDataObjectsRequestWithDefaults() *WorkflowSetDataObjectsRequest {
	this := WorkflowSetDataObjectsRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowSetDataObjectsRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSetDataObjectsRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowSetDataObjectsRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowSetDataObjectsRequest) GetWorkflowRunId() string {
	if o == nil || IsNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSetDataObjectsRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowSetDataObjectsRequest) HasWorkflowRunId() bool {
	if o != nil && !IsNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowSetDataObjectsRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *WorkflowSetDataObjectsRequest) GetObjects() []KeyValue {
	if o == nil || IsNil(o.Objects) {
		var ret []KeyValue
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSetDataObjectsRequest) GetObjectsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *WorkflowSetDataObjectsRequest) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []KeyValue and assigns it to the Objects field.
func (o *WorkflowSetDataObjectsRequest) SetObjects(v []KeyValue) {
	o.Objects = v
}

func (o WorkflowSetDataObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowSetDataObjectsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	if !IsNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if !IsNil(o.Objects) {
		toSerialize["objects"] = o.Objects
	}
	return toSerialize, nil
}

type NullableWorkflowSetDataObjectsRequest struct {
	value *WorkflowSetDataObjectsRequest
	isSet bool
}

func (v NullableWorkflowSetDataObjectsRequest) Get() *WorkflowSetDataObjectsRequest {
	return v.value
}

func (v *NullableWorkflowSetDataObjectsRequest) Set(val *WorkflowSetDataObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSetDataObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSetDataObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSetDataObjectsRequest(val *WorkflowSetDataObjectsRequest) *NullableWorkflowSetDataObjectsRequest {
	return &NullableWorkflowSetDataObjectsRequest{value: val, isSet: true}
}

func (v NullableWorkflowSetDataObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSetDataObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
