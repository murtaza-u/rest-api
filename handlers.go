package main

import (
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {

}

func create(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    createUser(user.Name, user.Dob, user.Address, user.Description)
    w.WriteHeader(http.StatusCreated)
}

func update(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}
