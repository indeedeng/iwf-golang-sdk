package iwf

import (
	"encoding/json"
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
)

type ObjectEncoder interface {
	// GetEncodingType returns the encoding info that it can handle
	GetEncodingType() string
	// Encode serialize an object
	Encode(obj interface{}) (*iwfidl.EncodedObject, error)
	// Decode deserialize an object
	Decode(encodedObj *iwfidl.EncodedObject, resultPtr interface{}) error
}

func GetDefaultObjectEncoder() ObjectEncoder {
	return &builtinJsonEncoder{}
}

type builtinJsonEncoder struct {
}

const encodingType = "builtinGolangJson"

func (b *builtinJsonEncoder) GetEncodingType() string {
	return encodingType
}

func (b *builtinJsonEncoder) Encode(obj interface{}) (*iwfidl.EncodedObject, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return &iwfidl.EncodedObject{
		Encoding: ptr.Any(encodingType),
		Data:     ptr.Any(string(data)),
	}, nil
}

func (b *builtinJsonEncoder) Decode(encodedObj *iwfidl.EncodedObject, resultPtr interface{}) error {
	if encodedObj == nil {
		return nil
	}
	return json.Unmarshal([]byte(encodedObj.GetData()), resultPtr)
}
