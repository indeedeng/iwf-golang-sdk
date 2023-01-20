package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type persistenceWorkflowState1 struct{}

const persistenceWorkflowState1Id = "persistenceWorkflowState1"

func (b persistenceWorkflowState1) GetStateId() string {
	return persistenceWorkflowState1Id
}

func (b persistenceWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	// TODO https://github.com/indeedeng/iwf/issues/146
	//kw, err := persistence.GetSearchAttributeKeyword(testSearchAttributeKeyword)
	//if kw != "init-1" {
	//	panic("must have a init value")
	//}
	//txt, err := persistence.GetSearchAttributeKeyword(testSearchAttributeText)
	//if txt != "init-2" {
	//	panic("must have a init value")
	//}

	var do ExampleDataObjectModel
	err := persistence.GetDataObject(testDataObjectKey, &do)
	if err != nil {
		return nil, err
	}
	if do.StrValue == "" && do.IntValue == 0 {
		err := input.Get(&do)
		if err != nil {
			return nil, err
		}
		if do.StrValue == "" || do.IntValue == 0 {
			panic("this value shouldn't be empty as we got it from start request")
		}
	} else {
		panic("this value should be empty because we haven't set it before")
	}
	err = persistence.SetDataObject(testDataObjectKey, do)
	if err != nil {
		return nil, err
	}
	err = persistence.SetSearchAttributeInt(testSearchAttributeInt, 1)
	if err != nil {
		return nil, err
	}

	return iwf.EmptyCommandRequest(), nil
}

func (b persistenceWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
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
	err = persistence.SetSearchAttributeDatetime(testSearchAttributeDatetime, do.Datetime)
	if err != nil {
		return nil, err
	}
	err = persistence.SetSearchAttributeBool(testSearchAttributeBool, true)
	if err != nil {
		return nil, err
	}
	return iwf.SingleNextState(persistenceWorkflowState2Id, nil), nil
}

func (b persistenceWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return nil
}
