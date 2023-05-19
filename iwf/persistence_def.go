package iwf

import (
	"fmt"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05-07:00"

type PersistenceFieldDef struct {
	Key       string
	FieldType PersistenceFieldType
	// SearchAttributeType is optional and only required for PersistenceFieldTypeSearchAttribute
	SearchAttributeType *iwfidl.SearchAttributeValueType
}

type PersistenceFieldType string

const (
	PersistenceFieldTypeDataObject      PersistenceFieldType = "DataAttribute"
	PersistenceFieldTypeSearchAttribute PersistenceFieldType = "SearchAttribute"
)

func DataAttributeDef(key string) PersistenceFieldDef {
	return PersistenceFieldDef{
		Key:       key,
		FieldType: PersistenceFieldTypeDataObject,
	}
}
func SearchAttributeDef(key string, saType iwfidl.SearchAttributeValueType) PersistenceFieldDef {
	return PersistenceFieldDef{
		Key:                 key,
		FieldType:           PersistenceFieldTypeSearchAttribute,
		SearchAttributeType: &saType,
	}
}

func getSearchAttributeValue(sa iwfidl.SearchAttribute) (interface{}, error) {
	switch *sa.ValueType {
	case iwfidl.TEXT, iwfidl.KEYWORD:
		return *sa.StringValue, nil
	case iwfidl.KEYWORD_ARRAY:
		return sa.StringArrayValue, nil
	case iwfidl.DOUBLE:
		return *sa.DoubleValue, nil
	case iwfidl.BOOL:
		return *sa.BoolValue, nil
	case iwfidl.DATETIME:
		t, err := time.Parse(DateTimeFormat, *sa.StringValue)
		if err != nil {
			return nil, err
		}
		return t, nil
	case iwfidl.INT:
		return *sa.IntegerValue, nil
	default:
		return nil, fmt.Errorf("unsupported search attribute type %v", sa.GetValueType())
	}
}
