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

// checks if the HealthInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthInfo{}

// HealthInfo struct for HealthInfo
type HealthInfo struct {
	Condition *string `json:"condition,omitempty"`
	Hostname  *string `json:"hostname,omitempty"`
	Duration  *int32  `json:"duration,omitempty"`
}

// NewHealthInfo instantiates a new HealthInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthInfo() *HealthInfo {
	this := HealthInfo{}
	return &this
}

// NewHealthInfoWithDefaults instantiates a new HealthInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthInfoWithDefaults() *HealthInfo {
	this := HealthInfo{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *HealthInfo) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthInfo) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *HealthInfo) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *HealthInfo) SetCondition(v string) {
	o.Condition = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HealthInfo) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthInfo) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HealthInfo) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HealthInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *HealthInfo) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthInfo) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *HealthInfo) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *HealthInfo) SetDuration(v int32) {
	o.Duration = &v
}

func (o HealthInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableHealthInfo struct {
	value *HealthInfo
	isSet bool
}

func (v NullableHealthInfo) Get() *HealthInfo {
	return v.value
}

func (v *NullableHealthInfo) Set(val *HealthInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthInfo(val *HealthInfo) *NullableHealthInfo {
	return &NullableHealthInfo{value: val, isSet: true}
}

func (v NullableHealthInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
