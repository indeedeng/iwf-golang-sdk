package integ

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
)

func TestAbnormalExitWorkflow(t *testing.T) {
	wfId := "TestAbnormalExitWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &abnormalExitWorkflow{}, wfId, 10, 1, &iwf.WorkflowOptions{
		WorkflowIdReusePolicy: ptr.Any(iwfidl.ALLOW_IF_PREVIOUS_EXITS_ABNORMALLY),
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	// start the same workflowId again will fail
	_, err = client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, nil, nil)
	assert.True(t, iwf.IsWorkflowAlreadyStartedError(err))

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)

	err = client.GetSimpleWorkflowResult(context.Background(), "a wrong workflowId", "", &output)
	assert.True(t, iwf.IsWorkflowNotExistsError(err))
}
