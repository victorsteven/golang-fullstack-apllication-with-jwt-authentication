package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
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
