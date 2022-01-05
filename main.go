package main

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
    id          int    `gorm:"primarykey"`
    name        string
    dob         string
    address     string
    description string
    gorm.Model
}

func main() {
    http.HandleFunc("/get", get)
    http.HandleFunc("/create", create)
    http.HandleFunc("/update", update)
    http.HandleFunc("/delete", delete)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
