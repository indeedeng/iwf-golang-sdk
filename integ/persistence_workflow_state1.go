package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type persistenceWorkflowState1 struct {
	iwf.DefaultStateIdAndOptions
}

func (b persistenceWorkflowState1) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	kw := persistence.GetSearchAttributeKeyword(testSearchAttributeKeyword)

	if kw != "init-keyword" {
		panic("incorrect init value: " + kw)
	}
	txt := persistence.GetSearchAttributeText(testSearchAttributeText)
	if txt != "init-text" {
		panic("incorrect init value: " + txt)
	}

	var do ExampleDataObjectModel
	persistence.GetDataAttribute(testDataObjectKey, &do)
	if do.StrValue == "" && do.IntValue == 0 {
		input.Get(&do)
		if do.StrValue == "" || do.IntValue == 0 {
			panic("this value shouldn't be empty as we got it from start request")
		}
	} else {
		panic("this value should be empty because we haven't set it before")
	}
	persistence.SetDataAttribute(testDataObjectKey, do)
	persistence.SetDataAttribute(testDataObjectKey2, "a string")
	persistence.SetSearchAttributeInt(testSearchAttributeInt, 1)

	return iwf.EmptyCommandRequest(), nil
}

func (b persistenceWorkflowState1) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	iv := persistence.GetSearchAttributeInt(testSearchAttributeInt)
	if iv != 1 {
		panic("this value must be 1 because it got set by WaitUntil API")
	}

	var do ExampleDataObjectModel
	persistence.GetDataAttribute(testDataObjectKey, &do)
	var str string
	persistence.GetDataAttribute(testDataObjectKey2, &str)
	if str != "a string" {
		panic("testDataObjectKey2 value is incorrect")
	}

	persistence.SetSearchAttributeDatetime(testSearchAttributeDatetime, do.Datetime)
	persistence.SetSearchAttributeBool(testSearchAttributeBool, true)
	return iwf.SingleNextState(persistenceWorkflowState2{}, nil), nil
}
