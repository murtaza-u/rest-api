package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// - Fetches user by ID from the database.
// - Returns a json response back.
// - Method `GET`
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

// - Creates a new user and add him/her in the Database.
// - Requires a json request body as such,
// - Method `POST`
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

// - Updates users data and saves it into the database.
// - Requires a json request body as mentioned above.
// - Method `PUT`
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

    // finally save the updated user to the database
    db.Save(&user)

    w.WriteHeader(http.StatusOK)
}

// - Deletes user by it ID.
// - All the user's data is deleted from the database as well.
func delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    err := deleteUser(params["id"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
