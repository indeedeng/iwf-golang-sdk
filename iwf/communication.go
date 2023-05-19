package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

type Communication interface {
	// PublishInternalChannel publishes a value to an interstate Channel
	PublishInternalChannel(channelName string, value interface{})

	// TriggerStateMovements trigger new state movements as the RPC results
	// NOTE: closing workflows like completing/failing are not supported
	// NOTE: Only used in RPC -- cannot be used in state APIs
	TriggerStateMovements(movements ...StateMovement)

	// below is for internal implementation
	communicationInternal
}

type communicationInternal interface {
	GetToPublishInternalChannel() map[string][]iwfidl.EncodedObject
	GetToTriggerStateMovements() []StateMovement
}
