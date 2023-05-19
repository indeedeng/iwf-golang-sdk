package iwf

type StateMovement struct {
	// NextStateId is required
	NextStateId string
	// NextStateInput is optional, it's also used as workflow result for GracefulCompletingWorkflowStateId and ForceCompletingWorkflowStateId
	NextStateInput interface{}
}

const (
	ReservedStateIdPrefix             = "_SYS_"
	GracefulCompletingWorkflowStateId = "_SYS_GRACEFUL_COMPLETING_WORKFLOW"
	ForceCompletingWorkflowStateId    = "_SYS_FORCE_COMPLETING_WORKFLOW"
	ForceFailingWorkflowStateId       = "_SYS_FORCE_FAILING_WORKFLOW"
)

func NewStateMovement(st WorkflowState, in interface{}) StateMovement {
	return StateMovement{
		NextStateId:    GetFinalWorkflowStateId(st),
		NextStateInput: in,
	}
}