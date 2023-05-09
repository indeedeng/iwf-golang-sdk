package integ

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRPCWorkflow(t *testing.T) {
	wfId := "TestRPCWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	wf := rpcWorkflow{}

	runId, err := client.StartWorkflow(context.Background(), wf, wfId, 10, 1, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	time.Sleep(time.Second)
	info, err := client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.RUNNING, info.Status)

	var rpcOutput int
	err = client.InvokeRPC(context.Background(), wfId, "", wf.TestRPC, 1, &rpcOutput)
	assert.Nil(t, err)
	assert.Equal(t, 2, rpcOutput)

	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)
}
