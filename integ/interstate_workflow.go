package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type interStateWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
}

const interStateChannel1 = "test-inter-state-channel-1"
const interStateChannel2 = "test-inter-state-channel-2"

func (b interStateWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&interStateWorkflowState0{}),
		iwf.NonStartingStateDef(&interStateWorkflowState1{}),
		iwf.NonStartingStateDef(&interStateWorkflowState2{}),
	}
}

func (b interStateWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.InterstateChannelDef(interStateChannel1),
		iwf.InterstateChannelDef(interStateChannel2),
	}
}

