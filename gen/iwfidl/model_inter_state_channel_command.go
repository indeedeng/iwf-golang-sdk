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

// checks if the InterStateChannelCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterStateChannelCommand{}

// InterStateChannelCommand struct for InterStateChannelCommand
type InterStateChannelCommand struct {
	CommandId   string `json:"commandId"`
	ChannelName string `json:"channelName"`
}

// NewInterStateChannelCommand instantiates a new InterStateChannelCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterStateChannelCommand(commandId string, channelName string) *InterStateChannelCommand {
	this := InterStateChannelCommand{}
	this.CommandId = commandId
	this.ChannelName = channelName
	return &this
}

// NewInterStateChannelCommandWithDefaults instantiates a new InterStateChannelCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterStateChannelCommandWithDefaults() *InterStateChannelCommand {
	this := InterStateChannelCommand{}
	return &this
}

// GetCommandId returns the CommandId field value
func (o *InterStateChannelCommand) GetCommandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetCommandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommandId, true
}

// SetCommandId sets field value
func (o *InterStateChannelCommand) SetCommandId(v string) {
	o.CommandId = v
}

// GetChannelName returns the ChannelName field value
func (o *InterStateChannelCommand) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *InterStateChannelCommand) SetChannelName(v string) {
	o.ChannelName = v
}

func (o InterStateChannelCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterStateChannelCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commandId"] = o.CommandId
	toSerialize["channelName"] = o.ChannelName
	return toSerialize, nil
}

type NullableInterStateChannelCommand struct {
	value *InterStateChannelCommand
	isSet bool
}

func (v NullableInterStateChannelCommand) Get() *InterStateChannelCommand {
	return v.value
}

func (v *NullableInterStateChannelCommand) Set(val *InterStateChannelCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableInterStateChannelCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableInterStateChannelCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterStateChannelCommand(val *InterStateChannelCommand) *NullableInterStateChannelCommand {
	return &NullableInterStateChannelCommand{value: val, isSet: true}
}

func (v NullableInterStateChannelCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterStateChannelCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
