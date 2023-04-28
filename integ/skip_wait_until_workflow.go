package integ

import "github.com/indeedeng/iwf-golang-sdk/iwf"

type skipWaitUntilWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (b skipWaitUntilWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&skipWaitUntilState1{}),
		iwf.NonStartingStateDef(&skipWaitUntilState2{}),
	}
}

type skipWaitUntilWorkflow2 struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
	iwf.EmptyCommunicationSchema
}

func (b skipWaitUntilWorkflow2) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(skipWaitUntilState1{}),
		iwf.NonStartingStateDef(skipWaitUntilState2{}),
	}
}
