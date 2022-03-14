package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

// get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Users)
	if err != nil {
		fmt.Println(err)
	}

}

// create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book User
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100)) // Mock ID - not safe
	Users = append(Users, book)
	json.NewEncoder(w).Encode(book)
}

// get single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Users {
		if item.ID == params["id"] {
			Users = append(Users[:index], Users[index+1:]...)
			break
		}
	}
	w.Write([]byte("User deleted"))
}

// update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Users {
		if item.ID == params["id"] {
			Users = append(Users[:index], Users[index+1:]...)
			break
		}
	}
	var book User
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	Users = append(Users, book)
	json.NewEncoder(w).Encode(book)
}
