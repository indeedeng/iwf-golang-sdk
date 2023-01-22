package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type interStateWorkflowState1 struct{}

const interStateWorkflowState1Id = "interStateWorkflowState1"

func (b interStateWorkflowState1) GetStateId() string {
	return interStateWorkflowState1Id
}

func (b interStateWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AnyCommandCompletedRequest(
			iwf.NewInterStateChannelCommand("id1", interStateChannel1),
			iwf.NewInterStateChannelCommand("id2", interStateChannel2)),
		nil
}

func (b interStateWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	cmd1 := commandResults.GetInterStateChannelCommandResultById("id1")
	cmd2 := commandResults.GetInterStateChannelCommandResultById("id2")
	cmd2.Value.Get(&i)
	if cmd1.Status == iwfidl.WAITING && i == 2 {
		return iwf.GracefulCompletingWorkflow, nil
	}
	return nil, fmt.Errorf("error in executing " + interStateWorkflowState1Id)
}

func (b interStateWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
