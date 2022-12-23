package integ

import (
	"context"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestBasicWorkflow(t *testing.T) {
	wfId := "TestBasicWorkflow" + strconv.Itoa(int(time.Now().Unix()))
	runId, err := client.StartWorkflow(context.Background(), &basicWorkflow{}, basicWorkflowState1Id, wfId, 10, 1, &iwf.WorkflowOptions{
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
	var output int
	err = client.GetSimpleWorkflowResult(context.Background(), wfId, "", &output)
	assert.Nil(t, err)
	assert.Equal(t, 3, output)
}
