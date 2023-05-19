package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type rpcWorkflow struct {
	iwf.DefaultWorkflowType
	iwf.EmptyPersistenceSchema
}

func (b rpcWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.InternalChannelDef("test"),
		iwf.RPCMethodDef(b.TestRPC, nil),
		iwf.RPCMethodDef(b.TestErrorRPC, nil),
	}
}

func (b rpcWorkflow) GetWorkflowStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.StartingStateDef(&rpcWorkflowState1{}),
	}
}

func (b rpcWorkflow) TestRPC(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (interface{}, error) {
	var i int
	input.Get(&i)
	i++
	communication.PublishInternalChannel("test", i)
	return i, nil
}

func (b rpcWorkflow) TestErrorRPC(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (interface{}, error) {
	return nil, fmt.Errorf("test error")
}

type rpcWorkflowState1 struct {
	iwf.DefaultStateIdAndOptions
}

func (b rpcWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	return iwf.AllCommandsCompletedRequest(
		iwf.NewInternalChannelCommand("", "test"),
	), nil
}

func (b rpcWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	var i int
	input.Get(&i)
	var j int
	commandResults.InternalChannelCommands[0].Value.Get(&j)
	return iwf.GracefulCompleteWorkflow(i + j), nil
}
