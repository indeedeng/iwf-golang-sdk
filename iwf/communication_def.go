package iwf

type CommunicationMethodDef struct {
	name                string
	communicationMethod CommunicationMethod
}

type CommunicationMethod string

const (
	CommunicationMethodSignalChannel     CommunicationMethod = "SignalChannel"
	CommunicationMethodInterstateChannel CommunicationMethod = "InterstateChannel"
)

func NewSignalChannelDef(channelName string) *CommunicationMethodDef {
	return &CommunicationMethodDef{
		name:                channelName,
		communicationMethod: CommunicationMethodSignalChannel,
	}
}

func NewInterstateChannelDef(channelName string) *CommunicationMethodDef {
	return &CommunicationMethodDef{
		name:                channelName,
		communicationMethod: CommunicationMethodInterstateChannel,
	}
}
