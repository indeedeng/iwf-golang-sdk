package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type StateDecision struct {
	NextStates []StateMovement
}

func SingleNextState(stateId string, input interface{}) StateDecision {
	return SingleNextStateWithOptions(stateId, input, nil)
}

func SingleNextStateWithOptions(stateId string, input interface{}, options *iwfidl.WorkflowStateOptions) StateDecision {
	return StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:      stateId,
				NextStateInput:   input,
				NextStateOptions: options,
			},
		},
	}
}

// DeadEnd means no next step for this thread. It is essentially graceful completion without a workflow result
var DeadEnd = StateDecision{}

var ForceFailingWorkflow = StateDecision{
	NextStates: []StateMovement{
		{
			NextStateId: ForceFailingWorkflowStateId,
		},
	},
}

var GracefulCompletingWorkflow = GracefulCompleteWorkflow(nil)

func GracefulCompleteWorkflow(output interface{}) StateDecision {
	return StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:    GracefulCompletingWorkflowStateId,
				NextStateInput: output,
			},
		},
	}
}

var ForceCompletingWorkflow = ForceCompleteWorkflow(nil)

func ForceCompleteWorkflow(output interface{}) StateDecision {
	return StateDecision{
		NextStates: []StateMovement{
			{
				NextStateId:    ForceCompletingWorkflowStateId,
				NextStateInput: output,
			},
		},
	}
}
