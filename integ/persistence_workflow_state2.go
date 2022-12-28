package integ

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
)

type persistenceWorkflowState2 struct{}

const persistenceWorkflowState2Id = "persistenceWorkflowState2"

const testText = "Hail iWF!"

func (b persistenceWorkflowState2) GetStateId() string {
	return persistenceWorkflowState2Id
}

func (b persistenceWorkflowState2) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	iv, err := persistence.GetSearchAttributeInt(testSearchAttributeInt)
	if err != nil {
		return nil, err
	}
	if iv != 1 {
		panic("this value must be 1 because it got set by Start API")
	}

	var do ExampleDataObjectModel
	err = persistence.GetDataObject(testDataObjectKey, &do)
	if err != nil {
		return nil, err
	}
	dv, err := persistence.GetSearchAttributeDatetime(testSearchAttributeDatetime)
	if err != nil {
		return nil, err
	}
	bv, err := persistence.GetSearchAttributeBool(testSearchAttributeBool)
	if err != nil {
		return nil, err
	}
	if dv.Unix() == do.Datetime.Unix() && bv == true {
		err := persistence.SetSearchAttributeText(testSearchAttributeText, testText)
		if err != nil {
			return nil, err
		}
		return iwf.EmptyCommandRequest(), nil
	}
	panic("the value of datatime or bool search attribute is incorrect")
}

func (b persistenceWorkflowState2) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	tv, err := persistence.GetSearchAttributeText(testSearchAttributeText)
	if err != nil {
		return nil, err
	}
	err = persistence.SetSearchAttributeKeyword(testSearchAttributeKeyword, "iWF")
	if err != nil {
		return nil, err
	}
	if tv == testText {
		return iwf.GracefulCompletingWorkflow, nil
	}
	panic("the value of text search attribute is incorrect")
}

func (b persistenceWorkflowState2) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
