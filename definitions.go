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
	*design.AttributeDefinition
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
	SourceMaps       map[string]*SourceMapping
	TargetMaps       map[string]*TargetMapping
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

type SourceMapping struct {
	dslengine.Definition
	*design.AttributeDefinition
	DefinitionDSL func()
	MappingName   string
	Description   string
	Remote        *design.UserTypeDefinition
	Parent        *RelationalModelDefinition
	Mappings      map[string]*MapDefinition // keyed by gorma field
}

type TargetMapping struct {
	dslengine.Definition
	*design.AttributeDefinition
	DefinitionDSL func()
	MappingName   string
	Description   string
	Remote        *design.UserTypeDefinition
	Parent        *RelationalModelDefinition
	Mappings      map[string]*MapDefinition // keyed by gorma field
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
type SourceMapIterator func(m *SourceMapping) error

// TargetMapIterator
type TargetMapIterator func(m *TargetMapping) error

// FieldIterator is a function that iterates over Fields
// in a RelationalModel
type FieldIterator func(m *RelationalFieldDefinition) error

// Kind implements DataKind.
func (u *RelationalModelDefinition) Kind() design.Kind { return design.UserTypeKind }

// Name returns the JSON type name.
func (u *RelationalModelDefinition) Name() string { return u.Type.Name() }

// IsPrimitive calls IsPrimitive on the user type underlying data type.
func (u *RelationalModelDefinition) IsPrimitive() bool { return u.Type.IsPrimitive() }

// IsObject calls IsObject on the user type underlying data type.
func (u *RelationalModelDefinition) IsObject() bool { return u.Type.IsObject() }

// IsArray calls IsArray on the user type underlying data type.
func (u *RelationalModelDefinition) IsArray() bool { return u.Type.IsArray() }

// IsHash calls IsHash on the user type underlying data type.
func (u *RelationalModelDefinition) IsHash() bool { return u.Type.IsHash() }

// ToObject calls ToObject on the user type underlying data type.
func (u *RelationalModelDefinition) ToObject() design.Object { return u.Type.ToObject() }

// ToArray calls ToArray on the user type underlying data type.
func (u *RelationalModelDefinition) ToArray() *design.Array { return u.Type.ToArray() }

// ToHash calls ToHash on the user type underlying data type.
func (u *RelationalModelDefinition) ToHash() *design.Hash { return u.Type.ToHash() }

// IsCompatible returns true if val is compatible with p.
func (u *RelationalModelDefinition) IsCompatible(val interface{}) bool {
	return u.Type.IsCompatible(val)
}

// Kind implements DataKind.
func (u *SourceMapping) Kind() design.Kind { return design.UserTypeKind }

// Name returns the JSON type name.
func (u *SourceMapping) Name() string { return u.Type.Name() }

// IsPrimitive calls IsPrimitive on the user type underlying data type.
func (u *SourceMapping) IsPrimitive() bool { return u.Type.IsPrimitive() }

// IsObject calls IsObject on the user type underlying data type.
func (u *SourceMapping) IsObject() bool { return u.Type.IsObject() }

// IsArray calls IsArray on the user type underlying data type.
func (u *SourceMapping) IsArray() bool { return u.Type.IsArray() }

// IsHash calls IsHash on the user type underlying data type.
func (u *SourceMapping) IsHash() bool { return u.Type.IsHash() }

// ToObject calls ToObject on the user type underlying data type.
func (u *SourceMapping) ToObject() design.Object { return u.Type.ToObject() }

// ToArray calls ToArray on the user type underlying data type.
func (u *SourceMapping) ToArray() *design.Array { return u.Type.ToArray() }

// ToHash calls ToHash on the user type underlying data type.
func (u *SourceMapping) ToHash() *design.Hash { return u.Type.ToHash() }

// IsCompatible returns true if val is compatible with p.
func (u *SourceMapping) IsCompatible(val interface{}) bool {
	return u.Type.IsCompatible(val)
}
