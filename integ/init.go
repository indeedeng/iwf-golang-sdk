package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

var registry = iwf.NewRegistry()
var client = iwf.NewClient(registry, nil)
var workerService = iwf.NewWorkerService(registry, nil)

func init() {
	err := registry.AddWorkflows(
		&basicWorkflow{},
		&proceedOnStateStartFailWorkflow{},
		&timerWorkflow{},
		&signalWorkflow{},
		&interStateWorkflow{},
		&persistenceWorkflow{},
		&forceFailWorkflow{},
		&stateApiFailWorkflow{},
		&stateApiTimeoutWorkflow{},
		&skipWaitUntilWorkflow{},
		skipWaitUntilWorkflow2{}, // test register by struct
		rpcWorkflow{},
		noStateWorkflow{},
		noStartStateWorkflow{},
		executeApiFailRecoveryWorkflow{},
	)
	if err != nil {
		panic(err)
	}
}
