package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"time"
)

// Persistence APIs are for read/write persistence of workflow
// All API would panic on error but the error --
// if you really need to do some customized handling(mostly you don't need to):
// 1. capturing panic yourself
// 2. get the error from WorkerService API, because WorkerService will use captureStateExecutionError to capture the error
type Persistence interface {
	GetDataAttribute(key string, valuePtr interface{})
	SetDataAttribute(key string, value interface{})

	GetSearchAttributeInt(key string) int64
	SetSearchAttributeInt(key string, value int64)

	GetSearchAttributeKeyword(key string) string
	SetSearchAttributeKeyword(key string, value string)

	GetSearchAttributeBool(key string) bool
	SetSearchAttributeBool(key string, value bool)

	GetSearchAttributeDouble(key string) float64
	SetSearchAttributeDouble(key string, value float64)

	GetSearchAttributeText(key string) string
	SetSearchAttributeText(key string, value string)

	GetSearchAttributeDatetime(key string) time.Time
	SetSearchAttributeDatetime(key string, value time.Time)

	GetSearchAttributeKeywordArray(key string) []string
	SetSearchAttributeKeywordArray(key string, value []string)

	// GetStateExecutionLocal retrieves a local state attribute
	// User code must make sure using the same type for both get and set
	GetStateExecutionLocal(key string, valuePtr interface{})
	// SetStateExecutionLocal sets a local attribute. The scope of the attribute is only within the execution of this state.
	// Usually it's for passing from State Start API to State Decide API
	// User code must make sure using the same type for both get and set
	SetStateExecutionLocal(key string, value interface{})

	// RecordEvent records an arbitrary event in State Start/Decide API for debugging/tracking purpose
	//  Name is the name of the event. Within a Start/Decide API, the same Name cannot be used for more than once.
	//  eventData is the data of the event.
	RecordEvent(key string, value interface{})

	// below is for internal implementation
	persistenceInternal
}

type persistenceInternal interface {
	GetToReturn() (
		dataObjectsToReturn []iwfidl.KeyValue,
		stateLocalToReturn []iwfidl.KeyValue,
		recordEvents []iwfidl.KeyValue,
		searchAttributes []iwfidl.SearchAttribute,
	)
}
