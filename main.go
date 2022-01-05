package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
    "github.com/gorilla/mux"
)

type User struct {
    ID          int    `gorm:"primarykey"`
    Name        string `json:"name"`
    Dob         string `json:"dob"`
    Address     string `json:"address"`
    Description string `json:"description"`
    gorm.Model
}

func main() {
    db = initDB()

    r := mux.NewRouter()

    r.HandleFunc("/get/{id}", get)
    r.HandleFunc("/create", create)
    r.HandleFunc("/update/{id}", update)
    r.HandleFunc("/delete/{id}", delete)

    fmt.Println("Listening on port :5000")
    log.Fatal(http.ListenAndServe(":5000", r))
}
