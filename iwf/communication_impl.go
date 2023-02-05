package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type communicationImpl struct {
	interStateChannelNames     map[string]bool
	toPublishInterStateChannel map[string][]iwfidl.EncodedObject
	encoder                    ObjectEncoder
}

func (c *communicationImpl) GetToPublishInterStateChannel() map[string][]iwfidl.EncodedObject {
	return c.toPublishInterStateChannel
}

func (c *communicationImpl) PublishInterstateChannel(channelName string, value interface{}) {
	if !c.interStateChannelNames[channelName] {
		panic(NewWorkflowDefinitionErrorFmt("channelName %v is not registered", channelName))
	}
	l := c.toPublishInterStateChannel[channelName]
	obj, err := c.encoder.Encode(value)
	if err != nil {
		panic(err)
	}
	l = append(l, *obj)
	c.toPublishInterStateChannel[channelName] = l
}

func newCommunication(encoder ObjectEncoder, interStateChannelNames map[string]bool) Communication {
	return &communicationImpl{
		encoder:                    encoder,
		interStateChannelNames:     interStateChannelNames,
		toPublishInterStateChannel: make(map[string][]iwfidl.EncodedObject),
	}
}
