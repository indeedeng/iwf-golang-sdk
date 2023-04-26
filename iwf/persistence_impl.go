package iwf

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf/ptr"
	"time"
)

type persistenceImpl struct {
	encoder ObjectEncoder

	// for data attributes
	dataObjectKeyMap    map[string]bool
	currentDataObjects  map[string]iwfidl.EncodedObject
	dataObjectsToReturn map[string]iwfidl.EncodedObject

	// for state locals
	recordedEvents     map[string]iwfidl.EncodedObject
	currentStateLocal  map[string]iwfidl.EncodedObject
	stateLocalToReturn map[string]iwfidl.EncodedObject

	// for search attributes
	saKeyToType          map[string]iwfidl.SearchAttributeValueType
	saCurrentIntValue    map[string]int64
	saIntToReturn        map[string]int64
	saCurrentStringValue map[string]string
	saStringToReturn     map[string]string
	saCurrentDoubleValue map[string]float64
	saDoubleToReturn     map[string]float64
	saCurrentBoolValue   map[string]bool
	saBoolToReturn       map[string]bool
	saCurrentStrArrValue map[string][]string
	saStrArrToReturn     map[string][]string
}

func newPersistence(
	encoder ObjectEncoder, dataObjectKeyMap map[string]bool, saKeyToType map[string]iwfidl.SearchAttributeValueType,
	dataObjects []iwfidl.KeyValue, searchAttributes []iwfidl.SearchAttribute, stateLocals []iwfidl.KeyValue,
) (Persistence, error) {
	currentDataObjects := make(map[string]iwfidl.EncodedObject)
	currentStateLocal := make(map[string]iwfidl.EncodedObject)
	saCurrentIntValue := make(map[string]int64)
	saCurrentStringValue := make(map[string]string)
	saCurrentDoubleValue := make(map[string]float64)
	saCurrentBoolValue := make(map[string]bool)
	saCurrentStrArrValue := make(map[string][]string)

	for _, do := range dataObjects {
		currentDataObjects[do.GetKey()] = do.GetValue()
	}

	for _, sl := range stateLocals {
		currentStateLocal[sl.GetKey()] = sl.GetValue()
	}

	for _, sa := range searchAttributes {
		switch sa.GetValueType() {
		case iwfidl.KEYWORD, iwfidl.DATETIME, iwfidl.TEXT:
			saCurrentStringValue[sa.GetKey()] = sa.GetStringValue()
		case iwfidl.BOOL:
			saCurrentBoolValue[sa.GetKey()] = sa.GetBoolValue()
		case iwfidl.DOUBLE:
			saCurrentDoubleValue[sa.GetKey()] = sa.GetDoubleValue()
		case iwfidl.KEYWORD_ARRAY:
			saCurrentStrArrValue[sa.GetKey()] = sa.GetStringArrayValue()
		case iwfidl.INT:
			saCurrentIntValue[sa.GetKey()] = sa.GetIntegerValue()
		default:
			return nil, newInternalError("invalid search attribute type %v for key %v ", sa.GetValueType(), sa.GetKey())
		}
	}

	return &persistenceImpl{
		encoder:              encoder,
		dataObjectKeyMap:     dataObjectKeyMap,
		currentDataObjects:   currentDataObjects,
		dataObjectsToReturn:  make(map[string]iwfidl.EncodedObject),
		recordedEvents:       make(map[string]iwfidl.EncodedObject),
		currentStateLocal:    currentStateLocal,
		stateLocalToReturn:   make(map[string]iwfidl.EncodedObject),
		saKeyToType:          saKeyToType,
		saCurrentIntValue:    saCurrentIntValue,
		saIntToReturn:        make(map[string]int64),
		saCurrentStringValue: saCurrentStringValue,
		saStringToReturn:     make(map[string]string),
		saCurrentDoubleValue: saCurrentDoubleValue,
		saDoubleToReturn:     make(map[string]float64),
		saCurrentBoolValue:   saCurrentBoolValue,
		saBoolToReturn:       make(map[string]bool),
		saCurrentStrArrValue: saCurrentStrArrValue,
		saStrArrToReturn:     make(map[string][]string),
	}, nil
}

func (p *persistenceImpl) GetDataAttribute(key string, valuePtr interface{}) {
	if !p.dataObjectKeyMap[key] {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a data object", key))
	}
	err := p.encoder.Decode(ptr.Any(p.currentDataObjects[key]), valuePtr)
	if err != nil {
		panic(err)
	}
}

func (p *persistenceImpl) SetDataAttribute(key string, value interface{}) {
	if !p.dataObjectKeyMap[key] {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a data object", key))
	}
	v, err := p.encoder.Encode(value)
	if err != nil {
		panic(err)
	}
	p.dataObjectsToReturn[key] = *v
	p.currentDataObjects[key] = *v
}

func (p *persistenceImpl) GetSearchAttributeInt(key string) int64 {
	if p.saKeyToType[key] != iwfidl.INT {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a INT search attribute", key))
	}
	return p.saCurrentIntValue[key]
}

func (p *persistenceImpl) SetSearchAttributeInt(key string, value int64) {
	if p.saKeyToType[key] != iwfidl.INT {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a INT search attribute", key))
	}
	p.saCurrentIntValue[key] = value
	p.saIntToReturn[key] = value
}

func (p *persistenceImpl) GetSearchAttributeKeyword(key string) string {
	if p.saKeyToType[key] != iwfidl.KEYWORD {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a KEYWORD search attribute", key))
	}
	return p.saCurrentStringValue[key]
}

func (p *persistenceImpl) SetSearchAttributeKeyword(key string, value string) {
	if p.saKeyToType[key] != iwfidl.KEYWORD {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a KEYWORD search attribute", key))
	}
	p.saCurrentStringValue[key] = value
	p.saStringToReturn[key] = value
}

