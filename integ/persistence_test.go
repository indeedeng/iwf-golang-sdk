package integ

import (
	"context"
	"fmt"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestPersistenceWorkflow(t *testing.T) {
	wfId := "TestPersistenceWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	input := ExampleDataObjectModel{
		IntValue: time.Now().UnixNano(),
		StrValue: wfId,
		Datetime: time.Now(),
	}
	runId, err := client.StartWorkflow(context.Background(), &persistenceWorkflow{}, persistenceWorkflowState1Id, wfId, 10, input, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)
	assert.Nil(t, err)
	info, err := client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.COMPLETED, info.Status)
	dos, err := client.GetWorkflowDataObjects(context.Background(), &persistenceWorkflow{}, wfId, "", []string{
		testDataObjectKey,
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(dos))
	var do ExampleDataObjectModel
	dos[testDataObjectKey].Get(&do)
	assert.Equal(t, wfId, do.StrValue)

	sas, err := client.GetWorkflowSearchAttributes(context.Background(), &persistenceWorkflow{}, wfId, "", []string{
		testSearchAttributeKeyword,
		testSearchAttributeText,
		testSearchAttributeBool,
		testSearchAttributeDatetime,
		testSearchAttributeInt,
		testSearchAttributeDouble,
	})
	assert.Nil(t, err)
	expectedSas := map[string]interface{}{
		testSearchAttributeKeyword:  "iWF",
		testSearchAttributeText:     "Hail iWF!",
		testSearchAttributeBool:     true,
		testSearchAttributeDatetime: sas[testSearchAttributeDatetime], // skip this one
		testSearchAttributeInt:      int64(1),
		testSearchAttributeDouble:   1.0,
	}
	assert.Equal(t, expectedSas, sas)

	resp, err := client.SearchWorkflow(context.Background(), iwfidl.WorkflowSearchRequest{
		Query:         fmt.Sprintf("IwfWorkflowType='%v'", iwf.GetDefaultWorkflowType(&persistenceWorkflow{})),
		PageSize:      iwfidl.PtrInt32(1),
		NextPageToken: nil,
	})
	assert.Nil(t, err, iwf.GetOpenApiErrorDetailedMessage(err))
	assert.True(t, len(resp.WorkflowExecutions) > 0)
}
