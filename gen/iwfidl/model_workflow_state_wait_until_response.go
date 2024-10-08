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

// checks if the WorkflowStateWaitUntilResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateWaitUntilResponse{}

// WorkflowStateWaitUntilResponse struct for WorkflowStateWaitUntilResponse
type WorkflowStateWaitUntilResponse struct {
	LocalActivityInput         *string                       `json:"localActivityInput,omitempty"`
	UpsertSearchAttributes     []SearchAttribute             `json:"upsertSearchAttributes,omitempty"`
	UpsertDataObjects          []KeyValue                    `json:"upsertDataObjects,omitempty"`
	CommandRequest             *CommandRequest               `json:"commandRequest,omitempty"`
	UpsertStateLocals          []KeyValue                    `json:"upsertStateLocals,omitempty"`
	RecordEvents               []KeyValue                    `json:"recordEvents,omitempty"`
	PublishToInterStateChannel []InterStateChannelPublishing `json:"publishToInterStateChannel,omitempty"`
}

// NewWorkflowStateWaitUntilResponse instantiates a new WorkflowStateWaitUntilResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateWaitUntilResponse() *WorkflowStateWaitUntilResponse {
	this := WorkflowStateWaitUntilResponse{}
	return &this
}

// NewWorkflowStateWaitUntilResponseWithDefaults instantiates a new WorkflowStateWaitUntilResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateWaitUntilResponseWithDefaults() *WorkflowStateWaitUntilResponse {
	this := WorkflowStateWaitUntilResponse{}
	return &this
}

// GetLocalActivityInput returns the LocalActivityInput field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetLocalActivityInput() string {
	if o == nil || IsNil(o.LocalActivityInput) {
		var ret string
		return ret
	}
	return *o.LocalActivityInput
}

// GetLocalActivityInputOk returns a tuple with the LocalActivityInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetLocalActivityInputOk() (*string, bool) {
	if o == nil || IsNil(o.LocalActivityInput) {
		return nil, false
	}
	return o.LocalActivityInput, true
}

// HasLocalActivityInput returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasLocalActivityInput() bool {
	if o != nil && !IsNil(o.LocalActivityInput) {
		return true
	}

	return false
}

// SetLocalActivityInput gets a reference to the given string and assigns it to the LocalActivityInput field.
func (o *WorkflowStateWaitUntilResponse) SetLocalActivityInput(v string) {
	o.LocalActivityInput = &v
}

// GetUpsertSearchAttributes returns the UpsertSearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetUpsertSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.UpsertSearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.UpsertSearchAttributes
}

// GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetUpsertSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.UpsertSearchAttributes) {
		return nil, false
	}
	return o.UpsertSearchAttributes, true
}

// HasUpsertSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasUpsertSearchAttributes() bool {
	if o != nil && !IsNil(o.UpsertSearchAttributes) {
		return true
	}

	return false
}

// SetUpsertSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the UpsertSearchAttributes field.
func (o *WorkflowStateWaitUntilResponse) SetUpsertSearchAttributes(v []SearchAttribute) {
	o.UpsertSearchAttributes = v
}

// GetUpsertDataObjects returns the UpsertDataObjects field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetUpsertDataObjects() []KeyValue {
	if o == nil || IsNil(o.UpsertDataObjects) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertDataObjects
}

// GetUpsertDataObjectsOk returns a tuple with the UpsertDataObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetUpsertDataObjectsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.UpsertDataObjects) {
		return nil, false
	}
	return o.UpsertDataObjects, true
}

// HasUpsertDataObjects returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasUpsertDataObjects() bool {
	if o != nil && !IsNil(o.UpsertDataObjects) {
		return true
	}

	return false
}

// SetUpsertDataObjects gets a reference to the given []KeyValue and assigns it to the UpsertDataObjects field.
func (o *WorkflowStateWaitUntilResponse) SetUpsertDataObjects(v []KeyValue) {
	o.UpsertDataObjects = v
}

// GetCommandRequest returns the CommandRequest field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetCommandRequest() CommandRequest {
	if o == nil || IsNil(o.CommandRequest) {
		var ret CommandRequest
		return ret
	}
	return *o.CommandRequest
}

// GetCommandRequestOk returns a tuple with the CommandRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetCommandRequestOk() (*CommandRequest, bool) {
	if o == nil || IsNil(o.CommandRequest) {
		return nil, false
	}
	return o.CommandRequest, true
}

// HasCommandRequest returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasCommandRequest() bool {
	if o != nil && !IsNil(o.CommandRequest) {
		return true
	}

	return false
}

// SetCommandRequest gets a reference to the given CommandRequest and assigns it to the CommandRequest field.
func (o *WorkflowStateWaitUntilResponse) SetCommandRequest(v CommandRequest) {
	o.CommandRequest = &v
}

