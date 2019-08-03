package controllers

import (
	"encoding/json"
	"fmt"
	"golang_api_fullstack/api/database"
	"golang_api_fullstack/api/models"
	"golang_api_fullstack/api/repository"
	"golang_api_fullstack/api/repository/crud"
	"golang_api_fullstack/api/responses"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Create user"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err := usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
