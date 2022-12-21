package iwf

type Persistence interface {
	GetDataObject(key string, valuePtr interface{}) error
	SetDataObject(key string, value interface{}) error

	GetSearchAttributeInt64(key string) (int64, error)
	SetSearchAttributeInt64(key string, value int64) error

	GetSearchAttributeKeyword(key string) (string, error)
	SetSearchAttributeKeyword(key string, value string) error

	// GetStateLocal retrieves a local state attribute
	// User code must make sure using the same type for both get and set
	GetStateLocal(key string, valuePtr interface{}) error
	// SetStateLocal sets a local attribute. The scope of the attribute is only within the execution of this state.
	// Usually it's for passing from State Start API to State Decide API
	// User code must make sure using the same type for both get and set
	SetStateLocal(key string, value interface{}) error

	// RecordEvent records an arbitrary event in State Start/Decide API for debugging/tracking purpose
	// @param Name       the Name of the event. Within a Start/Decide API, the same Name cannot be used for more than once.
	// @param eventData the data of the event.
	RecordEvent(key string, value interface{}) error
}
