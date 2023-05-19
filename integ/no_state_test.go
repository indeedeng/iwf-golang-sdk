package integ

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNoStateWorkflow(t *testing.T) {
	wfId := "TestNoStateWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	wf := noStateWorkflow{}

	runId, err := client.StartWorkflow(context.Background(), wf, wfId, 10, 1, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	time.Sleep(time.Second)
	info, err := client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.RUNNING, info.Status)

	err = client.InvokeRPC(context.Background(), wfId, "", wf.TestErrorRPC, 1, nil)
	assert.NotNil(t, err)
	assert.True(t, iwf.IsRPCError(err))
	rpcErr, _ := err.(*iwf.ApiError)
	assert.Equal(t, "worker API error, status:501, errorType:test-error-type", rpcErr.Response.GetDetail())

	err = client.StopWorkflow(context.Background(), wfId, "", &iwf.WorkflowStopOptions{
		StopType: iwfidl.FAIL,
		Reason:   "test",
	})
	assert.Nil(t, err)
	time.Sleep(time.Second * 2)
	info, err = client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.FAILED, info.Status)
}
