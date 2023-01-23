package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type signalWorkflow struct{}

const testChannelName1 = "test-channel-name-1"
const testChannelName2 = "test-channel-name-2"

func (b signalWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&signalWorkflowState1{}),
		iwf.NonStartingStateDef(&signalWorkflowState2{}),
	}
}

func (b signalWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return nil
}

func (b signalWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.SignalChannelDef(testChannelName1),
		iwf.SignalChannelDef(testChannelName2),
	}
}

func (b signalWorkflow) GetWorkflowType() string {
	return ""
}
