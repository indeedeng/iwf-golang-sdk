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

// checks if the WorkflowStateOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateOptions{}

// WorkflowStateOptions struct for WorkflowStateOptions
type WorkflowStateOptions struct {
	SearchAttributesLoadingPolicy *PersistenceLoadingPolicy `json:"searchAttributesLoadingPolicy,omitempty"`
	DataAttributesLoadingPolicy *PersistenceLoadingPolicy `json:"dataAttributesLoadingPolicy,omitempty"`
	CommandCarryOverPolicy *CommandCarryOverPolicy `json:"commandCarryOverPolicy,omitempty"`
	WaitUntilApiTimeoutSeconds *int32 `json:"waitUntilApiTimeoutSeconds,omitempty"`
	ExecuteApiTimeoutSeconds *int32 `json:"executeApiTimeoutSeconds,omitempty"`
	WaitUntilApiRetryPolicy *RetryPolicy `json:"waitUntilApiRetryPolicy,omitempty"`
	ExecuteApiRetryPolicy *RetryPolicy `json:"executeApiRetryPolicy,omitempty"`
	WaitUntilApiFailurePolicy *WaitUntilApiFailurePolicy `json:"waitUntilApiFailurePolicy,omitempty"`
	SkipWaitUntil *bool `json:"skipWaitUntil,omitempty"`
}

// NewWorkflowStateOptions instantiates a new WorkflowStateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateOptions() *WorkflowStateOptions {
	this := WorkflowStateOptions{}
	return &this
}

// NewWorkflowStateOptionsWithDefaults instantiates a new WorkflowStateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateOptionsWithDefaults() *WorkflowStateOptions {
	this := WorkflowStateOptions{}
	return &this
}

// GetSearchAttributesLoadingPolicy returns the SearchAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || IsNil(o.SearchAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.SearchAttributesLoadingPolicy
}

// GetSearchAttributesLoadingPolicyOk returns a tuple with the SearchAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || IsNil(o.SearchAttributesLoadingPolicy) {
		return nil, false
	}
	return o.SearchAttributesLoadingPolicy, true
}

// HasSearchAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasSearchAttributesLoadingPolicy() bool {
	if o != nil && !IsNil(o.SearchAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetSearchAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the SearchAttributesLoadingPolicy field.
func (o *WorkflowStateOptions) SetSearchAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.SearchAttributesLoadingPolicy = &v
}

// GetDataAttributesLoadingPolicy returns the DataAttributesLoadingPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetDataAttributesLoadingPolicy() PersistenceLoadingPolicy {
	if o == nil || IsNil(o.DataAttributesLoadingPolicy) {
		var ret PersistenceLoadingPolicy
		return ret
	}
	return *o.DataAttributesLoadingPolicy
}

// GetDataAttributesLoadingPolicyOk returns a tuple with the DataAttributesLoadingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetDataAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool) {
	if o == nil || IsNil(o.DataAttributesLoadingPolicy) {
		return nil, false
	}
	return o.DataAttributesLoadingPolicy, true
}

// HasDataAttributesLoadingPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasDataAttributesLoadingPolicy() bool {
	if o != nil && !IsNil(o.DataAttributesLoadingPolicy) {
		return true
	}

	return false
}

// SetDataAttributesLoadingPolicy gets a reference to the given PersistenceLoadingPolicy and assigns it to the DataAttributesLoadingPolicy field.
func (o *WorkflowStateOptions) SetDataAttributesLoadingPolicy(v PersistenceLoadingPolicy) {
	o.DataAttributesLoadingPolicy = &v
}

// GetCommandCarryOverPolicy returns the CommandCarryOverPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetCommandCarryOverPolicy() CommandCarryOverPolicy {
	if o == nil || IsNil(o.CommandCarryOverPolicy) {
		var ret CommandCarryOverPolicy
		return ret
	}
	return *o.CommandCarryOverPolicy
}

// GetCommandCarryOverPolicyOk returns a tuple with the CommandCarryOverPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetCommandCarryOverPolicyOk() (*CommandCarryOverPolicy, bool) {
	if o == nil || IsNil(o.CommandCarryOverPolicy) {
		return nil, false
	}
	return o.CommandCarryOverPolicy, true
}

// HasCommandCarryOverPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasCommandCarryOverPolicy() bool {
	if o != nil && !IsNil(o.CommandCarryOverPolicy) {
		return true
	}

	return false
}

