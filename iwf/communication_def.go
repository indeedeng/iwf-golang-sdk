package iwf

type CommunicationMethodDef struct {
	Name                string
	CommunicationMethod CommunicationMethod
}

type CommunicationMethod string

const (
	CommunicationMethodSignalChannel     CommunicationMethod = "SignalChannel"
	CommunicationMethodInterstateChannel CommunicationMethod = "InterstateChannel"
)

func SignalChannelDef(channelName string) CommunicationMethodDef {
	return NewSignalChannelDef(channelName)
}

// Deprecated: use SignalChannelDef instead to be more concise and readable
func NewSignalChannelDef(channelName string) CommunicationMethodDef {
	return CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodSignalChannel,
	}
}

func InterstateChannelDef(channelName string) CommunicationMethodDef {
	return NewInterstateChannelDef(channelName)
}

// Deprecated: use InterstateChannelDef instead to be more concise and readable
func NewInterstateChannelDef(channelName string) CommunicationMethodDef {
	return CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodInterstateChannel,
	}
}

