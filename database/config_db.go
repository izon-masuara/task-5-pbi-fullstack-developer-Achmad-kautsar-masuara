package database

import (
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

	// if err := db.AutoMigrate(); err != nil {
	// 	panic(err.Error())
	// }

	Db = db
}
