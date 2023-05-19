package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type communicationImpl struct {
	internalChannelNames     map[string]bool
	toPublishInternalChannel map[string][]iwfidl.EncodedObject
	stateMovements           []StateMovement
	encoder                  ObjectEncoder
}

func (c *communicationImpl) GetToTriggerStateMovements() []StateMovement {
	return c.stateMovements
}

func (c *communicationImpl) TriggerStateMovements(movements ...StateMovement) {
	c.stateMovements = append(c.stateMovements, movements...)
}

func (c *communicationImpl) GetToPublishInternalChannel() map[string][]iwfidl.EncodedObject {
	return c.toPublishInternalChannel
}

func (c *communicationImpl) PublishInternalChannel(channelName string, value interface{}) {
	if !c.internalChannelNames[channelName] {
		panic(NewWorkflowDefinitionErrorFmt("channelName %v is not registered", channelName))
	}
	l := c.toPublishInternalChannel[channelName]
	obj, err := c.encoder.Encode(value)
	if err != nil {
		panic(err)
	}
	l = append(l, *obj)
	c.toPublishInternalChannel[channelName] = l
}

func newCommunication(encoder ObjectEncoder, internalChannelNames map[string]bool) Communication {
	return &communicationImpl{
		encoder:                  encoder,
		internalChannelNames:     internalChannelNames,
		toPublishInternalChannel: make(map[string][]iwfidl.EncodedObject),
	}
}
