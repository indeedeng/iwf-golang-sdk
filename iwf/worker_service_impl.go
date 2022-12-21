package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
)

type workerServiceImpl struct {
	registry Registry
	options  WorkerOptions
}

func (w *workerServiceImpl) HandleWorkflowStateStart(request iwfidl.WorkflowStateStartRequest) (iwfidl.WorkflowStateStartResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w *workerServiceImpl) HandleWorkflowStateDecide(request iwfidl.WorkflowStateDecideRequest) (iwfidl.WorkflowStateDecideResponse, error) {
	//TODO implement me
	panic("implement me")
}
