package genmodels

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gopkg.in/inconshreveable/log15.v2"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB
var logger log15.Logger
var ctx *goa.Context

func TestMain(m *testing.M) {
	//c := dockertest.Launch("registry.xordataexchange.com/xor/postgres-dev")
	//fmt.Println(c)
	var err error
	//port, err := strconv.Atoi(strings.Split(c.Port(5432), ":")[1])
	//host := strings.Split(c.Port(5432), ":")[0]
	url := fmt.Sprintf("dbname=xorapidb user=docker password=docker sslmode=disable port=%d host=%s", 32779, "192.168.100.5")
	fmt.Println(url)
	time.Sleep(10)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.DropTable(&Bottle{}, &Account{})
	db.AutoMigrate(&Bottle{}, &Account{})
	logger = log15.New("tests", "bottle")
	gctx := context.Background()
	ctx = goa.NewContext(gctx, goa.New("test"), nil, nil, nil)
	ctx.Logger = logger
	setup()
	//defer c.Close()
	os.Exit(m.Run())
}

func TestOneBottle(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db)
	btl := bdb.OneBottle(*ctx, 1)
	fmt.Println(btl.ID)
}
func TestGetBottle(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db)
	btl := bdb.GetBottle(*ctx, 1)
	if btl.ID != 1 {
		t.Error("Expected Bottle")
	}
}
func TestOneAccount(t *testing.T) {
	db.LogMode(true)
	adb := NewAccountDB(db)
	act := adb.OneAccount(*ctx, 1)
	fmt.Println(act.ID)
}
func TestGetAccount(t *testing.T) {
	db.LogMode(true)
	adb := NewAccountDB(db)
	act := adb.GetAccount(*ctx, 1)
	if act.ID != 1 {
		t.Error("Expected account")
	}
}
func setup() error {
	adb := NewAccountDB(db)
	act, err := adb.Add(*ctx, Account{
		CreatedBy: "Brian",
		Href:      "href",
		Name:      "Account1",
	})
	if err != nil {
		panic(err)
	}

	bdb := NewBottleDB(db)

	var Color string
	var Country string
	var Myvintage string
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
	Myvintage = "MyVintage"
	Name = "Red Horse"
	Region = "South"
	Review = "crappy"
	Sweetness = 4
	Rating = 99
	Varietal = "Merlot"
	Vineyard = "Robert Mondavi"
	Vintage = "1999"
	VinyardCounty = "Cork"
	btl, err := bdb.Add(*ctx, Bottle{
		AccountID:     act.ID,
		Color:         &Color,
		Country:       &Country,
		Myvintage:     &Myvintage,
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
