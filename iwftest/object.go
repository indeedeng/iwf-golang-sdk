package iwftest

import "github.com/indeedeng/iwf-golang-sdk/iwf"

func NewTestObject(obj interface{}) iwf.Object {
	obj2, err := iwf.GetDefaultObjectEncoder().Encode(obj)
	if err != nil {
		panic(err)
	}
	return iwf.NewObject(obj2, iwf.GetDefaultObjectEncoder())
}

func NewTestObjectWithEncoder(obj interface{}, encoder iwf.ObjectEncoder) iwf.Object {
	obj2, err := encoder.Encode(obj)
	if err != nil {
		panic(err)
	}
	return iwf.NewObject(obj2, encoder)
}
