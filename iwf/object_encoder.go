package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
)

type ObjectEncoder interface {
	// GetEncodingType returns the encoding info that it can handle
	GetEncodingType() string
	// Encode serialize an object
	Encode(obj interface{}) (*iwfidl.EncodedObject, error)
	// Decode deserialize an object
	Decode(encodedObj *iwfidl.EncodedObject, resultPtr interface{}) error
}
