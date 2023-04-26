package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"testing"
)

type myWf struct {
}

type myState struct {
}

func (m myState) GetStateId() string {
	return "stateId"
}

func (m myState) Start(ctx WorkflowContext, input Object, persistence Persistence, communication Communication) (*CommandRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (m myState) Decide(ctx WorkflowContext, input Object, commandResults CommandResults, persistence Persistence, communication Communication) (*StateDecision, error) {
	//TODO implement me
	panic("implement me")
}

func (m myState) GetStateOptions() *iwfidl.WorkflowStateOptions {
	//TODO implement me
	panic("implement me")
}

func (m myWf) GetWorkflowStates() []StateDef {
	return []StateDef{
		{
			State:            myState{},
			CanStartWorkflow: true,
		},
	}
}

func (m myWf) GetPersistenceSchema() []PersistenceFieldDef {
	return nil
}

func (m myWf) GetCommunicationSchema() []CommunicationMethodDef {
	return nil
}

func (m myWf) GetWorkflowType() string {
	return ""
}

func TestNewRegistry(t *testing.T) {
	registry := NewRegistry()
	err := registry.AddWorkflow(&myWf{})
	NewClient(registry, nil)
	if err != nil {
		t.Fail()
	}
}
