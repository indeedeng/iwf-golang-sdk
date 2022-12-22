package iwf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaptureError(t *testing.T) {
	skipCaptureErrorLogging = true
	assertions := assert.New(t)
	err := testExecuteWithError()
	assertions.NotNilf(err, "cannot be nil")
	err = testExecuteWithSuccess()
	assertions.Nil(err, "must be nil")
	err = testExecuteWithPanic()
	assertions.NotNilf(err, "cannot be nil")
}

func testExecuteWithError() (retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()
	return fmt.Errorf("some error")
}

func testExecuteWithSuccess() (retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()
	return nil
}

func testExecuteWithPanic() (retErr error) {
	defer func() { captureStateExecutionError(recover(), &retErr) }()
	panic("some panic")
}