// GetUpsertStateLocals returns the UpsertStateLocals field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetUpsertStateLocals() []KeyValue {
	if o == nil || IsNil(o.UpsertStateLocals) {
		var ret []KeyValue
		return ret
	}
	return o.UpsertStateLocals
}

// GetUpsertStateLocalsOk returns a tuple with the UpsertStateLocals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetUpsertStateLocalsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.UpsertStateLocals) {
		return nil, false
	}
	return o.UpsertStateLocals, true
}

// HasUpsertStateLocals returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasUpsertStateLocals() bool {
	if o != nil && !IsNil(o.UpsertStateLocals) {
		return true
	}

	return false
}

// SetUpsertStateLocals gets a reference to the given []KeyValue and assigns it to the UpsertStateLocals field.
func (o *WorkflowStateWaitUntilResponse) SetUpsertStateLocals(v []KeyValue) {
	o.UpsertStateLocals = v
}

// GetRecordEvents returns the RecordEvents field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetRecordEvents() []KeyValue {
	if o == nil || IsNil(o.RecordEvents) {
		var ret []KeyValue
		return ret
	}
	return o.RecordEvents
}

// GetRecordEventsOk returns a tuple with the RecordEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetRecordEventsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.RecordEvents) {
		return nil, false
	}
	return o.RecordEvents, true
}

// HasRecordEvents returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasRecordEvents() bool {
	if o != nil && !IsNil(o.RecordEvents) {
		return true
	}

	return false
}

// SetRecordEvents gets a reference to the given []KeyValue and assigns it to the RecordEvents field.
func (o *WorkflowStateWaitUntilResponse) SetRecordEvents(v []KeyValue) {
	o.RecordEvents = v
}

// GetPublishToInterStateChannel returns the PublishToInterStateChannel field value if set, zero value otherwise.
func (o *WorkflowStateWaitUntilResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing {
	if o == nil || IsNil(o.PublishToInterStateChannel) {
		var ret []InterStateChannelPublishing
		return ret
	}
	return o.PublishToInterStateChannel
}

// GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateWaitUntilResponse) GetPublishToInterStateChannelOk() ([]InterStateChannelPublishing, bool) {
	if o == nil || IsNil(o.PublishToInterStateChannel) {
		return nil, false
	}
	return o.PublishToInterStateChannel, true
}

// HasPublishToInterStateChannel returns a boolean if a field has been set.
func (o *WorkflowStateWaitUntilResponse) HasPublishToInterStateChannel() bool {
	if o != nil && !IsNil(o.PublishToInterStateChannel) {
		return true
	}

	return false
}

// SetPublishToInterStateChannel gets a reference to the given []InterStateChannelPublishing and assigns it to the PublishToInterStateChannel field.
func (o *WorkflowStateWaitUntilResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing) {
	o.PublishToInterStateChannel = v
}

func (o WorkflowStateWaitUntilResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateWaitUntilResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalActivityInput) {
		toSerialize["localActivityInput"] = o.LocalActivityInput
	}
	if !IsNil(o.UpsertSearchAttributes) {
		toSerialize["upsertSearchAttributes"] = o.UpsertSearchAttributes
	}
	if !IsNil(o.UpsertDataObjects) {
		toSerialize["upsertDataObjects"] = o.UpsertDataObjects
	}
	if !IsNil(o.CommandRequest) {
		toSerialize["commandRequest"] = o.CommandRequest
	}
	if !IsNil(o.UpsertStateLocals) {
		toSerialize["upsertStateLocals"] = o.UpsertStateLocals
	}
	if !IsNil(o.RecordEvents) {
		toSerialize["recordEvents"] = o.RecordEvents
	}
	if !IsNil(o.PublishToInterStateChannel) {
		toSerialize["publishToInterStateChannel"] = o.PublishToInterStateChannel
	}
	return toSerialize, nil
}

type NullableWorkflowStateWaitUntilResponse struct {
	value *WorkflowStateWaitUntilResponse
	isSet bool
}

func (v NullableWorkflowStateWaitUntilResponse) Get() *WorkflowStateWaitUntilResponse {
	return v.value
}

func (v *NullableWorkflowStateWaitUntilResponse) Set(val *WorkflowStateWaitUntilResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateWaitUntilResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateWaitUntilResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateWaitUntilResponse(val *WorkflowStateWaitUntilResponse) *NullableWorkflowStateWaitUntilResponse {
	return &NullableWorkflowStateWaitUntilResponse{value: val, isSet: true}
}

func (v NullableWorkflowStateWaitUntilResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateWaitUntilResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
