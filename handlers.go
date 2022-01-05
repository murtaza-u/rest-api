package main

import (
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    user, err = getUserByName(user.Name)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    resp := User {
        ID: user.ID,
        Name: user.Name,
        Dob: user.Dob,
        Address: user.Address,
        Description: user.Description,
    }

    if err := json.NewEncoder(w).Encode(resp); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    err = deleteUser(user.Name)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
