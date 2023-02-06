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

// WorkflowErrorType the model 'WorkflowErrorType'
type WorkflowErrorType string

// List of WorkflowErrorType
const (
	STATE_DECISION_FAILING_WORKFLOW_ERROR_TYPE WorkflowErrorType = "STATE_DECISION_FAILING_WORKFLOW_ERROR_TYPE"
	STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE WorkflowErrorType = "STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE"
	INVALID_USER_WORKFLOW_CODE_ERROR_TYPE WorkflowErrorType = "INVALID_USER_WORKFLOW_CODE_ERROR_TYPE"
	SERVER_INTERNAL_ERROR_TYPE WorkflowErrorType = "SERVER_INTERNAL_ERROR_TYPE"
)

// All allowed values of WorkflowErrorType enum
var AllowedWorkflowErrorTypeEnumValues = []WorkflowErrorType{
	"STATE_DECISION_FAILING_WORKFLOW_ERROR_TYPE",
	"STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE",
	"INVALID_USER_WORKFLOW_CODE_ERROR_TYPE",
	"SERVER_INTERNAL_ERROR_TYPE",
}

func (v *WorkflowErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkflowErrorType(value)
	for _, existing := range AllowedWorkflowErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkflowErrorType", value)
}

// NewWorkflowErrorTypeFromValue returns a pointer to a valid WorkflowErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkflowErrorTypeFromValue(v string) (*WorkflowErrorType, error) {
	ev := WorkflowErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkflowErrorType: valid values are %v", v, AllowedWorkflowErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkflowErrorType) IsValid() bool {
	for _, existing := range AllowedWorkflowErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowErrorType value
func (v WorkflowErrorType) Ptr() *WorkflowErrorType {
	return &v
}

type NullableWorkflowErrorType struct {
	value *WorkflowErrorType
	isSet bool
}

func (v NullableWorkflowErrorType) Get() *WorkflowErrorType {
	return v.value
}

func (v *NullableWorkflowErrorType) Set(val *WorkflowErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowErrorType(val *WorkflowErrorType) *NullableWorkflowErrorType {
	return &NullableWorkflowErrorType{value: val, isSet: true}
}

func (v NullableWorkflowErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

