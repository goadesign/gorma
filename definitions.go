package gorma

import "github.com/raphael/goa/design"

// StorageGroup is the parent configuration structure for Gorma definitions
type StorageGroup struct {
	api             *design.APIDefinition
	RelationalStore *RelationalStore
	KVStore         *KVStore
}

// Store represents a single object storage mechanism
// i.e. a relational database or KV store
type Store interface {
	Models() []Model
}

// Model is an interface for the representation of a single
// domain object
type Model interface {
	Name() string
	Fields() []Field
	StorageAlias() string
	PrimaryKey() Field
	BelongsTo() []Model
	HasMany() []Model
	HasOne() []Model
	ManyToMany() []ManyToMany
	Cached() bool
}

// ManyToMany stores information about a ManyToMany
// relationship between two domain objects
type ManyToMany struct {
	Left             *RelationalModel
	Right            *RelationalModel
	LeftNamePlural   string
	RightNamePlural  string
	LeftName         string
	RightName        string
	RelationshipName string
	DatabaseField    string
}

// Field is an abstraction of a field in a data store
type Field interface {
	Name() string
	Datatype() string
	Nullable() bool // use pointer
}

// RelationalStore is the parent configuration structure for Gorm relational model definitions
type RelationalStore struct {
	Models map[string]*RelationalModel
}

// KVStore is the parent configuration structure for Gorm KV model definitions
type KVStore struct {
	Models map[string]*KVModel
}

// RelationalField implements the Field interface and represents
// a field in a relational database
type RelationalField struct {
	a                 *design.AttributeDefinition
	Name              string
	Datatype          string
	SQLTag            string
	DatabaseFieldName string
	Description       string
	Nullable          bool
	PrimaryKey        bool
	Timestamp         bool
	Aliased           bool
	BelongsTo         string
	HasOne            string
	HasMany           string
	Many2Many         string
}

// KVField implements the Field interface and represents
// a field in a KV store
type KVField struct {
	Name     string
	Datatype string
}

// RelationalModel implements the Model interface and
// implements the storage of a domain model into a
// table in a relational database
type RelationalModel struct {
	utd              *design.UserTypeDefinition
	BelongsTo        map[string]*RelationalModel
	HasMany          map[string]*RelationalModel
	HasOne           map[string]*RelationalModel
	ManyToMany       map[string]*ManyToMany
	Fields           map[string]*RelationalField
	TableName        string
	Name             string
	Alias            string
	Description      string
	Cached           bool
	CacheDuration    int
	NoMedia          bool
	Roler            bool
	DynamicTableName bool
	SQLTag           string
	PrimaryKeys      []*RelationalField
	belongsto        []string
	hasmany          []string
	hasone           []string
	many2many        []string
}

// KVModel implements the Model interface and represents
// the storage of a domain model in a KV store
type KVModel struct {
	BelongsTo  map[string]*KVModel
	HasMany    map[string]*KVModel
	HasOne     map[string]*KVModel
	ManyToMany map[string]*ManyToMany
}

type ModelIterator func(m *RelationalModel) error
type FieldIterator func(m *RelationalField) error
