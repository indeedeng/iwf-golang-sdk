package internal

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
	"reflect"
)

type registry struct {
	workflowStore              map[string]iwf.Workflow
	workflowStateStore         map[string]map[string]iwf.StateDef
	signalNameStore            map[string]map[string]bool
	interStateChannelNameStore map[string]map[string]bool
	dataObjectKeyStore         map[string]map[string]bool
	searchAttributeTypeStore   map[string]map[string]iwfidl.SearchAttributeValueType
}

func NewRegistry() iwf.Registry {
	return &registry{}
}

func (r *registry) AddWorkflow(wf iwf.Workflow) error {
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
func (r *registry) GetWorkflowType(wf iwf.Workflow) string {
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

func (r *registry) registerWorkflow(wf iwf.Workflow) error {
	wfType := r.GetWorkflowType(wf)
	_, ok := r.workflowStore[wfType]
	if ok {
		return iwf.NewWorkflowDefinitionError("workflow type conflict: " + wfType)
	}
	r.workflowStore[wfType] = wf
	return nil
}

func (r *registry) registerWorkflowState(wf iwf.Workflow) error {
	wfType := r.GetWorkflowType(wf)
	if len(wf.GetStates()) == 0 {
		return iwf.NewWorkflowDefinitionFmtError("Workflow type %s must contain at least one workflow state", wfType)
	}
	var stateMap map[string]iwf.StateDef
	for _, state := range wf.GetStates() {
		stateMap[state.State.GetStateId()] = state
	}
	r.workflowStateStore[wfType] = stateMap
	return nil
}

func (r *registry) registerWorkflowCommunicationSchema(wf iwf.Workflow) error {
	wfType := r.GetWorkflowType(wf)
	var signalMap map[string]bool
	var interStateChannel map[string]bool
	for _, methodDef := range wf.GetCommunicationSchema() {
		if methodDef.CommunicationMethod == iwf.CommunicationMethodSignalChannel {
			signalMap[methodDef.Name] = true
		} else if methodDef.CommunicationMethod == iwf.CommunicationMethodInterstateChannel {
			interStateChannel[methodDef.Name] = true
		} else {
			return iwf.NewWorkflowDefinitionError("invalid CommunicationMethod definition " + string(methodDef.CommunicationMethod))
		}
	}
	r.signalNameStore[wfType] = signalMap
	r.interStateChannelNameStore[wfType] = interStateChannel
	return nil
}

func (r *registry) registerWorkflowPersistenceSchema(wf iwf.Workflow) error {
	wfType := r.GetWorkflowType(wf)
	var dataObjectKeys map[string]bool
	var searchAttributes map[string]iwfidl.SearchAttributeValueType
	for _, pers := range wf.GetPersistenceSchema() {
		if pers.FieldType == iwf.PersistenceFieldTypeDataObject {
			dataObjectKeys[pers.Key] = true
		} else if pers.FieldType == iwf.PersistenceFieldTypeSearchAttribute && pers.SearchAttributeType != nil {
			searchAttributes[pers.Key] = *pers.SearchAttributeType
		} else {
			return iwf.NewWorkflowDefinitionFmtError("invalid PersistenceField definition %s for key %s ", string(pers.FieldType), pers.Key)
		}
	}
	r.dataObjectKeyStore[wfType] = dataObjectKeys
	r.searchAttributeTypeStore[wfType] = searchAttributes
	return nil
}
