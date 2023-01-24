package iwf

import "github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"

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
// It just panics on error but the error can still be accessible if really need to do some customized handling(mostly you don't need to):
// 1. capturing panic yourself
// 2. get the error from WorkerService API, because WorkerService will use captureStateExecutionError to capture the error
func (o Object) Get(resultPtr interface{}) {
	err := o.ObjectEncoder.Decode(o.EncodedObject, resultPtr)
	if err != nil {
		panic(err)
	}
}

