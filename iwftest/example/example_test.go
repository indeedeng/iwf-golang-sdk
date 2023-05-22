package example

import (
	"github.com/golang/mock/gomock"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwftest"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockWfCtx *iwftest.MockWorkflowContext
var mockPersistence *iwftest.MockPersistence
var mockCommunication *iwftest.MockCommunication
var emptyCmdResults = iwf.CommandResults{}
var testCustomer = "customer1"
var emptyObj = iwftest.NewTestObject(testCustomer)

func beforeEach(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockWfCtx = iwftest.NewMockWorkflowContext(ctrl)
	mockPersistence = iwftest.NewMockPersistence(ctrl)
	mockCommunication = iwftest.NewMockCommunication(ctrl)
}

func TestInitState_WaitUntil(t *testing.T) {
	beforeEach(t)

	state := NewInitState()

	mockPersistence.EXPECT().SetDataAttribute(keyCustomer, testCustomer)
	cmdReq, err := state.WaitUntil(mockWfCtx, emptyObj, mockPersistence, mockCommunication)
	assert.Nil(t, err)
	assert.Equal(t, iwf.EmptyCommandRequest(), cmdReq)
}

func TestInitState_Execute(t *testing.T) {
	beforeEach(t)

	state := NewInitState()
	input := iwftest.NewTestObject(testCustomer)

	decision, err := state.Execute(mockWfCtx, input, emptyCmdResults, mockPersistence, mockCommunication)
	assert.Nil(t, err)
	assert.Equal(t, iwf.GracefulCompletingWorkflow, decision)
}
