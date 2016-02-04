package gorma

import (
	"github.com/goadesign/goa/design"

	"github.com/goadesign/goa/dslengine"
)

// RelationalStorageType is the type of database
type RelationalStorageType string

// FieldType is the storage data type for a database field
type FieldType string

// StorageGroupDefinition is the parent configuration structure for Gorma definitions
type StorageGroupDefinition struct {
	dslengine.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	RelationalStores map[string]*RelationalStoreDefinition
}

// RelationalStoreDefinition is the parent configuration structure for Gorm relational model definitions
type RelationalStoreDefinition struct {
	dslengine.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	Parent           *StorageGroupDefinition
	Type             RelationalStorageType
	RelationalModels map[string]*RelationalModelDefinition
	NoAutoIDFields   bool
	NoAutoTimestamps bool
	NoAutoSoftDelete bool
}

// RelationalModelDefinition implements the storage of a domain model into a
// table in a relational database
type RelationalModelDefinition struct {
	dslengine.Definition
	*design.UserTypeDefinition
	DefinitionDSL    func()
	ModelName        string
	Description      string
	GoaType          *design.MediaTypeDefinition
	Parent           *RelationalStoreDefinition
	BuiltFrom        map[string]*design.UserTypeDefinition
	RenderTo         map[string]*design.MediaTypeDefinition
	BelongsTo        map[string]*RelationalModelDefinition
	HasMany          map[string]*RelationalModelDefinition
	HasOne           map[string]*RelationalModelDefinition
	ManyToMany       map[string]*ManyToManyDefinition
	Maps             map[string]*Mapping
	Alias            string // gorm:tablename
	Cached           bool
	CacheDuration    int
	Roler            bool
	DynamicTableName bool
	SQLTag           string
	RelationalFields map[string]*RelationalFieldDefinition
	PrimaryKeys      []*RelationalFieldDefinition
	many2many        []string
}

type Mapping struct {
	dslengine.Definition
	*design.AttributeDefinition
	DefinitionDSL func()
	MappingName   string
	Description   string
	Remote        *design.UserTypeDefinition
	Parent        *RelationalModelDefinition
	Mappings      map[string]*MapDefinition // keyed by gorma field
}

func NewMapping() *Mapping {
	baseAttr := &design.AttributeDefinition{
		Description: "mapping",
	}
	m := &Mapping{
		Mappings:            make(map[string]*MapDefinition),
		AttributeDefinition: baseAttr,
	}
	return m
}

// MapDefinition represents something
type MapDefinition struct {
	RemoteField string
	ParentField string
	GormaType   FieldType //  Override computed field type
}

// RelationalFieldDefinition represents
// a field in a relational database
type RelationalFieldDefinition struct {
	dslengine.Definition
	DefinitionDSL     func()
	Parent            *RelationalModelDefinition
	a                 *design.AttributeDefinition
	Name              string
	Datatype          FieldType
	SQLTag            string
	DatabaseFieldName string
	Description       string
	Nullable          bool
	PrimaryKey        bool
	Timestamp         bool
	Size              int    // string field size
	Alias             string // gorm:column
	BelongsTo         string
	HasOne            string
	HasMany           string
	Many2Many         string
}

// ManyToManyDefinition stores information about a ManyToMany
// relationship between two domain objects
type ManyToManyDefinition struct {
	dslengine.Definition
	DefinitionDSL    func()
	Left             *RelationalModelDefinition
	Right            *RelationalModelDefinition
	RelationshipName string // ??
	DatabaseField    string
}

// StoreIterator is a function that iterates over Relational Stores in a
// StorageGroup
type StoreIterator func(m *RelationalStoreDefinition) error

// ModelIterator is a function that iterates over Models in a
// RelationalStore
type ModelIterator func(m *RelationalModelDefinition) error

// SourceMapIterator
type MapIterator func(m *Mapping) error

// FieldIterator is a function that iterates over Fields
// in a RelationalModel
type FieldIterator func(m *RelationalFieldDefinition) error
