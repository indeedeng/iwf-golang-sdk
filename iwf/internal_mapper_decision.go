package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"strings"
)

func toIdlDecision(from *StateDecision, wfType string, registry Registry, encoder ObjectEncoder) (*iwfidl.StateDecision, error) {
	var mvs []iwfidl.StateMovement
	for _, fromMv := range from.NextStates {
		input, err := encoder.Encode(fromMv.NextStateInput)
		if err != nil {
			return nil, err
		}
		var options *iwfidl.WorkflowStateOptions
		if !strings.HasPrefix(fromMv.NextStateId, ReservedStateIdPrefix) {
			stateDef := registry.getWorkflowStateDef(wfType, fromMv.NextStateId)
			options = stateDef.State.GetStateOptions()
		}
		mv := iwfidl.StateMovement{
			StateId:      fromMv.NextStateId,
			StateInput:   input,
			StateOptions: options,
		}
		mvs = append(mvs, mv)
	}
	return &iwfidl.StateDecision{
		NextStates: mvs,
	}, nil
}
