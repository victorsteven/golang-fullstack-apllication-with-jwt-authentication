package auto

import "golang_api_fullstack/api/models"

var users = []models.User{
	models.User{Nickname: "steven victor", Email: "steven@gmail.com", Password: "password"},
	// models.User{Nickname:"steven victor", Email:"steven@gmail.com", password:"23423434"}
	// models.User{Nickname:"steven victor", Email:"steven@gmail.com", password:"23423434"}
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Hello world",
	},
}
