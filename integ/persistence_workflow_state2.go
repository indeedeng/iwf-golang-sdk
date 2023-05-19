package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type persistenceWorkflowState2 struct {
	iwf.WorkflowStateDefaults
}

const testText = "Hail iWF!"

func (b persistenceWorkflowState2) WaitUntil(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	iv := persistence.GetSearchAttributeInt(testSearchAttributeInt)
	if iv != 1 {
		panic("this value must be 1 because it got set by WaitUntil API")
	}

	var do ExampleDataObjectModel
	persistence.GetDataAttribute(testDataObjectKey, &do)
	dv := persistence.GetSearchAttributeDatetime(testSearchAttributeDatetime)
	bv := persistence.GetSearchAttributeBool(testSearchAttributeBool)
	persistence.SetSearchAttributeDouble(testSearchAttributeDouble, 1.0)
	if dv.Unix() == do.Datetime.Unix() && bv == true {
		persistence.SetSearchAttributeText(testSearchAttributeText, testText)
		return iwf.EmptyCommandRequest(), nil
	}
	panic("the value of datatime or bool search attribute is incorrect")

}

func (b persistenceWorkflowState2) Execute(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	tv := persistence.GetSearchAttributeText(testSearchAttributeText)
	persistence.SetSearchAttributeKeyword(testSearchAttributeKeyword, "iWF")
	if tv == testText {
		return iwf.GracefulCompletingWorkflow, nil
	}
	panic("the value of text search attribute is incorrect")
}
