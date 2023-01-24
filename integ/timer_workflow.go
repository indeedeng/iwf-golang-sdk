package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type timerWorkflow struct{}

func (b timerWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&timerWorkflowState1{}),
	}
}

func (b timerWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return nil
}

func (b timerWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return nil
}

func (b timerWorkflow) GetWorkflowType() string {
	return ""
}
