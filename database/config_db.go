package database

import (
	"github.com/izon-masuara/app"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	psqlconn := "postgres://postgres:wppq@localhost:5432/User_BPTNS"
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
