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

// WorkflowResetType the model 'WorkflowResetType'
type WorkflowResetType string

// List of WorkflowResetType
const (
	HISTORY_EVENT_ID   WorkflowResetType = "HISTORY_EVENT_ID"
	BEGINNING          WorkflowResetType = "BEGINNING"
	HISTORY_EVENT_TIME WorkflowResetType = "HISTORY_EVENT_TIME"
	STATE_ID           WorkflowResetType = "STATE_ID"
	STATE_EXECUTION_ID WorkflowResetType = "STATE_EXECUTION_ID"
)

// All allowed values of WorkflowResetType enum
var AllowedWorkflowResetTypeEnumValues = []WorkflowResetType{
	"HISTORY_EVENT_ID",
	"BEGINNING",
	"HISTORY_EVENT_TIME",
	"STATE_ID",
	"STATE_EXECUTION_ID",
}

func (v *WorkflowResetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkflowResetType(value)
	for _, existing := range AllowedWorkflowResetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkflowResetType", value)
}

// NewWorkflowResetTypeFromValue returns a pointer to a valid WorkflowResetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkflowResetTypeFromValue(v string) (*WorkflowResetType, error) {
	ev := WorkflowResetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkflowResetType: valid values are %v", v, AllowedWorkflowResetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkflowResetType) IsValid() bool {
	for _, existing := range AllowedWorkflowResetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowResetType value
func (v WorkflowResetType) Ptr() *WorkflowResetType {
	return &v
}

type NullableWorkflowResetType struct {
	value *WorkflowResetType
	isSet bool
}

func (v NullableWorkflowResetType) Get() *WorkflowResetType {
	return v.value
}

func (v *NullableWorkflowResetType) Set(val *WorkflowResetType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowResetType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowResetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowResetType(val *WorkflowResetType) *NullableWorkflowResetType {
	return &NullableWorkflowResetType{value: val, isSet: true}
}

func (v NullableWorkflowResetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowResetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
