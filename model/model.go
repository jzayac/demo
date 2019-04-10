package model

import (
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"

	"demo/util/database"
)

var logger *log.Logger

type Iata struct {
	gorm.Model
	Index string `gorm:"not null;unique"`
	Value string
}

func dbConnect() (*gorm.DB, error) {
	db, err := database.PgNewClient()
	return db, err
}

func InitDatabaseModel() error {
	db, err := dbConnect()

	if err != nil {
		return database.ErrSqlInitDbDriver
	}
	defer db.Close()

	// Automatically migrate your schema, to keep your schema update to date.
	db.AutoMigrate(Iata{})

	// db.Create(&Iata{
	// 	Index: "1234",
	// 	Value: "test airlines",
	// })

	// db.Create(&Iata{
	// 	Index: "1235",
	// 	Value: "slovak airlines",
	// })

	return nil
}
