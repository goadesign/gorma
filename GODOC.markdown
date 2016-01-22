# gorma
--
    import "github.com/bketelsen/gorma"


## Usage

```go
const (
	// Gorma is the constant string used as the index in the
	// goa DesignConstructs map
	Gorma = "gorma"
	// StorageGroup is the constant string used as the index in the
	// GormaConstructs map
	StorageGroup                            = "storagegroup"
	MySQL             RelationalStorageType = "mysql"
	Postgres          RelationalStorageType = "postgres"
	Boolean           FieldType             = "bool"
	Integer           FieldType             = "integer"
	BigInteger        FieldType             = "biginteger"
	AutoInteger       FieldType             = "auto_integer"
	AutoBigInteger    FieldType             = "auto_biginteger"
	Decimal           FieldType             = "decimal"
	BigDecimal        FieldType             = "bigdecimal"
	String            FieldType             = "string"
	Text              FieldType             = "text"
	UUID              FieldType             = "uuid"
	PKInteger         FieldType             = "pkinteger"
	PKBigInteger      FieldType             = "pkbiginteger"
	PKUUID            FieldType             = "pkuuid"
	Timestamp         FieldType             = "timestamp"
	NullableTimestamp FieldType             = "nulltimestamp"
	NotFound          FieldType             = "notfound"
)
```

```go
var (
	// TargetPackage is the name of the generated Go package.
	TargetPackage string
	AppPackage    string
)
```

#### func  AppOutputDir

```go
func AppOutputDir() string
```
AppOutputDir returns the directory containing the generated files.

#### func  AppPackagePath

```go
func AppPackagePath() (string, error)
```
AppPackagePath returns the Go package path to the generated package.

#### func  Generate

```go
func Generate(api *design.APIDefinition) ([]string, error)
```
Generate is the generator entry point called by the meta generator.

#### func  Init

```go
func Init()
```
Init creates the necessary data structures for parsing a DSL

#### func  ModelOutputDir

```go
func ModelOutputDir() string
```
ModelOutputDir returns the directory containing the generated files.

#### func  ModelPackagePath

```go
func ModelPackagePath() (string, error)
```
ModelPackagePath returns the Go package path to the generated package.

#### type Command

```go
type Command struct {
	*codegen.BaseCommand
}
```

Command is the goa application code generator command line data structure.

#### func  NewCommand

```go
func NewCommand() *Command
```
NewCommand instantiates a new command.

#### func (*Command) RegisterFlags

```go
func (c *Command) RegisterFlags(r codegen.FlagRegistry)
```
RegisterFlags registers the command line flags with the given registry.

#### func (*Command) Run

```go
func (c *Command) Run() ([]string, error)
```
Run simply calls the meta generator.

#### type FieldIterator

```go
type FieldIterator func(m *RelationalFieldDefinition) error
```

FieldIterator is a function that iterates over Fields in a RelationalModel

#### type FieldType

```go
type FieldType string
```

FieldType is the storage type for a database field

#### type Generator

```go
type Generator struct {
}
```

Generator is the application code generator.

#### func  NewGenerator

```go
func NewGenerator() (*Generator, error)
```
NewGenerator returns the application code generator.

#### func (*Generator) Cleanup

```go
func (g *Generator) Cleanup()
```
Cleanup removes the entire "app" directory if it was created by this generator.

#### func (*Generator) Generate

```go
func (g *Generator) Generate(api *design.APIDefinition) (_ []string, err error)
```
Generate the application code, implement codegen.Generator.

#### type ManyToManyDefinition

```go
type ManyToManyDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Left             *RelationalModelDefinition
	Right            *RelationalModelDefinition
	RelationshipName string // ??
	DatabaseField    string
}
```

ManyToManyDefinition stores information about a ManyToMany relationship between
two domain objects

#### func (*ManyToManyDefinition) LeftName

```go
func (m *ManyToManyDefinition) LeftName() string
```

#### func (*ManyToManyDefinition) LeftNamePlural

```go
func (m *ManyToManyDefinition) LeftNamePlural() string
```

#### func (*ManyToManyDefinition) LowerLeftName

```go
func (m *ManyToManyDefinition) LowerLeftName() string
```

#### func (*ManyToManyDefinition) LowerRightName

```go
func (m *ManyToManyDefinition) LowerRightName() string
```

#### func (*ManyToManyDefinition) RightName

```go
func (m *ManyToManyDefinition) RightName() string
```

#### func (*ManyToManyDefinition) RightNamePlural

```go
func (m *ManyToManyDefinition) RightNamePlural() string
```

