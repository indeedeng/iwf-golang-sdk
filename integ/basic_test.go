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

func TestBasicWorkflow(t *testing.T) {
	wfId := "TestBasicWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &basicWorkflow{}, wfId, 10, 1, &iwf.WorkflowOptions{
		WorkflowIdReusePolicy: ptr.Any(iwfidl.REJECT_DUPLICATE),
		WorkflowRetryPolicy: &iwfidl.RetryPolicy{
			InitialIntervalSeconds: iwfidl.PtrInt32(10),
			MaximumAttempts:        iwfidl.PtrInt32(3),
			MaximumIntervalSeconds: iwfidl.PtrInt32(100),
			BackoffCoefficient:     iwfidl.PtrFloat32(3),
		},
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
