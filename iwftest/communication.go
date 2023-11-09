// Code generated by MockGen. DO NOT EDIT.
// Source: ./communication.go
//
// Generated by this command:
//
//	mockgen -source=./communication.go -package=iwftest -destination=../iwftest/communication.go
//
// Package iwftest is a generated GoMock package.
package iwftest

import (
	reflect "reflect"

	iwfidl "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	iwf "github.com/indeedeng/iwf-golang-sdk/iwf"
	gomock "go.uber.org/mock/gomock"
)

// MockCommunication is a mock of Communication interface.
type MockCommunication struct {
	ctrl     *gomock.Controller
	recorder *MockCommunicationMockRecorder
}

// MockCommunicationMockRecorder is the mock recorder for MockCommunication.
type MockCommunicationMockRecorder struct {
	mock *MockCommunication
}

// NewMockCommunication creates a new mock instance.
func NewMockCommunication(ctrl *gomock.Controller) *MockCommunication {
	mock := &MockCommunication{ctrl: ctrl}
	mock.recorder = &MockCommunicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommunication) EXPECT() *MockCommunicationMockRecorder {
	return m.recorder
}

// GetToPublishInternalChannel mocks base method.
func (m *MockCommunication) GetToPublishInternalChannel() map[string][]iwfidl.EncodedObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToPublishInternalChannel")
	ret0, _ := ret[0].(map[string][]iwfidl.EncodedObject)
	return ret0
}

// GetToPublishInternalChannel indicates an expected call of GetToPublishInternalChannel.
func (mr *MockCommunicationMockRecorder) GetToPublishInternalChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToPublishInternalChannel", reflect.TypeOf((*MockCommunication)(nil).GetToPublishInternalChannel))
}

// GetToTriggerStateMovements mocks base method.
func (m *MockCommunication) GetToTriggerStateMovements() []iwf.StateMovement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToTriggerStateMovements")
	ret0, _ := ret[0].([]iwf.StateMovement)
	return ret0
}

// GetToTriggerStateMovements indicates an expected call of GetToTriggerStateMovements.
func (mr *MockCommunicationMockRecorder) GetToTriggerStateMovements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToTriggerStateMovements", reflect.TypeOf((*MockCommunication)(nil).GetToTriggerStateMovements))
}

// PublishInternalChannel mocks base method.
func (m *MockCommunication) PublishInternalChannel(channelName string, value any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishInternalChannel", channelName, value)
}

// PublishInternalChannel indicates an expected call of PublishInternalChannel.
func (mr *MockCommunicationMockRecorder) PublishInternalChannel(channelName, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishInternalChannel", reflect.TypeOf((*MockCommunication)(nil).PublishInternalChannel), channelName, value)
}

// TriggerStateMovements mocks base method.
func (m *MockCommunication) TriggerStateMovements(movements ...iwf.StateMovement) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range movements {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "TriggerStateMovements", varargs...)
}

// TriggerStateMovements indicates an expected call of TriggerStateMovements.
func (mr *MockCommunicationMockRecorder) TriggerStateMovements(movements ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerStateMovements", reflect.TypeOf((*MockCommunication)(nil).TriggerStateMovements), movements...)
}

// MockcommunicationInternal is a mock of communicationInternal interface.
type MockcommunicationInternal struct {
	ctrl     *gomock.Controller
	recorder *MockcommunicationInternalMockRecorder
}

// MockcommunicationInternalMockRecorder is the mock recorder for MockcommunicationInternal.
type MockcommunicationInternalMockRecorder struct {
	mock *MockcommunicationInternal
}

// NewMockcommunicationInternal creates a new mock instance.
func NewMockcommunicationInternal(ctrl *gomock.Controller) *MockcommunicationInternal {
	mock := &MockcommunicationInternal{ctrl: ctrl}
	mock.recorder = &MockcommunicationInternalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcommunicationInternal) EXPECT() *MockcommunicationInternalMockRecorder {
	return m.recorder
}

// GetToPublishInternalChannel mocks base method.
func (m *MockcommunicationInternal) GetToPublishInternalChannel() map[string][]iwfidl.EncodedObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToPublishInternalChannel")
	ret0, _ := ret[0].(map[string][]iwfidl.EncodedObject)
	return ret0
}

// GetToPublishInternalChannel indicates an expected call of GetToPublishInternalChannel.
func (mr *MockcommunicationInternalMockRecorder) GetToPublishInternalChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToPublishInternalChannel", reflect.TypeOf((*MockcommunicationInternal)(nil).GetToPublishInternalChannel))
}

// GetToTriggerStateMovements mocks base method.
func (m *MockcommunicationInternal) GetToTriggerStateMovements() []iwf.StateMovement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToTriggerStateMovements")
	ret0, _ := ret[0].([]iwf.StateMovement)
	return ret0
}

// GetToTriggerStateMovements indicates an expected call of GetToTriggerStateMovements.
func (mr *MockcommunicationInternalMockRecorder) GetToTriggerStateMovements() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToTriggerStateMovements", reflect.TypeOf((*MockcommunicationInternal)(nil).GetToTriggerStateMovements))
}
