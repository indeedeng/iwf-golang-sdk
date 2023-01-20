package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type registryImpl struct {
	workflowStore              map[string]Workflow
	workflowStartingState      map[string]WorkflowState
	workflowStateStore         map[string]map[string]StateDef
	signalNameStore            map[string]map[string]bool
	interStateChannelNameStore map[string]map[string]bool
	dataObjectKeyStore         map[string]map[string]bool
	searchAttributeTypeStore   map[string]map[string]iwfidl.SearchAttributeValueType
}

func (r *registryImpl) AddWorkflows(workflows ...Workflow) error {
	for _, wf := range workflows {
		err := r.AddWorkflow(wf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *registryImpl) AddWorkflow(wf Workflow) error {
	if err := r.registerWorkflow(wf); err != nil {
		return err
	}
	if err := r.registerWorkflowState(wf); err != nil {
		return err
	}
	if err := r.registerWorkflowCommunicationSchema(wf); err != nil {
		return err
	}
	return r.registerWorkflowPersistenceSchema(wf)
}

func (r *registryImpl) GetAllRegisteredWorkflowTypes() []string {
	var res []string
	for wfType := range r.workflowStore {
		res = append(res, wfType)
	}
	return res
}

func (r *registryImpl) getWorkflowStartingState(wfType string) WorkflowState {
	return r.workflowStartingState[wfType]
}

func (r *registryImpl) getWorkflowStateDef(wfType string, id string) StateDef {
	return r.workflowStateStore[wfType][id]
}

func (r *registryImpl) getWorkflowSignalNameStore(wfType string) map[string]bool {
	return r.signalNameStore[wfType]
}

func (r *registryImpl) getWorkflowInterStateChannelNameStore(wfType string) map[string]bool {
	return r.interStateChannelNameStore[wfType]
}

func (r *registryImpl) getWorkflowDataObjectKeyStore(wfType string) map[string]bool {
	return r.dataObjectKeyStore[wfType]
}

func (r *registryImpl) getSearchAttributeTypeStore(wfType string) map[string]iwfidl.SearchAttributeValueType {
	return r.searchAttributeTypeStore[wfType]
}

func (r *registryImpl) registerWorkflow(wf Workflow) error {
	wfType := GetDefaultWorkflowType(wf)
	_, ok := r.workflowStore[wfType]
	if ok {
		return NewWorkflowDefinitionError("workflow type conflict: " + wfType)
	}
	r.workflowStore[wfType] = wf
	return nil
}

func (r *registryImpl) registerWorkflowState(wf Workflow) error {
	wfType := GetDefaultWorkflowType(wf)
	if len(wf.GetStates()) == 0 {
		return NewWorkflowDefinitionFmtError("Workflow type %s must contain at least one workflow state", wfType)
	}
	stateMap := map[string]StateDef{}
	var startingState WorkflowState
	for _, state := range wf.GetStates() {
		stateMap[state.State.GetStateId()] = state
		if state.CanStartWorkflow {
			if startingState == nil {
				startingState = state.State
			} else {
				return NewWorkflowDefinitionError("Workflow must contain exactly one starting states: " + wfType)
			}
		}
	}
	r.workflowStartingState[wfType] = startingState
	r.workflowStateStore[wfType] = stateMap
	return nil
}

func (r *registryImpl) registerWorkflowCommunicationSchema(wf Workflow) error {
	wfType := GetDefaultWorkflowType(wf)
	signalMap := map[string]bool{}
	interStateChannel := map[string]bool{}
	for _, methodDef := range wf.GetCommunicationSchema() {
		if methodDef.CommunicationMethod == CommunicationMethodSignalChannel {
			signalMap[methodDef.Name] = true
		} else if methodDef.CommunicationMethod == CommunicationMethodInterstateChannel {
			interStateChannel[methodDef.Name] = true
		} else {
			return NewWorkflowDefinitionError("invalid CommunicationMethod definition " + string(methodDef.CommunicationMethod))
		}
	}
	r.signalNameStore[wfType] = signalMap
	r.interStateChannelNameStore[wfType] = interStateChannel
	return nil
}

func (r *registryImpl) registerWorkflowPersistenceSchema(wf Workflow) error {
	wfType := GetDefaultWorkflowType(wf)
	dataObjectKeys := map[string]bool{}
	searchAttributes := map[string]iwfidl.SearchAttributeValueType{}
	for _, pers := range wf.GetPersistenceSchema() {
		if pers.FieldType == PersistenceFieldTypeDataObject {
			dataObjectKeys[pers.Key] = true
		} else if pers.FieldType == PersistenceFieldTypeSearchAttribute && pers.SearchAttributeType != nil {
			searchAttributes[pers.Key] = *pers.SearchAttributeType
		} else {
			return NewWorkflowDefinitionFmtError("invalid PersistenceField definition %s for key %s ", string(pers.FieldType), pers.Key)
		}
	}
	r.dataObjectKeyStore[wfType] = dataObjectKeys
	r.searchAttributeTypeStore[wfType] = searchAttributes
	return nil
}
