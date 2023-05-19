package integ

import (
	"context"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNoStartStateWorkflow(t *testing.T) {
	wfId := "TestNoStartStateWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	wf := noStartStateWorkflow{}

	runId, err := client.StartWorkflow(context.Background(), wf, wfId, 10, 1, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)

	var rpcOutput int
	err = client.InvokeRPC(context.Background(), wfId, "", wf.TestRPC, 1, &rpcOutput)
	assert.Nil(t, err)
	assert.Equal(t, 2, rpcOutput)

	time.Sleep(time.Second * 2)
	info, err := client.DescribeWorkflow(context.Background(), wfId, "")
	assert.Nil(t, err)
	assert.Equal(t, iwfidl.COMPLETED, info.Status)
}
