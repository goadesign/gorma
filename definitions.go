package gorma

// StorageGroup is the parent configuration structure for Gorma definitions
type StorageGroup struct {
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
	Left             Model
	Right            Model
	LeftKeyField     Field
	RightKeyField    Field
	RelationshipName string
}

// Field is an abstraction of a field in a data store
type Field interface {
	Name() string
	Datatype() string
}

// RelationalStorage is the parent configuration structure for Gorm relational model definitions
type RelationalStore struct {
	Models map[string]*RelationalModel
}

// KVStorage is the parent configuration structure for Gorm KV model definitions
type KVStore struct {
	Models map[string]*KVModel
}

// RelationalField implements the Field interface and represents
// a field in a relational database
type RelationalField struct {
	Name     string
	Datatype string
	SQLTag   string
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
	BelongsTo  map[string]*RelationalModel
	HasMany    map[string]*RelationalModel
	HasOne     map[string]*RelationalModel
	HanyToMany map[string]*ManyToMany
	Fields     map[string]*RelationalField
	TableName  string
	Alias      string
	Cached     bool
	Media      bool
}

// KVModel implements the Model interface and represents
// the storage of a domain model in a KV store
type KVModel struct {
	BelongsTo  map[string]*KVModel
	HasMany    map[string]*KVModel
	HasOne     map[string]*KVModel
	ManyToMany map[string]*ManyToMany
}
