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

// InterStateChannelResult struct for InterStateChannelResult
type InterStateChannelResult struct {
	CommandId string `json:"commandId"`
	RequestStatus string `json:"requestStatus"`
	ChannelName string `json:"channelName"`
	Value *EncodedObject `json:"value,omitempty"`
}

// NewInterStateChannelResult instantiates a new InterStateChannelResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterStateChannelResult(commandId string, requestStatus string, channelName string) *InterStateChannelResult {
	this := InterStateChannelResult{}
	this.CommandId = commandId
	this.RequestStatus = requestStatus
	this.ChannelName = channelName
	return &this
}

// NewInterStateChannelResultWithDefaults instantiates a new InterStateChannelResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterStateChannelResultWithDefaults() *InterStateChannelResult {
	this := InterStateChannelResult{}
	return &this
}

// GetCommandId returns the CommandId field value
func (o *InterStateChannelResult) GetCommandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelResult) GetCommandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandId, true
}

// SetCommandId sets field value
func (o *InterStateChannelResult) SetCommandId(v string) {
	o.CommandId = v
}

// GetRequestStatus returns the RequestStatus field value
func (o *InterStateChannelResult) GetRequestStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelResult) GetRequestStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestStatus, true
}

// SetRequestStatus sets field value
func (o *InterStateChannelResult) SetRequestStatus(v string) {
	o.RequestStatus = v
}

// GetChannelName returns the ChannelName field value
func (o *InterStateChannelResult) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelResult) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *InterStateChannelResult) SetChannelName(v string) {
	o.ChannelName = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterStateChannelResult) GetValue() EncodedObject {
	if o == nil || o.Value == nil {
		var ret EncodedObject
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterStateChannelResult) GetValueOk() (*EncodedObject, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterStateChannelResult) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given EncodedObject and assigns it to the Value field.
func (o *InterStateChannelResult) SetValue(v EncodedObject) {
	o.Value = &v
}

func (o InterStateChannelResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commandId"] = o.CommandId
	}
	if true {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if true {
		toSerialize["channelName"] = o.ChannelName
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInterStateChannelResult struct {
	value *InterStateChannelResult
	isSet bool
}

func (v NullableInterStateChannelResult) Get() *InterStateChannelResult {
	return v.value
}

func (v *NullableInterStateChannelResult) Set(val *InterStateChannelResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInterStateChannelResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInterStateChannelResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterStateChannelResult(val *InterStateChannelResult) *NullableInterStateChannelResult {
	return &NullableInterStateChannelResult{value: val, isSet: true}
}

func (v NullableInterStateChannelResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterStateChannelResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


