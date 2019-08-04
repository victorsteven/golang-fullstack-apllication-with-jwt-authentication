package auto

import (
	"golang_api_fullstack/api/database"
	"golang_api_fullstack/api/models"
	"log"
)

func Load() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalf("cannot connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		// err = db.Debug().Model(&models.Post{}).Where("author_id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
		// console.Pretty(posts[i])
	}
}
