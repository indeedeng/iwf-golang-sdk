package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"time"
)

type Persistence interface {
	GetDataObject(key string, valuePtr interface{}) error
	SetDataObject(key string, value interface{}) error

	GetSearchAttributeInt(key string) (int64, error)
	SetSearchAttributeInt(key string, value int64) error

	GetSearchAttributeKeyword(key string) (string, error)
	SetSearchAttributeKeyword(key string, value string) error

	GetSearchAttributeBool(key string) (bool, error)
	SetSearchAttributeBool(key string, value bool) error

	GetSearchAttributeDouble(key string) (float64, error)
	SetSearchAttributeDouble(key string, value float64) error

	GetSearchAttributeText(key string) (string, error)
	SetSearchAttributeText(key string, value string) error

	GetSearchAttributeDatetime(key string) (time.Time, error)
	SetSearchAttributeDatetime(key string, value time.Time) error

	GetSearchAttributeKeywordArray(key string) ([]string, error)
	SetSearchAttributeKeywordArray(key string, value []string) error

	// GetStateLocal retrieves a local state attribute
	// User code must make sure using the same type for both get and set
	GetStateLocal(key string, valuePtr interface{}) error
	// SetStateLocal sets a local attribute. The scope of the attribute is only within the execution of this state.
	// Usually it's for passing from State Start API to State Decide API
	// User code must make sure using the same type for both get and set
	SetStateLocal(key string, value interface{}) error

	// RecordEvent records an arbitrary event in State Start/Decide API for debugging/tracking purpose
	//  Name is the name of the event. Within a Start/Decide API, the same Name cannot be used for more than once.
	//  eventData is the data of the event.
	RecordEvent(key string, value interface{}) error

	// below is for internal implementation
	getToReturn() (
		dataObjectsToReturn []iwfidl.KeyValue,
		stateLocalToReturn []iwfidl.KeyValue,
		recordEvents []iwfidl.KeyValue,
		searchAttributes []iwfidl.SearchAttribute,
	)
}

type PersistenceX interface {
	GetDataObjectX(key string, valuePtr interface{})
	SetDataObjectX(key string, value interface{})

	GetSearchAttributeIntX(key string) int64
	SetSearchAttributeIntX(key string, value int64)

	GetSearchAttributeKeywordX(key string) string
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

	// GetStateLocal retrieves a local state attribute
	// User code must make sure using the same type for both get and set
	GetStateLocal(key string, valuePtr interface{})
	// SetStateLocal sets a local attribute. The scope of the attribute is only within the execution of this state.
	// Usually it's for passing from State Start API to State Decide API
	// User code must make sure using the same type for both get and set
	SetStateLocal(key string, value interface{})

	// RecordEvent records an arbitrary event in State Start/Decide API for debugging/tracking purpose
	//  Name is the name of the event. Within a Start/Decide API, the same Name cannot be used for more than once.
	//  eventData is the data of the event.
	RecordEvent(key string, value interface{})
}
