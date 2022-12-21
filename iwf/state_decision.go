package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type StateDecision struct {
	NextStates []StateMovement
}

type StateMovement struct {
	NextStateId      string
	NextStateInput   interface{}
	NextStateOptions *iwfidl.WorkflowStateOptions
}
