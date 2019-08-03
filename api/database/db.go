package database

import (
	"golang_api_fullstack/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		// log.Printf("error connecting to the database: %v", err)
		return nil, err
	}
	return db, nil
}
