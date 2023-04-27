package iwf

type CommunicationMethodDef struct {
	Name                string
	CommunicationMethod CommunicationMethod
}

type CommunicationMethod string

const (
	CommunicationMethodSignalChannel   CommunicationMethod = "SignalChannel"
	CommunicationMethodInternalChannel CommunicationMethod = "InternalChannel"
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

