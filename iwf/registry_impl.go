package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"reflect"
)

type registry struct {
	workflowStore              map[string]Workflow
	workflowStateStore         map[string]map[string]StateDef
	signalNameStore            map[string]map[string]bool
	interStateChannelNameStore map[string]map[string]bool
	dataObjectKeyStore         map[string]map[string]bool
	searchAttributeTypeStore   map[string]map[string]iwfidl.SearchAttributeValueType
}

func (r *registry) AddWorkflow(wf Workflow) error {
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

// GetWorkflowType returns the workflow type that will be registered and used as IwfWorkflowType
// if the workflow is from &myStruct{} under mywf package, the method returns "*mywf.myStruct"
// the "*" is from pointer. If the instance is initiated as myStruct{}, then it returns "mywf.myStruct" without the "*"
func (r *registry) GetWorkflowType(wf Workflow) string {
	wfType := wf.GetWorkflowType()
	if wfType == "" {
		rt := reflect.TypeOf(wf)
		return rt.String()
	}
	return wfType
}

func (r *registry) GetAllWorkflowTypes() []string {
	var res []string
	for wfType := range r.workflowStore {
		res = append(res, wfType)
	}
	return res
}

func (r *registry) registerWorkflow(wf Workflow) error {
	wfType := r.GetWorkflowType(wf)
	_, ok := r.workflowStore[wfType]
	if ok {
		return NewWorkflowDefinitionError("workflow type conflict: " + wfType)
	}
	r.workflowStore[wfType] = wf
	return nil
}

func (r *registry) registerWorkflowState(wf Workflow) error {
	wfType := r.GetWorkflowType(wf)
	if len(wf.GetStates()) == 0 {
		return NewWorkflowDefinitionFmtError("Workflow type %s must contain at least one workflow state", wfType)
	}
	stateMap := map[string]StateDef{}
	for _, state := range wf.GetStates() {
		stateMap[state.State.GetStateId()] = state
	}
	r.workflowStateStore[wfType] = stateMap
	return nil
}

func (r *registry) registerWorkflowCommunicationSchema(wf Workflow) error {
	wfType := r.GetWorkflowType(wf)
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

func (r *registry) registerWorkflowPersistenceSchema(wf Workflow) error {
	wfType := r.GetWorkflowType(wf)
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
