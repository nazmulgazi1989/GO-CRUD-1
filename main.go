package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	han "github.com/tusharhow/handlers"
)

func main() {

	r := mux.NewRouter()
	p := ":8080"
	r.HandleFunc("/users", han.GetUsers).Methods("GET")
	r.HandleFunc("/createUser", han.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", han.GetUser).Methods("GET")
	r.HandleFunc("/deleteUser/{id}", han.DeleteUser).Methods("DELETE")
	r.HandleFunc("/updateUser/{id}", han.UpdateUser).Methods("PUT")

	fmt.Println("Server is running on port: ", strings.Split(p, ":")[1])

	han.Users = append(han.Users, han.User{ID: "12", Name: "Tushar Mah", Age: 30})
	han.Users = append(han.Users, han.User{ID: "33", Name: "T Mah", Age: 55})
	log.Fatal(http.ListenAndServe(p, r))

}
