package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

const (
	WorkflowStateStartApi  = "/api/v1/workflowState/start"
	WorkflowStateDecideApi = "/api/v1/workflowState/decide"
)

type WorkerService interface {
	HandleWorkflowStateStart(request iwfidl.WorkflowStateStartRequest) (iwfidl.WorkflowStateStartResponse, error)
	HandleWorkflowStateDecide(request iwfidl.WorkflowStateDecideRequest) (iwfidl.WorkflowStateDecideResponse, error)
}
