package integ

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
)

type persistenceWorkflow struct{}

const (
	testDataObjectKey = "test-data-object"

	testSearchAttributeInt      = "CustomIntField"
	testSearchAttributeDatetime = "CustomDatetimeField"
	testSearchAttributeBool     = "CustomBoolField"
	testSearchAttributeDouble   = "CustomDoubleField"
	testSearchAttributeText     = "CustomStringField"
	testSearchAttributeKeyword  = "CustomKeywordField"
)

func (b persistenceWorkflow) GetStates() []iwf.StateDef {
	return []iwf.StateDef{
		iwf.NewStartingState(&persistenceWorkflowState1{}),
		iwf.NewNonStartingState(&persistenceWorkflowState2{}),
	}
}

func (b persistenceWorkflow) GetPersistenceSchema() []iwf.PersistenceFieldDef {
	return []iwf.PersistenceFieldDef{
		iwf.NewDataObjectDef(testDataObjectKey),
		iwf.NewSearchAttributeDef(testSearchAttributeInt, iwfidl.INT),
		iwf.NewSearchAttributeDef(testSearchAttributeDatetime, iwfidl.DATETIME),
		iwf.NewSearchAttributeDef(testSearchAttributeBool, iwfidl.BOOL),
		iwf.NewSearchAttributeDef(testSearchAttributeDouble, iwfidl.DOUBLE),
		iwf.NewSearchAttributeDef(testSearchAttributeText, iwfidl.TEXT),
		iwf.NewSearchAttributeDef(testSearchAttributeKeyword, iwfidl.KEYWORD),
	}
}

func (b persistenceWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return nil
}

func (b persistenceWorkflow) GetWorkflowType() string {
	return ""
}
