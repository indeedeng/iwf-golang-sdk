package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type Communication interface {
	// PublishInterstateChannel publishes a value to an interstate Channel
	PublishInternalChannel(channelName string, value interface{})

	// below is for internal implementation
	communicationInternal
}

type communicationInternal interface {
	GetToPublishInternalChannel() map[string][]iwfidl.EncodedObject
}