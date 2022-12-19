package iwf

type PersistenceFieldDef struct {
	key                 string
	fieldType           PersistenceFieldType
	searchAttributeType SearchAttributeType
}

type PersistenceFieldType string
type SearchAttributeType string

const (
	PersistenceFieldTypeDataObject      PersistenceFieldType = "DataObject"
	PersistenceFieldTypeSearchAttribute PersistenceFieldType = "SearchAttribute"
)

const (
	SearchAttributeTypeKeyword SearchAttributeType = "keyword"
	SearchAttributeTypeInt64   SearchAttributeType = "int64"
)

func NewDataObjectDef(key string) PersistenceFieldDef {
	return PersistenceFieldDef{
		key:       key,
		fieldType: PersistenceFieldTypeDataObject,
	}
}

func NewSearchAttributeDef(key string, saType SearchAttributeType) PersistenceFieldDef {
	return PersistenceFieldDef{
		key:                 key,
		fieldType:           PersistenceFieldTypeSearchAttribute,
		searchAttributeType: saType,
	}
}
