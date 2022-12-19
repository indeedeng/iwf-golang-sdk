package iwf

type Communication interface {
	// PublishInterstateChannel publishes a value to an interstate Channel
	PublishInterstateChannel(channelName string, value interface{}) error
}
