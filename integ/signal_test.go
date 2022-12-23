package integ

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSignalWorkflow(t *testing.T) {
	wfId := "TestSignalWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &signalWorkflow{}, signalWorkflowState1Id, wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	err = client.SignalWorkflow(context.Background(), &signalWorkflow{}, wfId, "", testChannelName2, 10)
	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 10, output)
}

func TestSignalWorkflowWithUntypedClient(t *testing.T) {
	client := iwf.NewUntypedClient(nil)

	wfType := iwf.GetDefaultWorkflowType(&signalWorkflow{})
	wfId := "TestSignalWorkflowWithUntypedClient" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), wfType, signalWorkflowState1Id, wfId, 10, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, runId)
	err = client.SignalWorkflow(context.Background(), nil, wfId, "", testChannelName2, 10)
	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 10, output)
}
