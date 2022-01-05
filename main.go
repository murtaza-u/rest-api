package main

import (
    "fmt"
    "log"
    "net/http"

    "gorm.io/gorm"
    "github.com/gorilla/mux"
)

// the user model
type User struct {
    ID          int    `gorm:"primarykey"`  // unique id
    Name        string `json:"name"`        // username
    Dob         string `json:"dob"`         // DOB
    Address     string `json:"address"`     // address
    Description string `json:"description"` // description
    gorm.Model
}

func main() {
    // initalising database
    db = initDB()

    // initialising a new gorilla mux router
    r := mux.NewRouter()

    // functions to handle API calls
    r.HandleFunc("/get/{id}", get)
    r.HandleFunc("/create", create)
    r.HandleFunc("/update/{id}", update)
    r.HandleFunc("/delete/{id}", delete)

    // leave a debug message for the user
    fmt.Println("Listening on port :5000")

    // listen to new requests and server them accordingly
    // incase of an error, log it to stderr
    log.Fatal(http.ListenAndServe(":5000", r))
}
