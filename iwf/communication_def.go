package iwf

type CommunicationMethodDef struct {
	Name                string
	CommunicationMethod CommunicationMethod
	RPC                 RPC
	RPCOptions          *RPCOptions
}

type CommunicationMethod string

const (
	CommunicationMethodSignalChannel   CommunicationMethod = "SignalChannel"
	CommunicationMethodInternalChannel CommunicationMethod = "InternalChannel"
	CommunicationMethodRPCMethod       CommunicationMethod = "RPCMethod"
)

func SignalChannelDef(channelName string) CommunicationMethodDef {
	return CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodSignalChannel,
	}
}

func InternalChannelDef(channelName string) CommunicationMethodDef {
	return CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodInternalChannel,
	}
}

func RPCMethodDef(rpc RPC, rpcOptions *RPCOptions) CommunicationMethodDef {
	return CommunicationMethodDef{
		Name:                GetFinalRPCMethodName(rpc, rpcOptions),
		CommunicationMethod: CommunicationMethodRPCMethod,
		RPC:                 rpc,
		RPCOptions:          rpcOptions,
	}
}

