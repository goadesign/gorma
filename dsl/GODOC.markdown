# dsl
--
    import "github.com/bketelsen/gorma/dsl"


## Usage

#### func  Alias

```go
func Alias(d string)
```
Alias overrides the name of the sql store's table or field

#### func  BelongsTo

```go
func BelongsTo(parent string)
```
BelongsTo signifies a relationship between this model and a Parent. The Parent
has the child, and the Child belongs to the Parent

#### func  BuiltFrom

```go
func BuiltFrom(mts ...*design.UserTypeDefinition)
```

#### func  Cached

```go
func Cached(d string)
```
Cached caches the models for `duration` seconds

#### func  Description

```go
func Description(d string)
```
Description sets the definition description. Description can be called inside
StorageGroup, RelationalStore, RelationalModel, RelationalField

#### func  DynamicTableName

```go
func DynamicTableName()
```
DynamicTableName sets a boolean flag that cause the generator generate function
definitions in the database models that specify the name of the database table.
Useful when using multiple tables with different names but same schema e.g.
Users, AdminUsers

#### func  Field

```go
func Field(name string, args ...interface{})
```
RelationalField is a DSL definition for a field in a Relational Model Examples
and more docs here later Parameter Options: Field("Title") Field("Title",
gorma.String) Field("Title", func(){... other field level dsl ...})
Field("Title", gorma.String, func(){... other field level dsl ...})

#### func  HasMany

```go
func HasMany(name, child string)
```
HasMany signifies a relationship between this model and a set of Children. The
Parent has the children, and the Children belong to the Parent

#### func  HasOne

```go
func HasOne(child string)
```
HasOne signifies a relationship between this model and a Child. The Parent has
the child, and the Child belongs to the Parent

#### func  ManyToMany

```go
func ManyToMany(other, tablename string)
```
ManyToMany creates a join table to store the intersection relationship between
this model and another model. For example, in retail an Order can contain many
products, and a product can belong to many orders. To express this relationship
use the following syntax: Model("Order", func(){

    ManyToMany("Product", "order_lines")

}) This specifies that the Order and Product tables have a "junction" table
called `order_lines` that contains the order and product information The
generated model will have a field called `Products` that will be an array of
type `product.Product`

#### func  Model

```go
func Model(name string, dsl func())
```
RelationalModel is the DSL that represents a Relational Model Examples and more
docs here later

#### func  NoMedia

```go
func NoMedia()
```
NoMedia sets a boolean flag that prevents the generation of media helpers

#### func  RenderTo

```go
func RenderTo(mts ...*design.MediaTypeDefinition)
```

#### func  Roler

```go
func Roler()
```
Roler sets a boolean flag that cause the generation of a Role() function that
returns the model's Role value Requires a field in the model named Role, type
String

#### func  SQLTag

```go
func SQLTag(d string)
```
SQLTag sets the model's struct tag `sql` value for indexing and other purposes

#### func  StorageGroup

```go
func StorageGroup(name string, dsli func()) *gorma.StorageGroupDefinition
```
StorageGroup implements the top level Gorma DSL Examples and more docs here
later

#### func  Store

```go
func Store(name string, storeType gorma.RelationalStorageType, dsl func())
```
StorageGroup implements the top level Gorma DSL Examples and more docs here
later

#### func  TableName

```go
func TableName(d string)
```
TableName creates a TableName() function that returns the name of the table to
query. Useful for pre-existing schemas
