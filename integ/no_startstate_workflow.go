package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type noStartStateWorkflow struct {
	iwf.WorkflowDefaults
}

func (b noStartStateWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.RPCMethodDef(b.TestRPC, nil),
	}
}

func (b noStartStateWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.NonStartingStateDef(&noStartStateWorkflowState1{}),
	}
}

func (b noStartStateWorkflow) TestRPC(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (interface{}, error) {
	var i int
	input.Get(&i)
	i++
	communication.TriggerStateMovements(iwf.NewStateMovement(noStartStateWorkflowState1{}, nil))
	return i, nil
}

type noStartStateWorkflowState1 struct {
	iwf.WorkflowStateDefaults
	iwf.NoWaitUntil
}

func (b noStartStateWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.GracefulCompletingWorkflow, nil
}
