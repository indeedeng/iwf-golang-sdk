package integ

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestWorkflowTimeout(t *testing.T) {
	wfId := "TestWorkflowTimeout" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &signalWorkflow{}, wfId, 1, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)

	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.TIMEOUT, nil, nil, nil, iwf.GetDefaultObjectEncoder()), wErr)

	out, err2 := client.GetComplexWorkflowResults(context.Background(), wfId, "")
	assert.Nil(t, out)
	assert.Equal(t, err, err2)

	assert.Equal(t, "workflow is not completed succesfully, closedStatus: TIMEOUT, failedErrorType(applies if failed as closedStatus):<nil>, error message:<nil>", err.Error())
}

func TestWorkflowCancel(t *testing.T) {
	wfId := "TestWorkflowCancel" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &signalWorkflow{}, wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.StopWorkflow(context.Background(), wfId, "", nil)
	assert.Nil(t, err)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)

	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.CANCELED, nil, nil, nil, iwf.GetDefaultObjectEncoder()), wErr)

	out, err2 := client.GetComplexWorkflowResults(context.Background(), wfId, "")
	assert.Nil(t, out)
	assert.Equal(t, err, err2)

	assert.Equal(t, "workflow is not completed succesfully, closedStatus: CANCELED, failedErrorType(applies if failed as closedStatus):<nil>, error message:<nil>", err.Error())
}

func TestForceFailWorkflow(t *testing.T) {
	wfId := "TestForceFailWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &forceFailWorkflow{}, wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)

	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.FAILED, ptr.Any(iwfidl.STATE_DECISION_FAILING_WORKFLOW_ERROR_TYPE), nil, wErr.StateResults, iwf.GetDefaultObjectEncoder()), wErr)

	out, err2 := client.GetComplexWorkflowResults(context.Background(), wfId, "")
	assert.Nil(t, out)
	assert.Equal(t, err, err2)
	assert.Equal(t, "workflow is not completed succesfully, closedStatus: FAILED, failedErrorType(applies if failed as closedStatus):STATE_DECISION_FAILING_WORKFLOW_ERROR_TYPE, error message:<nil>", err.Error())

	var output string
	err = wErr.GetStateResult(0, &output)
	assert.Nil(t, err)
	assert.Equal(t, "a failing message", output)
}
