package main

import (
	"fmt"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/gorma/example/app"
	"github.com/goadesign/gorma/example/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
)

var db *gorm.DB
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

	// Create service
	service := goa.New("Congo")

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(false))
	service.Use(middleware.Recover())

	// Mount "auth" controller
	c := NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "ui" controller
	c2 := NewUIController(service)
	app.MountUIController(service, c2)
	// Mount "user" controller
	c3 := NewUserController(service)
	app.MountUserController(service, c3)

	// Mount "proposal" controller
	c4 := NewProposalController(service)
	app.MountProposalController(service, c4)
	// Mount "review" controller
	c5 := NewReviewController(service)
	app.MountReviewController(service, c5)

	// Start service, listen on port 8080
	service.ListenAndServe(":8080")
}