func (p *persistenceImpl) GetSearchAttributeBool(key string) bool {
	if p.saKeyToType[key] != iwfidl.BOOL {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a BOOL search attribute", key))
	}
	return p.saCurrentBoolValue[key]
}

func (p *persistenceImpl) SetSearchAttributeBool(key string, value bool) {
	if p.saKeyToType[key] != iwfidl.BOOL {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a BOOL search attribute", key))
	}
	p.saBoolToReturn[key] = value
	p.saCurrentBoolValue[key] = value
}

func (p *persistenceImpl) GetSearchAttributeDouble(key string) float64 {
	if p.saKeyToType[key] != iwfidl.DOUBLE {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a DOUBLE search attribute", key))
	}
	return p.saCurrentDoubleValue[key]
}

func (p *persistenceImpl) SetSearchAttributeDouble(key string, value float64) {
	if p.saKeyToType[key] != iwfidl.DOUBLE {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a DOUBLE search attribute", key))
	}
	p.saCurrentDoubleValue[key] = value
	p.saDoubleToReturn[key] = value
}

func (p *persistenceImpl) GetSearchAttributeText(key string) string {
	if p.saKeyToType[key] != iwfidl.TEXT {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a TEXT search attribute", key))
	}
	return p.saCurrentStringValue[key]
}

func (p *persistenceImpl) SetSearchAttributeText(key string, value string) {
	if p.saKeyToType[key] != iwfidl.TEXT {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a TEXT search attribute", key))
	}
	p.saStringToReturn[key] = value
	p.saCurrentStringValue[key] = value
}

func (p *persistenceImpl) GetSearchAttributeDatetime(key string) time.Time {
	if p.saKeyToType[key] != iwfidl.DATETIME {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a DATETIME search attribute", key))
	}
	d, err := time.Parse(DateTimeFormat, p.saCurrentStringValue[key])
	if err != nil {
		panic(err)
	}
	return d
}

func (p *persistenceImpl) SetSearchAttributeDatetime(key string, value time.Time) {
	if p.saKeyToType[key] != iwfidl.DATETIME {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a DATETIME search attribute", key))
	}

	v := value.Format(DateTimeFormat)
	p.saCurrentStringValue[key] = v
	p.saStringToReturn[key] = v
}

func (p *persistenceImpl) GetSearchAttributeKeywordArray(key string) []string {
	if p.saKeyToType[key] != iwfidl.KEYWORD_ARRAY {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a KEYWORD_ARRAY search attribute", key))
	}
	return p.saCurrentStrArrValue[key]
}

func (p *persistenceImpl) SetSearchAttributeKeywordArray(key string, value []string) {
	if p.saKeyToType[key] != iwfidl.KEYWORD_ARRAY {
		panic(NewWorkflowDefinitionErrorFmt("key %v is not registered as a KEYWORD_ARRAY search attribute", key))
	}
	p.saCurrentStrArrValue[key] = value
	p.saStrArrToReturn[key] = value
}

func (p *persistenceImpl) GetStateExecutionLocal(key string, valuePtr interface{}) {
	err := p.encoder.Decode(ptr.Any(p.currentStateLocal[key]), valuePtr)
	if err != nil {
		panic(err)
	}
}

func (p *persistenceImpl) SetStateExecutionLocal(key string, value interface{}) {
	v, err := p.encoder.Encode(value)
	if err != nil {
		panic(err)
	}
	p.currentStateLocal[key] = *v
	p.stateLocalToReturn[key] = *v
}

func (p *persistenceImpl) RecordEvent(key string, value interface{}) {
	v, err := p.encoder.Encode(value)
	if err != nil {
		panic(err)
	}
	p.recordedEvents[key] = *v
}

func (p *persistenceImpl) GetToReturn() (
	dataObjectsToReturn []iwfidl.KeyValue,
	stateLocalToReturn []iwfidl.KeyValue,
	recordEvents []iwfidl.KeyValue,
	searchAttributes []iwfidl.SearchAttribute) {
	for k, v := range p.dataObjectsToReturn {
		dataObjectsToReturn = append(dataObjectsToReturn, iwfidl.KeyValue{
			Key:   ptr.Any(k),
			Value: ptr.Any(v),
		})
	}

	for k, v := range p.stateLocalToReturn {
		stateLocalToReturn = append(stateLocalToReturn, iwfidl.KeyValue{
			Key:   ptr.Any(k),
			Value: ptr.Any(v),
		})
	}

	for k, v := range p.recordedEvents {
		recordEvents = append(recordEvents, iwfidl.KeyValue{
			Key:   ptr.Any(k),
			Value: ptr.Any(v),
		})
	}
	for k, sa := range p.saIntToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:          ptr.Any(k),
			ValueType:    ptr.Any(p.saKeyToType[k]),
			IntegerValue: ptr.Any(sa),
		})
	}
	for k, sa := range p.saStringToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:         ptr.Any(k),
			ValueType:   ptr.Any(p.saKeyToType[k]),
			StringValue: ptr.Any(sa),
		})
	}
	for k, sa := range p.saDoubleToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:         ptr.Any(k),
			ValueType:   ptr.Any(p.saKeyToType[k]),
			DoubleValue: ptr.Any(sa),
		})
	}
	for k, sa := range p.saBoolToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:       ptr.Any(k),
			ValueType: ptr.Any(p.saKeyToType[k]),
			BoolValue: ptr.Any(sa),
		})
	}
	for k, sa := range p.saStrArrToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:              ptr.Any(k),
			ValueType:        ptr.Any(p.saKeyToType[k]),
			StringArrayValue: sa,
		})
	}
	return
}
