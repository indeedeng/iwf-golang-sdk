package integ

import (
	"context"
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
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

func TestStateApiFailWorkflow(t *testing.T) {
	wfId := "TestStateApiFailWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &stateApiFailWorkflow{}, wfId, 10, nil, &iwf.WorkflowOptions{})
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)

	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.FAILED, ptr.Any(iwfidl.STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE), wErr.ErrorMessage, nil, iwf.GetDefaultObjectEncoder()), wErr)

	assert.True(t, strings.Contains(*wErr.ErrorMessage, "test api failing"), "must contain api failing message")

	out, err2 := client.GetComplexWorkflowResults(context.Background(), wfId, "")
	assert.Nil(t, out)
	assert.Equal(t, err, err2)

	assert.True(t, strings.Contains(err.Error(), "workflow is not completed succesfully, closedStatus: FAILED, failedErrorType(applies if failed as closedStatus):STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE, error message:statusCode: 400, responseBody: {\"error\":\"error message:test api failing"))
}

func TestStateApiTimeoutWorkflow(t *testing.T) {
	wfId := "TestStateApiTimeoutWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &stateApiTimeoutWorkflow{}, wfId, 10, nil, &iwf.WorkflowOptions{})
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", nil)

	wErr, ok := iwf.AsWorkflowUncompletedError(err)
	assert.True(t, ok)
	assert.Equal(t, iwf.NewWorkflowUncompletedError(runId, iwfidl.FAILED, ptr.Any(iwfidl.STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE), wErr.ErrorMessage, nil, iwf.GetDefaultObjectEncoder()), wErr)

	fmt.Println(err)

	expectedMsg := "workflow is not completed succesfully, closedStatus: FAILED, failedErrorType(applies if failed as closedStatus):STATE_API_FAIL_MAX_OUT_RETRY_ERROR_TYPE, error message:activity error (type: StateApiWaitUntil, scheduledEventID: 9, startedEventID: 10, identity: ): activity StartToClose timeout (type: StartToClose)"
	assert.Equal(t, expectedMsg, err.Error())

	out, err2 := client.GetComplexWorkflowResults(context.Background(), wfId, "")
	assert.Nil(t, out)
	assert.Equal(t, err, err2)
}

// TODO need to support terminate operation in Stop API first
//func TestWorkflowTerminated(t *testing.T) {
//
//}
