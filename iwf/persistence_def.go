package iwf

import "github.com/iworkflowio/iwf-golang-sdk/gen/iwfidl"

type PersistenceFieldDef struct {
	Key       string
	FieldType PersistenceFieldType
	// SearchAttributeType is optional and only required for PersistenceFieldTypeSearchAttribute
	SearchAttributeType *iwfidl.SearchAttributeValueType
}

type PersistenceFieldType string

const (
	PersistenceFieldTypeDataObject      PersistenceFieldType = "DataObject"
	PersistenceFieldTypeSearchAttribute PersistenceFieldType = "SearchAttribute"
)


func NewDataObjectDef(key string) PersistenceFieldDef {
	return PersistenceFieldDef{
		Key:       key,
		FieldType: PersistenceFieldTypeDataObject,
	}
}

func NewSearchAttributeDef(key string, saType iwfidl.SearchAttributeValueType) PersistenceFieldDef {
	return PersistenceFieldDef{
		Key:                 key,
		FieldType:           PersistenceFieldTypeSearchAttribute,
		SearchAttributeType: &saType,
	}
}
