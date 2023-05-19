package integ

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

type noStateWorkflow struct {
	iwf.WorkflowDefaults
}

func (b noStateWorkflow) GetCommunicationSchema() []iwf.CommunicationMethodDef {
	return []iwf.CommunicationMethodDef{
		iwf.RPCMethodDef(b.TestErrorRPC, nil),
	}
}

func (b noStateWorkflow) TestErrorRPC(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (interface{}, error) {
	return nil, fmt.Errorf("test error")
}
