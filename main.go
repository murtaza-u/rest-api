package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
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

    http.HandleFunc("/get", get)
    http.HandleFunc("/create", create)
    http.HandleFunc("/update", update)
    http.HandleFunc("/delete", delete)

    fmt.Println("Listening on port :5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}
