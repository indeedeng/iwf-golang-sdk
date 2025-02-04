package integ

import (
	"context"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
)

func TestAbnormalExitWorkflow(t *testing.T) {
	wfId := "TestAbnormalExitWorkflow" + strconv.Itoa(int(time.Now().Unix()))

	opt := iwf.WorkflowOptions{
		WorkflowIdReusePolicy: ptr.Any(iwfidl.ALLOW_IF_PREVIOUS_EXITS_ABNORMALLY),
	}

	runId, err := client.StartWorkflow(context.Background(), &abnormalExitWorkflow{}, wfId, 10, nil, &opt)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)
	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.True(t, strings.Contains(*wErr.ErrorMessage, "abnormal exit state"))
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.FAILED, ptr.Any(iwfidl.STATE_API_FAIL_ERROR_TYPE), wErr.ErrorMessage, wErr.StateResults, iwf.GetDefaultObjectEncoder()), wErr)

	// Starting a workflow with the same ID should be allowed since the previous failed abnormally
	_, err = client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, 1, &opt)
	assert.False(t, iwf.IsWorkflowAlreadyStartedError(err))

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)
}
