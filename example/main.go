package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/goadesign/gorma/example/app"
	"github.com/goadesign/gorma/example/app/v1"
	"github.com/goadesign/gorma/example/models"
	"github.com/goadesign/gorma/example/swagger"
	"github.com/goadesign/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
)

var db gorm.DB
var logger log15.Logger
var udb *models.UserDB
var rdb *models.ReviewDB
var pdb *models.ProposalDB

func main() {
	// Create service
	var err error
	//port, err := strconv.Atoi(strings.Split(c.Port(5432), ":")[1])
	//host := strings.Split(c.Port(5432), ":")[0]
	url := fmt.Sprintf("dbname=xorapidb user=docker password=docker sslmode=disable port=%d host=%s", 5432, "local.docker")
	fmt.Println(url)
	logger = log15.New("gorma", "example")
	time.Sleep(10)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DropTable(&models.Proposal{}, &models.Review{}, &models.User{})
	db.AutoMigrate(&models.Proposal{}, &models.Review{}, &models.User{})

	setup()
	// Create service
	service := goa.New("Congo")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest())
	service.Use(middleware.Recover())

	// Mount "auth" controller
	c := NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "ui" controller
	c2 := NewUiController(service)
	app.MountUiController(service, c2)
	// Mount "user" controller
	c3 := NewUserController(service)
	app.MountUserController(service, c3)

	// Version v1
	// Mount "proposal" controller
	c4 := NewProposalV1Controller(service)
	v1.MountProposalController(service, c4)
	// Mount "review" controller
	c5 := NewReviewV1Controller(service)
	v1.MountReviewController(service, c5)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}

func setup() error {
	gctx := context.Background()
	ctx := goa.NewContext(gctx, goa.New("setup"), nil, nil, nil)
	ctx.Logger = logger
	udb = models.NewUserDB(db, logger)

	bio := "A prolific debugger"
	city := "Tampa"
	country := "USA"
	email := "dude@congo.com"
	firstname := "Joe"
	lastname := "Bloggs"
	state := "Florida"
	act, err := udb.Add(ctx, models.User{
		Bio:       &bio,
		City:      &city,
		Country:   &country,
		Email:     &email,
		Firstname: &firstname,
		Lastname:  &lastname,
		State:     &state,
	})
	if err != nil {
		panic(err)
	}
	ctx.Info("created first acct", "account", act)
	abstract := "This is the abstract"
	detail := "This is the detail"
	title := "The TITLE"
	pdb = models.NewProposalDB(db, logger)

	prop, err := pdb.Add(ctx, models.Proposal{
		Abstract: &abstract,
		Detail:   &detail,
		Title:    &title,
		User:     act,
	})
	ctx.Info("created first proposal", "proposal", prop)
	comment := "Great Proposal!"
	rating := 5

	rdb := models.NewReviewDB(db, logger)

	rvw, err := rdb.Add(ctx, models.Review{
		Comment:  &comment,
		Rating:   &rating,
		User:     act,
		Proposal: prop,
	})
	ctx.Info("created first review", "review", rvw)
	return err
}
