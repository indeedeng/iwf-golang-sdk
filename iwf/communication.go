package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type Communication interface {
	// PublishInterstateChannel publishes a value to an interstate Channel
	PublishInterstateChannel(channelName string, value interface{})

	// below is for internal implementation
	getToPublishInterStateChannel() map[string][]iwfidl.EncodedObject
}
