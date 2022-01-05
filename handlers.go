package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    user, err := getUserByID(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    if err := json.NewEncoder(w).Encode(user); err != nil {
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
    params := mux.Vars(r)

    user, err := getUserByID(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    db.Save(&user)
    w.WriteHeader(http.StatusOK)
}

func delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := deleteUser(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
