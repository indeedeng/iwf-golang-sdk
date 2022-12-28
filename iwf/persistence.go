package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
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
		dataObjectsToReturn map[string]iwfidl.EncodedObject,
		stateLocalToReturn map[string]iwfidl.EncodedObject,
		saIntToReturn map[string]int64,
		saStringToReturn map[string]string,
		saDoubleToReturn map[string]float64,
		saBoolToReturn map[string]bool,
		saStrArrToReturn map[string][]string,
	)
}