// SetCommandCarryOverPolicy gets a reference to the given CommandCarryOverPolicy and assigns it to the CommandCarryOverPolicy field.
func (o *WorkflowStateOptions) SetCommandCarryOverPolicy(v CommandCarryOverPolicy) {
	o.CommandCarryOverPolicy = &v
}

// GetWaitUntilApiTimeoutSeconds returns the WaitUntilApiTimeoutSeconds field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetWaitUntilApiTimeoutSeconds() int32 {
	if o == nil || IsNil(o.WaitUntilApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.WaitUntilApiTimeoutSeconds
}

// GetWaitUntilApiTimeoutSecondsOk returns a tuple with the WaitUntilApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetWaitUntilApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitUntilApiTimeoutSeconds) {
		return nil, false
	}
	return o.WaitUntilApiTimeoutSeconds, true
}

// HasWaitUntilApiTimeoutSeconds returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasWaitUntilApiTimeoutSeconds() bool {
	if o != nil && !IsNil(o.WaitUntilApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetWaitUntilApiTimeoutSeconds gets a reference to the given int32 and assigns it to the WaitUntilApiTimeoutSeconds field.
func (o *WorkflowStateOptions) SetWaitUntilApiTimeoutSeconds(v int32) {
	o.WaitUntilApiTimeoutSeconds = &v
}

// GetExecuteApiTimeoutSeconds returns the ExecuteApiTimeoutSeconds field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetExecuteApiTimeoutSeconds() int32 {
	if o == nil || IsNil(o.ExecuteApiTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.ExecuteApiTimeoutSeconds
}

// GetExecuteApiTimeoutSecondsOk returns a tuple with the ExecuteApiTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetExecuteApiTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.ExecuteApiTimeoutSeconds) {
		return nil, false
	}
	return o.ExecuteApiTimeoutSeconds, true
}

// HasExecuteApiTimeoutSeconds returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasExecuteApiTimeoutSeconds() bool {
	if o != nil && !IsNil(o.ExecuteApiTimeoutSeconds) {
		return true
	}

	return false
}

// SetExecuteApiTimeoutSeconds gets a reference to the given int32 and assigns it to the ExecuteApiTimeoutSeconds field.
func (o *WorkflowStateOptions) SetExecuteApiTimeoutSeconds(v int32) {
	o.ExecuteApiTimeoutSeconds = &v
}

// GetWaitUntilApiRetryPolicy returns the WaitUntilApiRetryPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetWaitUntilApiRetryPolicy() RetryPolicy {
	if o == nil || IsNil(o.WaitUntilApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.WaitUntilApiRetryPolicy
}

// GetWaitUntilApiRetryPolicyOk returns a tuple with the WaitUntilApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetWaitUntilApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || IsNil(o.WaitUntilApiRetryPolicy) {
		return nil, false
	}
	return o.WaitUntilApiRetryPolicy, true
}

// HasWaitUntilApiRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasWaitUntilApiRetryPolicy() bool {
	if o != nil && !IsNil(o.WaitUntilApiRetryPolicy) {
		return true
	}

	return false
}

// SetWaitUntilApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the WaitUntilApiRetryPolicy field.
func (o *WorkflowStateOptions) SetWaitUntilApiRetryPolicy(v RetryPolicy) {
	o.WaitUntilApiRetryPolicy = &v
}

// GetExecuteApiRetryPolicy returns the ExecuteApiRetryPolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetExecuteApiRetryPolicy() RetryPolicy {
	if o == nil || IsNil(o.ExecuteApiRetryPolicy) {
		var ret RetryPolicy
		return ret
	}
	return *o.ExecuteApiRetryPolicy
}

// GetExecuteApiRetryPolicyOk returns a tuple with the ExecuteApiRetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetExecuteApiRetryPolicyOk() (*RetryPolicy, bool) {
	if o == nil || IsNil(o.ExecuteApiRetryPolicy) {
		return nil, false
	}
	return o.ExecuteApiRetryPolicy, true
}

// HasExecuteApiRetryPolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasExecuteApiRetryPolicy() bool {
	if o != nil && !IsNil(o.ExecuteApiRetryPolicy) {
		return true
	}

	return false
}

// SetExecuteApiRetryPolicy gets a reference to the given RetryPolicy and assigns it to the ExecuteApiRetryPolicy field.
func (o *WorkflowStateOptions) SetExecuteApiRetryPolicy(v RetryPolicy) {
	o.ExecuteApiRetryPolicy = &v
}

