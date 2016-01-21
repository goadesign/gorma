package gorma

import "github.com/goadesign/goa/design"

// RelationalStorageType is the type of database
type RelationalStorageType string

// FieldType is the storage type for a database field
type FieldType string

// StorageGroupDefinition is the parent configuration structure for Gorma definitions
type StorageGroupDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	RelationalStores map[string]*RelationalStoreDefinition
}

// RelationalStoreDefinition is the parent configuration structure for Gorm relational model definitions
type RelationalStoreDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	Parent           *StorageGroupDefinition
	Type             RelationalStorageType
	RelationalModels map[string]*RelationalModelDefinition
}

// RelationalModelDefinition implements the storage of a domain model into a
// table in a relational database
type RelationalModelDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string //
	Parent           *RelationalStoreDefinition
	ModeledType      *design.UserTypeDefinition
	BelongsTo        map[string]*RelationalModelDefinition
	HasMany          map[string]*RelationalModelDefinition
	HasOne           map[string]*RelationalModelDefinition
	ManyToMany       map[string]*ManyToManyDefinition
	Adapters         map[string]func()
	TableName        string
	Alias            string // gorm:tablename
	Cached           bool
	CacheDuration    int
	NoMedia          bool
	Roler            bool
	DynamicTableName bool
	SQLTag           string
	RelationalFields map[string]*RelationalFieldDefinition
	PrimaryKeys      []*RelationalFieldDefinition
	many2many        []string
}

// MediaTypeAdapterDefinition represents the transformation of a
// Goa media type into a Gorma Model
type MediaTypeAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.MediaTypeDefinition
	Right         *RelationalModelDefinition
}

// UserTypeAdapterDefinition represents the transformation of a Goa
// user type into a Gorma Model
type UserTypeAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *RelationalModelDefinition
	Right         *RelationalModelDefinition
}

// PayloadAdapterDefinition represents the transformation of a Goa
// Payload (which is really a UserTypeDefinition
// into a Gorma model
type PayloadAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.UserTypeDefinition
	Right         *RelationalModelDefinition
}

// RelationalFieldDefinition implements the Field interface and represents
// a field in a relational database
type RelationalFieldDefinition struct {
	design.Definition
	DefinitionDSL     func()
	Parent            *RelationalModelDefinition
	a                 *design.AttributeDefinition
	Name              string
	Datatype          FieldType
	SQLTag            string //
	DatabaseFieldName string
	Description       string
	Nullable          bool
	PrimaryKey        bool
	Timestamp         bool
	Size              int // string field size
	Alias             string
	BelongsTo         string
	HasOne            string
	HasMany           string
	Many2Many         string
}

// ManyToManyDefinition stores information about a ManyToMany
// relationship between two domain objects
type ManyToManyDefinition struct {
	design.Definition
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

// FieldIterator is a function that iterates over Fields
// in a RelationalModel
type FieldIterator func(m *RelationalFieldDefinition) error
