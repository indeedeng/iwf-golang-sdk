package iwf

type CommunicationMethodDef struct {
	Name                string // for signal and internal channel
	CommunicationMethod CommunicationMethod
	RPC                 RPC         // only for CommunicationMethodRPCMethod
	RPCOptions          *RPCOptions // only for CommunicationMethodRPCMethod
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
		CommunicationMethod: CommunicationMethodRPCMethod,
		RPC:                 rpc,
		RPCOptions:          rpcOptions,
	}
}