// GetWaitUntilApiFailurePolicy returns the WaitUntilApiFailurePolicy field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetWaitUntilApiFailurePolicy() WaitUntilApiFailurePolicy {
	if o == nil || IsNil(o.WaitUntilApiFailurePolicy) {
		var ret WaitUntilApiFailurePolicy
		return ret
	}
	return *o.WaitUntilApiFailurePolicy
}

// GetWaitUntilApiFailurePolicyOk returns a tuple with the WaitUntilApiFailurePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetWaitUntilApiFailurePolicyOk() (*WaitUntilApiFailurePolicy, bool) {
	if o == nil || IsNil(o.WaitUntilApiFailurePolicy) {
		return nil, false
	}
	return o.WaitUntilApiFailurePolicy, true
}

// HasWaitUntilApiFailurePolicy returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasWaitUntilApiFailurePolicy() bool {
	if o != nil && !IsNil(o.WaitUntilApiFailurePolicy) {
		return true
	}

	return false
}

// SetWaitUntilApiFailurePolicy gets a reference to the given WaitUntilApiFailurePolicy and assigns it to the WaitUntilApiFailurePolicy field.
func (o *WorkflowStateOptions) SetWaitUntilApiFailurePolicy(v WaitUntilApiFailurePolicy) {
	o.WaitUntilApiFailurePolicy = &v
}

// GetSkipWaitUntil returns the SkipWaitUntil field value if set, zero value otherwise.
func (o *WorkflowStateOptions) GetSkipWaitUntil() bool {
	if o == nil || IsNil(o.SkipWaitUntil) {
		var ret bool
		return ret
	}
	return *o.SkipWaitUntil
}

// GetSkipWaitUntilOk returns a tuple with the SkipWaitUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateOptions) GetSkipWaitUntilOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipWaitUntil) {
		return nil, false
	}
	return o.SkipWaitUntil, true
}

// HasSkipWaitUntil returns a boolean if a field has been set.
func (o *WorkflowStateOptions) HasSkipWaitUntil() bool {
	if o != nil && !IsNil(o.SkipWaitUntil) {
		return true
	}

	return false
}

// SetSkipWaitUntil gets a reference to the given bool and assigns it to the SkipWaitUntil field.
func (o *WorkflowStateOptions) SetSkipWaitUntil(v bool) {
	o.SkipWaitUntil = &v
}

func (o WorkflowStateOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchAttributesLoadingPolicy) {
		toSerialize["searchAttributesLoadingPolicy"] = o.SearchAttributesLoadingPolicy
	}
	if !IsNil(o.DataAttributesLoadingPolicy) {
		toSerialize["dataAttributesLoadingPolicy"] = o.DataAttributesLoadingPolicy
	}
	if !IsNil(o.CommandCarryOverPolicy) {
		toSerialize["commandCarryOverPolicy"] = o.CommandCarryOverPolicy
	}
	if !IsNil(o.WaitUntilApiTimeoutSeconds) {
		toSerialize["waitUntilApiTimeoutSeconds"] = o.WaitUntilApiTimeoutSeconds
	}
	if !IsNil(o.ExecuteApiTimeoutSeconds) {
		toSerialize["executeApiTimeoutSeconds"] = o.ExecuteApiTimeoutSeconds
	}
	if !IsNil(o.WaitUntilApiRetryPolicy) {
		toSerialize["waitUntilApiRetryPolicy"] = o.WaitUntilApiRetryPolicy
	}
	if !IsNil(o.ExecuteApiRetryPolicy) {
		toSerialize["executeApiRetryPolicy"] = o.ExecuteApiRetryPolicy
	}
	if !IsNil(o.WaitUntilApiFailurePolicy) {
		toSerialize["waitUntilApiFailurePolicy"] = o.WaitUntilApiFailurePolicy
	}
	if !IsNil(o.SkipWaitUntil) {
		toSerialize["skipWaitUntil"] = o.SkipWaitUntil
	}
	return toSerialize, nil
}

type NullableWorkflowStateOptions struct {
	value *WorkflowStateOptions
	isSet bool
}

func (v NullableWorkflowStateOptions) Get() *WorkflowStateOptions {
	return v.value
}

func (v *NullableWorkflowStateOptions) Set(val *WorkflowStateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateOptions(val *WorkflowStateOptions) *NullableWorkflowStateOptions {
	return &NullableWorkflowStateOptions{value: val, isSet: true}
}

func (v NullableWorkflowStateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


