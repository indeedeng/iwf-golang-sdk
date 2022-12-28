package integ

import (
	"context"
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
}
