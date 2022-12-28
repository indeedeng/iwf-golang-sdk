package iwf

import (
	"github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"
	"github.com/iworkflowio/iwf-golang-sdk/iwf/ptr"
	"time"
)

type persistenceImpl struct {
	encoder ObjectEncoder

	// for data objects
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

func (p *persistenceImpl) GetDataObject(key string, valuePtr interface{}) error {
	if !p.dataObjectKeyMap[key] {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a data object", key)
	}
	return p.encoder.Decode(ptr.Any(p.currentDataObjects[key]), valuePtr)
}

func (p *persistenceImpl) SetDataObject(key string, value interface{}) error {
	if !p.dataObjectKeyMap[key] {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a data object", key)
	}
	v, err := p.encoder.Encode(value)
	if err != nil {
		return err
	}
	p.dataObjectsToReturn[key] = *v
	p.currentDataObjects[key] = *v
	return nil
}

func (p *persistenceImpl) GetSearchAttributeInt(key string) (int64, error) {
	if p.saKeyToType[key] != iwfidl.INT {
		return 0, NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentIntValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeInt(key string, value int64) error {
	if p.saKeyToType[key] != iwfidl.INT {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saCurrentIntValue[key] = value
	p.saIntToReturn[key] = value
	return nil
}

func (p *persistenceImpl) GetSearchAttributeKeyword(key string) (string, error) {
	if p.saKeyToType[key] != iwfidl.KEYWORD {
		return "", NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentStringValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeKeyword(key string, value string) error {
	if p.saKeyToType[key] != iwfidl.KEYWORD {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saCurrentStringValue[key] = value
	p.saStringToReturn[key] = value
	return nil
}

func (p *persistenceImpl) GetSearchAttributeBool(key string) (bool, error) {
	if p.saKeyToType[key] != iwfidl.BOOL {
		return false, NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentBoolValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeBool(key string, value bool) error {
	if p.saKeyToType[key] != iwfidl.BOOL {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saBoolToReturn[key] = value
	p.saCurrentBoolValue[key] = value
	return nil
}

func (p *persistenceImpl) GetSearchAttributeDouble(key string) (float64, error) {
	if p.saKeyToType[key] != iwfidl.DOUBLE {
		return 0, NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentDoubleValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeDouble(key string, value float64) error {
	if p.saKeyToType[key] != iwfidl.DOUBLE {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saCurrentDoubleValue[key] = value
	p.saDoubleToReturn[key] = value
	return nil
}

func (p *persistenceImpl) GetSearchAttributeText(key string) (string, error) {
	if p.saKeyToType[key] != iwfidl.TEXT {
		return "", NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentStringValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeText(key string, value string) error {
	if p.saKeyToType[key] != iwfidl.TEXT {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saStringToReturn[key] = value
	p.saCurrentStringValue[key] = value
	return nil
}

func (p *persistenceImpl) GetSearchAttributeDatetime(key string) (time.Time, error) {
	if p.saKeyToType[key] != iwfidl.DATETIME {
		return time.Time{}, NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return time.Parse(DateTimeFormat, p.saCurrentStringValue[key])
}

func (p *persistenceImpl) SetSearchAttributeDatetime(key string, value time.Time) error {
	if p.saKeyToType[key] != iwfidl.DATETIME {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}

	v := value.Format(DateTimeFormat)
	p.saCurrentStringValue[key] = v
	p.saStringToReturn[key] = v
	return nil
}

func (p *persistenceImpl) GetSearchAttributeKeywordArray(key string) ([]string, error) {
	if p.saKeyToType[key] != iwfidl.KEYWORD_ARRAY {
		return nil, NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	return p.saCurrentStrArrValue[key], nil
}

func (p *persistenceImpl) SetSearchAttributeKeywordArray(key string, value []string) error {
	if p.saKeyToType[key] != iwfidl.KEYWORD_ARRAY {
		return NewWorkflowDefinitionFmtError("key %v is not registered as a INT search attribute", key)
	}
	p.saCurrentStrArrValue[key] = value
	p.saStrArrToReturn[key] = value
	return nil
}

func (p *persistenceImpl) GetStateLocal(key string, valuePtr interface{}) error {
	return p.encoder.Decode(ptr.Any(p.currentStateLocal[key]), valuePtr)
}

func (p *persistenceImpl) SetStateLocal(key string, value interface{}) error {
	v, err := p.encoder.Encode(value)
	if err != nil {
		return err
	}
	p.currentStateLocal[key] = *v
	p.stateLocalToReturn[key] = *v
	return nil
}

func (p *persistenceImpl) RecordEvent(key string, value interface{}) error {
	v, err := p.encoder.Encode(value)
	if err != nil {
		return err
	}
	p.recordedEvents[key] = *v
	return nil
}

func (p *persistenceImpl) getToReturn() (
	dataObjectsToReturn []iwfidl.KeyValue,
	stateLocalToReturn []iwfidl.KeyValue,
	recordEvents []iwfidl.KeyValue,
	searchAttributes []iwfidl.SearchAttribute) {
	for k, v := range p.dataObjectsToReturn {
		dataObjectsToReturn = append(dataObjectsToReturn, iwfidl.KeyValue{
			Key:   &k,
			Value: &v,
		})
	}

	for k, v := range p.stateLocalToReturn {
		stateLocalToReturn = append(stateLocalToReturn, iwfidl.KeyValue{
			Key:   &k,
			Value: &v,
		})
	}

	for k, v := range p.recordedEvents {
		recordEvents = append(recordEvents, iwfidl.KeyValue{
			Key:   &k,
			Value: &v,
		})
	}
	for k, sa := range p.saIntToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:          &k,
			ValueType:    ptr.Any(p.saKeyToType[k]),
			IntegerValue: &sa,
		})
	}
	for k, sa := range p.saStringToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:         &k,
			ValueType:   ptr.Any(p.saKeyToType[k]),
			StringValue: &sa,
		})
	}
	for k, sa := range p.saDoubleToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:         &k,
			ValueType:   ptr.Any(p.saKeyToType[k]),
			DoubleValue: &sa,
		})
	}
	for k, sa := range p.saBoolToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:       &k,
			ValueType: ptr.Any(p.saKeyToType[k]),
			BoolValue: &sa,
		})
	}
	for k, sa := range p.saStrArrToReturn {
		searchAttributes = append(searchAttributes, iwfidl.SearchAttribute{
			Key:              &k,
			ValueType:        ptr.Any(p.saKeyToType[k]),
			StringArrayValue: sa,
		})
	}
	return
}
