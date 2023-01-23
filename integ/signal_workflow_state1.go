package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type signalWorkflowState1 struct {
	iwf.DefaultStateIdAndOptions
}

func (b signalWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AnyCommandCompletedRequest(
		iwf.NewSignalCommand("", testChannelName1),
		iwf.NewSignalCommand("", testChannelName2),
	), nil
}

func (b signalWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	signal0 := commandResults.Signals[0]
	signal1 := commandResults.Signals[1]
	if signal0.CommandId != "" || signal0.ChannelName != testChannelName1 || signal0.Status != iwfidl.WAITING {
		panic(testChannelName1 + " should be waiting....")
	}
	if signal1.CommandId == "" && signal1.ChannelName == testChannelName2 && signal1.Status == iwfidl.RECEIVED {
		var value int
		signal1.SignalValue.Get(&value)
		return iwf.SingleNextState(signalWorkflowState2{}, value), nil
	}
	return nil, fmt.Errorf(testChannelName2 + " doesn't receive correct value")
}
