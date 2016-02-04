//************************************************************************//
// API "cellar": Models
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/gorma/example
// --design=github.com/goadesign/gorma/example/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package genmodels

import (
	"github.com/goadesign/gorma/example/app"
	"github.com/jinzhu/gorm"
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

// This is the Container model
type Container struct {
	AccountID     int // Belongs To Account
	Color         string
	Country       *string
	CreatedAt     *time.Time
	Href          string
	ID            int
	Name          string
	Rating        *int
	Region        *string
	Review        *string
	Sweetness     *int
	UpdatedAt     *time.Time
	Varietal      string
	Vineyard      string
	Vintage       string
	VinyardCounty *string
	DeletedAt     *time.Time // nullable timestamp (soft delete)
	Account       Account
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Container) TableName() string {
	return "containers"

}

// ContainerDB is the implementation of the storage interface for
// Container.
type ContainerDB struct {
	Db gorm.DB
	log.Logger
}

// NewContainerDB creates a new storage type.
func NewContainerDB(db gorm.DB, logger log.Logger) *ContainerDB {
	glog := logger.New("db", "Container")
	return &ContainerDB{Db: db, Logger: glog}
}

// DB returns the underlying database.
func (m *ContainerDB) DB() interface{} {
	return &m.Db
}

// ContainerStorage represents the storage interface.
type ContainerStorage interface {
	DB() interface{}
	List(ctx goa.Context) []Container
	One(ctx goa.Context) (Container, error)
	Add(ctx goa.Context, container Container) (Container, error)
	Update(ctx goa.Context, container Container) error
	Delete(ctx goa.Context) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *ContainerDB) TableName() string {
	return "containers"

}

// Transformation

func BottleToBottleApp(source *Bottle) (target *app.Bottle) {
	target = new(app.Bottle)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	return
}

// CRUD Functions
// ListBottle returns an array of view: default
func (m *ContainerDB) ListBottle(ctx goa.Context) []app.Bottle {
	now := time.Now()
	defer ctx.Info("ListBottle", "duration", time.Since(now))
	var objs []app.Bottle
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.Country = t.Country
	container.Href = t.Href
	container.Name = t.Name
	container.Vineyard = t.Vineyard
	container.Color = t.Color
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.CreatedAt = t.CreatedAt
	container.Review = t.Review
	container.VinyardCounty = t.VinyardCounty
	container.UpdatedAt = t.UpdatedAt
	container.Rating = t.Rating
	container.Region = t.Region
	container.Varietal = t.Varietal
*/
// return container

// OneBottle returns an array of view: default
func (m *ContainerDB) OneBottle(ctx goa.Context, id int) app.Bottle {
	now := time.Now()
	defer ctx.Info("OneBottle", "duration", time.Since(now))
	var view app.Bottle
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// Transformation

func BottleToBottleFullApp(source *Bottle) (target *app.BottleFull) {
	target = new(app.BottleFull)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.color = source.color
	target.country = source.country
	target.created_at = source.created_at
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.region = source.region
	target.review = source.review
	target.sweetness = source.sweetness
	target.updated_at = source.updated_at
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	target.vinyard_county = source.vinyard_county
	return
}

// CRUD Functions
// ListBottleFull returns an array of view: full
func (m *ContainerDB) ListBottleFull(ctx goa.Context) []app.BottleFull {
	now := time.Now()
	defer ctx.Info("ListBottleFull", "duration", time.Since(now))
	var objs []app.BottleFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.Color = t.Color
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.CreatedAt = t.CreatedAt
	container.Review = t.Review
	container.VinyardCounty = t.VinyardCounty
	container.UpdatedAt = t.UpdatedAt
	container.Rating = t.Rating
	container.Region = t.Region
	container.Varietal = t.Varietal
	container.Country = t.Country
	container.Href = t.Href
	container.Name = t.Name
	container.Vineyard = t.Vineyard
*/
// return container

// OneBottleFull returns an array of view: full
func (m *ContainerDB) OneBottleFull(ctx goa.Context, id int) app.BottleFull {
	now := time.Now()
	defer ctx.Info("OneBottleFull", "duration", time.Since(now))
	var view app.BottleFull
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// Transformation

func BottleToBottleTinyApp(source *Bottle) (target *app.BottleTiny) {
	target = new(app.BottleTiny)
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	return
}

// CRUD Functions
// ListBottleTiny returns an array of view: tiny
func (m *ContainerDB) ListBottleTiny(ctx goa.Context) []app.BottleTiny {
	now := time.Now()
	defer ctx.Info("ListBottleTiny", "duration", time.Since(now))
	var objs []app.BottleTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.Region = t.Region
	container.Varietal = t.Varietal
	container.UpdatedAt = t.UpdatedAt
	container.Rating = t.Rating
	container.Name = t.Name
	container.Vineyard = t.Vineyard
	container.Country = t.Country
	container.Href = t.Href
	container.Color = t.Color
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.VinyardCounty = t.VinyardCounty
	container.CreatedAt = t.CreatedAt
	container.Review = t.Review
*/
// return container

// OneBottleTiny returns an array of view: tiny
func (m *ContainerDB) OneBottleTiny(ctx goa.Context, id int) app.BottleTiny {
	now := time.Now()
	defer ctx.Info("OneBottleTiny", "duration", time.Since(now))
	var view app.BottleTiny
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// Transformation

func BoxToBoxApp(source *Box) (target *app.Box) {
	target = new(app.Box)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	return
}

// CRUD Functions
// ListBox returns an array of view: default
func (m *ContainerDB) ListBox(ctx goa.Context) []app.Box {
	now := time.Now()
	defer ctx.Info("ListBox", "duration", time.Since(now))
	var objs []app.Box
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.Rating = t.Rating
	container.Region = t.Region
	container.Varietal = t.Varietal
	container.UpdatedAt = t.UpdatedAt
	container.Href = t.Href
	container.Name = t.Name
	container.Vineyard = t.Vineyard
	container.Country = t.Country
	container.Color = t.Color
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.Review = t.Review
	container.VinyardCounty = t.VinyardCounty
	container.CreatedAt = t.CreatedAt
*/
// return container

// OneBox returns an array of view: default
func (m *ContainerDB) OneBox(ctx goa.Context, id int) app.Box {
	now := time.Now()
	defer ctx.Info("OneBox", "duration", time.Since(now))
	var view app.Box
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// Transformation

func BoxToBoxFullApp(source *Box) (target *app.BoxFull) {
	target = new(app.BoxFull)
	target.account = new(app.Account)
	target.account.created_at = source.account.created_at
	target.account.created_by = source.account.created_by
	target.account.href = source.account.href
	target.account.id = source.account.id
	target.account.name = source.account.name
	target.color = source.color
	target.country = source.country
	target.created_at = source.created_at
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	target.region = source.region
	target.review = source.review
	target.sweetness = source.sweetness
	target.updated_at = source.updated_at
	target.varietal = source.varietal
	target.vineyard = source.vineyard
	target.vintage = source.vintage
	target.vinyard_county = source.vinyard_county
	return
}

// CRUD Functions
// ListBoxFull returns an array of view: full
func (m *ContainerDB) ListBoxFull(ctx goa.Context) []app.BoxFull {
	now := time.Now()
	defer ctx.Info("ListBoxFull", "duration", time.Since(now))
	var objs []app.BoxFull
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.UpdatedAt = t.UpdatedAt
	container.Rating = t.Rating
	container.Region = t.Region
	container.Varietal = t.Varietal
	container.Country = t.Country
	container.Href = t.Href
	container.Name = t.Name
	container.Vineyard = t.Vineyard
	container.Color = t.Color
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.CreatedAt = t.CreatedAt
	container.Review = t.Review
	container.VinyardCounty = t.VinyardCounty
*/
// return container

// OneBoxFull returns an array of view: full
func (m *ContainerDB) OneBoxFull(ctx goa.Context, id int) app.BoxFull {
	now := time.Now()
	defer ctx.Info("OneBoxFull", "duration", time.Since(now))
	var view app.BoxFull
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// Transformation

func BoxToBoxTinyApp(source *Box) (target *app.BoxTiny) {
	target = new(app.BoxTiny)
	target.href = source.href
	target.id = source.id
	target.name = source.name
	target.rating = source.rating
	return
}

// CRUD Functions
// ListBoxTiny returns an array of view: tiny
func (m *ContainerDB) ListBoxTiny(ctx goa.Context) []app.BoxTiny {
	now := time.Now()
	defer ctx.Info("ListBoxTiny", "duration", time.Since(now))
	var objs []app.BoxTiny
	err := m.Db.Table(m.TableName()).Preload("Account").Find(&objs).Error
	if err != nil {
		ctx.Error("error listing Container", "error", err.Error())
		return objs
	}

	return objs
}

//container := Container{}
/*
 		container.CreatedAt = t.CreatedAt
	container.Review = t.Review
	container.VinyardCounty = t.VinyardCounty
	container.Varietal = t.Varietal
	container.UpdatedAt = t.UpdatedAt
	container.Rating = t.Rating
	container.Region = t.Region
	container.Vineyard = t.Vineyard
	container.Country = t.Country
	container.Href = t.Href
	container.Name = t.Name
	container.Sweetness = t.Sweetness
	container.Vintage = t.Vintage
	container.Color = t.Color
*/
// return container

// OneBoxTiny returns an array of view: tiny
func (m *ContainerDB) OneBoxTiny(ctx goa.Context, id int) app.BoxTiny {
	now := time.Now()
	defer ctx.Info("OneBoxTiny", "duration", time.Since(now))
	var view app.BoxTiny
	var native Container

	m.Db.Table(m.TableName()).Preload("Account").Where("id = ?", id).Find(&native)
	fmt.Println(native)
	return view

}

// GetContainer returns a single Container as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *ContainerDB) GetContainer(ctx goa.Context, id int) Container {
	now := time.Now()
	defer ctx.Info("GetContainer", "duration", time.Since(now))
	var native Container
	m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native)
	return native
}

// Add creates a new record.
func (m *ContainerDB) Add(ctx goa.Context, model Container) (Container, error) {
	now := time.Now()
	defer ctx.Info("AddContainer", "duration", time.Since(now))
	err := m.Db.Create(&model).Error
	if err != nil {
		ctx.Error("error updating Container", "error", err.Error())
		return model, err
	}

	return model, err
}

// Update modifies a single record.
func (m *ContainerDB) Update(ctx goa.Context, model Container) error {
	now := time.Now()
	defer ctx.Info("UpdateContainer", "duration", time.Since(now))
	obj := m.GetContainer(ctx)
	err := m.Db.Model(&obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *ContainerDB) Delete(ctx goa.Context) error {
	now := time.Now()
	defer ctx.Info("DeleteContainer", "duration", time.Since(now))
	var obj Container
	err := m.Db.Delete(&obj).Where("").Error

	if err != nil {
		ctx.Error("error retrieving Container", "error", err.Error())
		return err
	}

	return nil
}
