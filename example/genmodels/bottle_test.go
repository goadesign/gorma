package genmodels

import (
	"fmt"
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"
	_ "github.com/lib/pq"
)

var db gorm.DB

func TestMain(m *testing.M) {
	//c := dockertest.Launch("registry.xordataexchange.com/xor/postgres-dev")
	//fmt.Println(c)
	var err error
	//port, err := strconv.Atoi(strings.Split(c.Port(5432), ":")[1])
	//host := strings.Split(c.Port(5432), ":")[0]
	url := fmt.Sprintf("dbname=xorapidb user=docker password=docker sslmode=disable port=%d host=%s", 32778, "192.168.100.5")
	fmt.Println(url)
	time.Sleep(10)
	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.DropTable(&Bottle{}, &Account{})
	db.AutoMigrate(&Bottle{}, &Account{})
	setup()
	//defer c.Close()
	os.Exit(m.Run())
}

func TestOneBottle(t *testing.T) {
	db.LogMode(true)
	bdb := NewBottleDB(db)
	btl := bdb.OneBottle(context.Background(), 1)
	pretty.Println(btl)
}

func setup() error {
	adb := NewAccountDB(db)
	act, err := adb.Add(context.Background(), Account{
		CreatedBy: "Brian",
		Href:      "href",
		Name:      "Account1",
	})
	if err != nil {
		panic(err)
	}
	pretty.Println(act)

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
	var VinyardCounty string
	Color = "Blue"
	Country = "Australia"
	Myvintage = "MyVintage"
	Name = "Red Horse"
	Region = "South"
	Review = "crappy"
	Sweetness = 4
	Varietal = "Merlot"
	Vineyard = "Robert Mondavi"
	Vintage = "1999"
	VinyardCounty = "Cork"
	btl, err := bdb.Add(context.Background(), Bottle{
		AccountID:     act.ID,
		Color:         &Color,
		Country:       &Country,
		Myvintage:     &Myvintage,
		Name:          &Name,
		Region:        &Region,
		Review:        &Review,
		Sweetness:     &Sweetness,
		Varietal:      &Varietal,
		Vineyard:      &Vineyard,
		Vintage:       &Vintage,
		VinyardCounty: &VinyardCounty,
	})
	pretty.Println(btl)
	return err
}
