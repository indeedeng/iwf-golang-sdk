package integ

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestBasicWorkflow(t *testing.T) {
	wfId := "basic" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &basicWorkflow{}, basicWorkflowState1Id, wfId, 10, 1, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)
}