#### type MediaTypeAdapterDefinition

```go
type MediaTypeAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.MediaTypeDefinition
	Right         *RelationalModelDefinition
}
```

MediaTypeAdapterDefinition represents the transformation of a Goa media type
into a Gorma Model

#### type ModelIterator

```go
type ModelIterator func(m *RelationalModelDefinition) error
```

ModelIterator is a function that iterates over Models in a RelationalStore

#### type PayloadAdapterDefinition

```go
type PayloadAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *design.UserTypeDefinition
	Right         *RelationalModelDefinition
}
```

PayloadAdapterDefinition represents the transformation of a Goa Payload (which
is really a UserTypeDefinition into a Gorma model

#### type RelationalFieldDefinition

```go
type RelationalFieldDefinition struct {
	design.Definition
	DefinitionDSL func()
	Parent        *RelationalModelDefinition

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
```

RelationalFieldDefinition implements the Field interface and represents a field
in a relational database

#### func (RelationalFieldDefinition) Children

```go
func (f RelationalFieldDefinition) Children() []design.Definition
```
Children returnsa slice of this objects children

#### func (*RelationalFieldDefinition) Context

```go
func (f *RelationalFieldDefinition) Context() string
```
Context returns the generic definition name used in error messages.

#### func (*RelationalFieldDefinition) DSL

```go
func (f *RelationalFieldDefinition) DSL() func()
```
DSL returns this object's DSL

#### func (*RelationalFieldDefinition) FieldDefinition

```go
func (f *RelationalFieldDefinition) FieldDefinition() string
```
Definition returns the field's struct definition

#### func (*RelationalFieldDefinition) LowerName

```go
func (f *RelationalFieldDefinition) LowerName() string
```

#### func (*RelationalFieldDefinition) Tags

```go
func (f *RelationalFieldDefinition) Tags() string
```
Tags returns teh sql and gorm struct tags for the Definition

#### func (*RelationalFieldDefinition) Validate

```go
func (field *RelationalFieldDefinition) Validate() *design.ValidationErrors
```
Validate tests whether the RelationalField definition is consistent

#### type RelationalModelDefinition

```go
type RelationalModelDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string //
	Parent           *RelationalStoreDefinition
	BuiltFrom        []*design.UserTypeDefinition
	RenderTo         []*design.MediaTypeDefinition
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
}
```

RelationalModelDefinition implements the storage of a domain model into a table
in a relational database

#### func (RelationalModelDefinition) Children

```go
func (sd RelationalModelDefinition) Children() []design.Definition
```
Children returnsa slice of this objects children

#### func (*RelationalModelDefinition) Context

```go
func (f *RelationalModelDefinition) Context() string
```
Context returns the generic definition name used in error messages.

#### func (*RelationalModelDefinition) DSL

```go
func (sd *RelationalModelDefinition) DSL() func()
```
DSL returns this object's DSL

#### func (*RelationalModelDefinition) IterateFields

```go
func (f *RelationalModelDefinition) IterateFields(it FieldIterator) error
```
IterateFields returns an iterator function useful for iterating through this
model's field list

#### func (*RelationalModelDefinition) LowerName

```go
func (f *RelationalModelDefinition) LowerName() string
```

#### func (*RelationalModelDefinition) PKAttributes

```go
func (f *RelationalModelDefinition) PKAttributes() string
```
PKAttributes constructs a pair of field + definition strings useful for method
parameters

#### func (*RelationalModelDefinition) PKUpdateFields

```go
func (f *RelationalModelDefinition) PKUpdateFields() string
```
PKUpdateFields returns something? This function doesn't look useful in current
form. Perhaps it isnt.

#### func (*RelationalModelDefinition) PKWhere

```go
func (f *RelationalModelDefinition) PKWhere() string
```
PKWhere returns an array of strings representing the where clause of a retrieval
by primary key(s) -- x = ? and y = ?

#### func (*RelationalModelDefinition) PKWhereFields

```go
func (f *RelationalModelDefinition) PKWhereFields() string
```
PKWhereFields returns the fields for a where clause for the primary keys of a
model

#### func (*RelationalModelDefinition) PopulateFromModeledType

```go
func (f *RelationalModelDefinition) PopulateFromModeledType()
```
PopulateFromModeledType creates fields for the model based on the goa
UserTypeDefinition it models This happens before fields are processed, so it's
ok to just assign without testing

#### func (*RelationalModelDefinition) StructDefinition

```go
func (f *RelationalModelDefinition) StructDefinition() string
```
Definition returns the struct definition for the model

#### func (*RelationalModelDefinition) Validate

```go
func (a *RelationalModelDefinition) Validate() *design.ValidationErrors
```
Validate tests whether the RelationalModel definition is consistent

#### type RelationalStorageType

```go
type RelationalStorageType string
```

RelationalStorageType is the type of database

#### type RelationalStoreDefinition

```go
type RelationalStoreDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	Parent           *StorageGroupDefinition
	Type             RelationalStorageType
	RelationalModels map[string]*RelationalModelDefinition
}
```

RelationalStoreDefinition is the parent configuration structure for Gorm
relational model definitions

#### func (RelationalStoreDefinition) Children

```go
func (sd RelationalStoreDefinition) Children() []design.Definition
```
Children returnsa slice of this objects children

#### func (*RelationalStoreDefinition) Context

```go
func (sd *RelationalStoreDefinition) Context() string
```
Context returns the generic definition name used in error messages.

#### func (*RelationalStoreDefinition) DSL

```go
func (sd *RelationalStoreDefinition) DSL() func()
```
DSL returns this object's DSL

#### func (*RelationalStoreDefinition) IterateModels

```go
func (sd *RelationalStoreDefinition) IterateModels(it ModelIterator) error
```
IterateModels runs an iterator function once per Model in the Store's model
list.

#### func (*RelationalStoreDefinition) Validate

```go
func (a *RelationalStoreDefinition) Validate() *design.ValidationErrors
```
Validate tests whether the RelationalStore definition is consistent

#### type StorageGroupDefinition

```go
type StorageGroupDefinition struct {
	design.Definition
	DefinitionDSL    func()
	Name             string
	Description      string
	RelationalStores map[string]*RelationalStoreDefinition
}
```

StorageGroupDefinition is the parent configuration structure for Gorma
definitions

```go
var GormaDesign *StorageGroupDefinition
```
GormaDesign is the root definition for Gorma

#### func (StorageGroupDefinition) Children

```go
func (sd StorageGroupDefinition) Children() []design.Definition
```
Children returnsa slice of this objects children

#### func (StorageGroupDefinition) Context

```go
func (sd StorageGroupDefinition) Context() string
```
Context returns the generic definition name used in error messages.

#### func (StorageGroupDefinition) DSL

```go
func (sd StorageGroupDefinition) DSL() func()
```
DSL returns this object's DSL

#### func (*StorageGroupDefinition) IterateSets

```go
func (a *StorageGroupDefinition) IterateSets(iterator design.SetIterator)
```
IterateSets goes over all the definition sets of the StorageGroup: The
StorageGroup definition itself, each store definition, models and fields.

#### func (*StorageGroupDefinition) IterateStores

```go
func (sd *StorageGroupDefinition) IterateStores(it StoreIterator) error
```
IterateStores runs an iterator function once per Relational Store in the
StorageGroup's Store list

#### func (*StorageGroupDefinition) Validate

```go
func (a *StorageGroupDefinition) Validate() *design.ValidationErrors
```
Validate tests whether the StorageGroup definition is consistent

#### type StoreIterator

```go
type StoreIterator func(m *RelationalStoreDefinition) error
```

StoreIterator is a function that iterates over Relational Stores in a
StorageGroup

#### type UserTypeAdapterDefinition

```go
type UserTypeAdapterDefinition struct {
	design.Definition
	DefinitionDSL func()
	Name          string
	Description   string
	Left          *RelationalModelDefinition
	Right         *RelationalModelDefinition
}
```

UserTypeAdapterDefinition represents the transformation of a Goa user type into
a Gorma Model

#### type UserTypeTemplateData

```go
type UserTypeTemplateData struct {
	APIDefinition *design.APIDefinition
	UserType      *RelationalModelDefinition
	DefaultPkg    string
	AppPkg        string
}
```

UserTypeTemplateData contains all the information used by the template to redner
the media types code.

#### type UserTypesWriter

```go
type UserTypesWriter struct {
	*codegen.SourceFile
	UserTypeTmpl *template.Template
}
```

UserTypesWriter generate code for a goa application user types. User types are
data structures defined in the DSL with "Type".

#### func  NewUserTypesWriter

```go
func NewUserTypesWriter(filename string) (*UserTypesWriter, error)
```
NewUserTypesWriter returns a contexts code writer. User types contain custom
data structured defined in the DSL with "Type".

#### func (*UserTypesWriter) Execute

```go
func (w *UserTypesWriter) Execute(data *UserTypeTemplateData) error
```
Execute writes the code for the context types to the writer.
