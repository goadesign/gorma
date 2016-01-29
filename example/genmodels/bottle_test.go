package genmodels

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/ory-am/dockertest"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	if c, err := dockertest.ConnectToPostgreSQL(15, time.Second, func(url string) bool {
		var err error
		db, err = gorm.Open("postgres", url)
		if err != nil {
			return false
		}
		return db.Ping() == nil
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer c.KillRemove()
	os.Exit(m.Run())
}

func TestFunction(t *testing.T) {
	db.Ping()
}
