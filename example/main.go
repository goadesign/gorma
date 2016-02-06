package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/goadesign/gorma/example/genmodels"
	"github.com/goadesign/gorma/example/swagger"
	"github.com/goadesign/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
)

var db gorm.DB
var logger log15.Logger
var adb *genmodels.AccountDB
var bdb *genmodels.BottleDB

func main() {
	// Create service
	service := goa.New("API")
	var err error
	//port, err := strconv.Atoi(strings.Split(c.Port(5432), ":")[1])
	//host := strings.Split(c.Port(5432), ":")[0]
	url := fmt.Sprintf("dbname=xorapidb user=docker password=docker sslmode=disable port=%d host=%s", 5432, "local.docker")
	fmt.Println(url)
	logger = log15.New("something", "example")
	time.Sleep(10)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.DropTable(&genmodels.Bottle{}, &genmodels.Account{})
	db.AutoMigrate(&genmodels.Bottle{}, &genmodels.Account{})

	setup()
	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)
	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}

func setup() error {
	gctx := context.Background()
	ctx := goa.NewContext(gctx, goa.New("test"), nil, nil, nil)
	ctx.Logger = logger
	adb = genmodels.NewAccountDB(db, logger)
	cb := "Brian"
	act, err := adb.Add(ctx, genmodels.Account{
		CreatedBy: cb,
		Href:      "href",
		Name:      "Account1",
	})
	if err != nil {
		panic(err)
	}

	bdb = genmodels.NewBottleDB(db, logger)

	var Color string
	var Country string
	var Name string
	var Region string
	var Review string
	var Sweetness int
	var Varietal string
	var Vineyard string
	var Vintage string
	var Rating int
	var VinyardCounty string
	Color = "Blue"
	Country = "Australia"
	Name = "Red Horse"
	Region = "South"
	Review = "crappy"
	Sweetness = 4
	Rating = 99
	Varietal = "Merlot"
	Vineyard = "Robert Mondavi"
	Vintage = "1999"
	VinyardCounty = "Cork"
	btl, err := bdb.Add(ctx, genmodels.Bottle{
		AccountID:     act.ID,
		Color:         &Color,
		Country:       &Country,
		Name:          &Name,
		Rating:        &Rating,
		Region:        &Region,
		Review:        &Review,
		Sweetness:     &Sweetness,
		Varietal:      &Varietal,
		Vineyard:      &Vineyard,
		Vintage:       &Vintage,
		VinyardCounty: &VinyardCounty,
	})
	fmt.Println(btl.ID, btl.AccountID)
	return err
}
