package iwf

type StateDecision interface {
	WaitForMoreCommandResults() bool
	GetNextStates() []StateMovement
	GetUpsertSearchAttributes()
	GetUpsertQueryAttributes()
}

func NewStateDecision(movement ...StateMovement) StateDecision {
	return nil
}

func WaitForMoreCommandResults() StateDecision {
	return nil
}

type StateMovement interface {
	GetNextStateId() string
	GetNextStateInput() interface{}
}

func NewStateMovement(nextStateId string) StateMovement {
	return nil
}

func NewStateMovementWithInput(nextStateId string, nextStateInput interface{}) StateMovement {
	return nil
}

type builtInStateDecision struct {
	movement StateMovement
}

func (b builtInStateDecision) WaitForMoreCommandResults() bool {
	return false
}

func (b builtInStateDecision) GetNextStates() []StateMovement {
	return nil
}

func (b builtInStateDecision) GetUpsertSearchAttributes() {
	return
}

func (b builtInStateDecision) GetUpsertQueryAttributes() {
	return
}
