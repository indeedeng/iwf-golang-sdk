package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type interStateWorkflow struct{}

const interStateChannel1 = "test-inter-state-channel-1"
const interStateChannel2 = "test-inter-state-channel-2"

func (b interStateWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.NewStartingState(&interStateWorkflowState0{}),
		iwf.NewNonStartingState(&interStateWorkflowState1{}),
		iwf.NewNonStartingState(&interStateWorkflowState2{}),
	}
}

func (b interStateWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return nil
}

func (b interStateWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.NewInterstateChannelDef(interStateChannel1),
		iwf.NewInterstateChannelDef(interStateChannel2),
	}
}

func (b interStateWorkflow) GetWorkflowType() string {
	return ""
}
