package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

// Object is a representation of EncodedObject
type Object struct {
	EncodedObject *iwfidl.EncodedObject
	ObjectEncoder ObjectEncoder
}

func NewObject(EncodedObject *iwfidl.EncodedObject, ObjectEncoder ObjectEncoder) Object {
	return Object{
		EncodedObject: EncodedObject,
		ObjectEncoder: ObjectEncoder,
	}
}

// Get retrieves the actual object
func (o Object) Get(resultPtr interface{}) error {
	return o.ObjectEncoder.Decode(o.EncodedObject, resultPtr)
}
