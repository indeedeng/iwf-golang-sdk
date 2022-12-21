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

func NewSignalChannelDef(channelName string) *CommunicationMethodDef {
	return &CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodSignalChannel,
	}
}

func NewInterstateChannelDef(channelName string) *CommunicationMethodDef {
	return &CommunicationMethodDef{
		Name:                channelName,
		CommunicationMethod: CommunicationMethodInterstateChannel,
	}
}
