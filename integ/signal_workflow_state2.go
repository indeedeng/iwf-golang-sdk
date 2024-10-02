package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"time"
)

type signalWorkflowState2 struct {
	iwf.WorkflowStateDefaults
}

const timerCommandId = "timerId"
const signalCommandId = "s1"

func (b signalWorkflowState2) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	var val int
	input.Get(&val)
	if val != 10 {
		panic(fmt.Sprintf("input value should be 10 but is %v", val))
	}

	return iwf.AnyCommandCombinationsCompletedRequest(
		[][]string{
			{signalCommandId, timerCommandId},
		},
		iwf.NewSignalCommand(signalCommandId, testChannelName1),
		iwf.NewSignalCommand(signalCommandId, testChannelName2),
		iwf.NewTimerCommandByDuration(timerCommandId, 24*time.Hour),
	), nil
}

func (b signalWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	signal0 := commandResults.Signals[0]
	signal1 := commandResults.Signals[1]
	timer := commandResults.Timers[0]

	if signal0.CommandId != signalCommandId || signal0.ChannelName != testChannelName1 || signal0.Status != iwfidl.RECEIVED {
		panic(testChannelName1 + " should be waiting....")
	}

	if signal1.CommandId != signalCommandId || signal1.ChannelName != testChannelName2 || signal1.Status != iwfidl.WAITING {
		panic(testChannelName2 + " should be received....")
	}

	if timer.CommandId != timerCommandId || timer.Status != iwfidl.FIRED {
		panic("timer should be fired")
	}

	var val int
	signal0.SignalValue.Get(&val)
	if val != 100 {
		panic("signal value should be 100")
	}

	return iwf.GracefulCompleteWorkflow(val), nil
}
