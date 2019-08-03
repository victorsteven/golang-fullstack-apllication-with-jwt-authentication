package auto

import (
	"golang_api_fullstack/api/database"
	"golang_api_fullstack/api/models"
	"golang_api_fullstack/api/utils/console"
	"log"
)

func Load() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalf("cannot connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed table: %v", err)
		}
		console.Pretty(user)
	}
}
