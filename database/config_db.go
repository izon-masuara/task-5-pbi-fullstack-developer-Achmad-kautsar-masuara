package database

import (
	"fmt"
	"os"

	"github.com/izon-masuara/app"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	type Env struct {
		USER string
		PASS string
		HOST string
		PORT string
		NAME string
	}
	var appConf = Env{
		USER: os.Getenv("DB_USER"),
		PASS: os.Getenv("DB_PASSWORD"),
		HOST: os.Getenv("DB_HOST"),
		PORT: os.Getenv("DB_PORT"),
		NAME: os.Getenv("DB_NAME"),
	}
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", appConf.USER, appConf.PASS, appConf.HOST, appConf.PORT, appConf.NAME)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(app.User{}, app.Photo{})

	if err != nil {
		panic(err.Error())
	}

	Db = db
}
