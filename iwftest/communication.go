// Code generated by MockGen. DO NOT EDIT.
// Source: iwf/communication.go

// Package iwftest is a generated GoMock package.
package iwftest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	iwfidl "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
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

// GetToPublishInterStateChannel mocks base method.
func (m *MockCommunication) GetToPublishInterStateChannel() map[string][]iwfidl.EncodedObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToPublishInterStateChannel")
	ret0, _ := ret[0].(map[string][]iwfidl.EncodedObject)
	return ret0
}

// GetToPublishInterStateChannel indicates an expected call of GetToPublishInterStateChannel.
func (mr *MockCommunicationMockRecorder) GetToPublishInterStateChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToPublishInterStateChannel", reflect.TypeOf((*MockCommunication)(nil).GetToPublishInterStateChannel))
}

// PublishInterstateChannel mocks base method.
func (m *MockCommunication) PublishInterstateChannel(channelName string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishInterstateChannel", channelName, value)
}

// PublishInterstateChannel indicates an expected call of PublishInterstateChannel.
func (mr *MockCommunicationMockRecorder) PublishInterstateChannel(channelName, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishInterstateChannel", reflect.TypeOf((*MockCommunication)(nil).PublishInterstateChannel), channelName, value)
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

// GetToPublishInterStateChannel mocks base method.
func (m *MockcommunicationInternal) GetToPublishInterStateChannel() map[string][]iwfidl.EncodedObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToPublishInterStateChannel")
	ret0, _ := ret[0].(map[string][]iwfidl.EncodedObject)
	return ret0
}

// GetToPublishInterStateChannel indicates an expected call of GetToPublishInterStateChannel.
func (mr *MockcommunicationInternalMockRecorder) GetToPublishInterStateChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToPublishInterStateChannel", reflect.TypeOf((*MockcommunicationInternal)(nil).GetToPublishInterStateChannel))
}