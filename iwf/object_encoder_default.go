package iwf

import (
	"encoding/json"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
)

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
	if obj == nil {
		return &iwfidl.EncodedObject{}, nil
	}
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
	if encodedObj == nil || resultPtr == nil || encodedObj.GetData() == "" {
		return nil
	}
	return json.Unmarshal([]byte(encodedObj.GetData()), resultPtr)
}
