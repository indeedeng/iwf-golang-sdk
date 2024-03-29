/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
	"fmt"
)

// CommandWaitingType the model 'CommandWaitingType'
type CommandWaitingType string

// List of CommandWaitingType
const (
	ALL_COMPLETED             CommandWaitingType = "ALL_COMPLETED"
	ANY_COMPLETED             CommandWaitingType = "ANY_COMPLETED"
	ANY_COMBINATION_COMPLETED CommandWaitingType = "ANY_COMBINATION_COMPLETED"
)

// All allowed values of CommandWaitingType enum
var AllowedCommandWaitingTypeEnumValues = []CommandWaitingType{
	"ALL_COMPLETED",
	"ANY_COMPLETED",
	"ANY_COMBINATION_COMPLETED",
}

func (v *CommandWaitingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommandWaitingType(value)
	for _, existing := range AllowedCommandWaitingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommandWaitingType", value)
}

// NewCommandWaitingTypeFromValue returns a pointer to a valid CommandWaitingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommandWaitingTypeFromValue(v string) (*CommandWaitingType, error) {
	ev := CommandWaitingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommandWaitingType: valid values are %v", v, AllowedCommandWaitingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommandWaitingType) IsValid() bool {
	for _, existing := range AllowedCommandWaitingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommandWaitingType value
func (v CommandWaitingType) Ptr() *CommandWaitingType {
	return &v
}

type NullableCommandWaitingType struct {
	value *CommandWaitingType
	isSet bool
}

func (v NullableCommandWaitingType) Get() *CommandWaitingType {
	return v.value
}

func (v *NullableCommandWaitingType) Set(val *CommandWaitingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandWaitingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandWaitingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandWaitingType(val *CommandWaitingType) *NullableCommandWaitingType {
	return &NullableCommandWaitingType{value: val, isSet: true}
}

func (v NullableCommandWaitingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandWaitingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
