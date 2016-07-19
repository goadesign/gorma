package gorma

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/dslengine"
)

// RelationalStorageType is the type of database.
type RelationalStorageType string

// FieldType is the storage data type for a database field.
type FieldType string

// StorageGroupDefinition is the parent configuration structure for Gorma definitions.
type StorageGroupDefinition struct {
	dslengine.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	RelationalStores map[string]*RelationalStoreDefinition
}

// RelationalStoreDefinition is the parent configuration structure for Gorm relational model definitions.
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
// table in a relational database.
type RelationalModelDefinition struct {
	dslengine.Definition
	*design.UserTypeDefinition
	DefinitionDSL    func()
	ModelName        string
	Description      string
	GoaType          *design.MediaTypeDefinition
	Parent           *RelationalStoreDefinition
	BuiltFrom        map[string]*design.UserTypeDefinition
	BuildSources     []*BuildSource
	RenderTo         map[string]*design.MediaTypeDefinition
	BelongsTo        map[string]*RelationalModelDefinition
	HasMany          map[string]*RelationalModelDefinition
	HasOne           map[string]*RelationalModelDefinition
	ManyToMany       map[string]*ManyToManyDefinition
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

// BuildSource stores the BuildsFrom sources
// for parsing.
type BuildSource struct {
	dslengine.Definition
	DefinitionDSL   func()
	Parent          *RelationalModelDefinition
	BuildSourceName string
}

// MapDefinition represents field mapping to and from
// Gorma models.
type MapDefinition struct {
	RemoteType  *design.UserTypeDefinition
	RemoteField string
}

// MediaTypeAdapterDefinition represents the transformation of a
// Goa media type into a Gorma Model.
//
// Unimplemented at this time.
type MediaTypeAdapterDefinition struct {
	dslengine.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.MediaTypeDefinition
	Right         *RelationalModelDefinition
}

// UserTypeAdapterDefinition represents the transformation of a Goa
// user type into a Gorma Model.
//
// Unimplemented at this time.
type UserTypeAdapterDefinition struct {
	dslengine.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *RelationalModelDefinition
	Right         *RelationalModelDefinition
}

// PayloadAdapterDefinition represents the transformation of a Goa
// Payload (which is really a UserTypeDefinition)
// into a Gorma model.
//
// Unimplemented at this time.
type PayloadAdapterDefinition struct {
	dslengine.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.UserTypeDefinition
	Right         *RelationalModelDefinition
}

// RelationalFieldDefinition represents
// a field in a relational database.
type RelationalFieldDefinition struct {
	dslengine.Definition
	DefinitionDSL     func()
	Parent            *RelationalModelDefinition
	a                 *design.AttributeDefinition
	FieldName         string
	TableName         string
	Datatype          FieldType
	SQLTag            string
	DatabaseFieldName string // gorm:column
	Description       string
	Nullable          bool
	PrimaryKey        bool
	Timestamp         bool
	Size              int // string field size
	BelongsTo         string
	HasOne            string
	HasMany           string
	Many2Many         string
	Mappings          map[string]*MapDefinition
}

// ManyToManyDefinition stores information about a ManyToMany
// relationship between two domain objects.
type ManyToManyDefinition struct {
	dslengine.Definition
	DefinitionDSL    func()
	Left             *RelationalModelDefinition
	Right            *RelationalModelDefinition
	RelationshipName string // ??
	DatabaseField    string
}

// StoreIterator is a function that iterates over Relational Stores in a
// StorageGroup.
type StoreIterator func(m *RelationalStoreDefinition) error

// ModelIterator is a function that iterates over Models in a
// RelationalStore.
type ModelIterator func(m *RelationalModelDefinition) error

// FieldIterator is a function that iterates over Fields
// in a RelationalModel.
type FieldIterator func(m *RelationalFieldDefinition) error

// BuildSourceIterator is a function that iterates over Fields
// in a RelationalModel.
type BuildSourceIterator func(m *BuildSource) error
