package integ

import (
	"github.com/iworkflowio/iwf-golang-sdk/iwf"
)

var registry = iwf.NewRegistry()
var client = iwf.NewClient(registry, nil)
var workerService = iwf.NewWorkerService(registry, nil)

func init() {
	err := registry.AddWorkflow(&basicWorkflow{})
	if err != nil {
		panic(err)
	}
}
