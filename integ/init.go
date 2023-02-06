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
		&timerWorkflow{},
		&signalWorkflow{},
		&interStateWorkflow{},
		&persistenceWorkflow{},
		&forceFailWorkflow{},
	)
	if err != nil {
		panic(err)
	}
}
