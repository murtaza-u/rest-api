package main

import (
    "log"
    "net/http"

    "gorm.io/gorm"
)

type User struct {
    ID          int    `gorm:"primarykey"`
    Name        string
    Dob         string
    Address     string
    Description string
    gorm.Model
}

func main() {
    db = initDB()

    http.HandleFunc("/get", get)
    http.HandleFunc("/create", create)
    http.HandleFunc("/update", update)
    http.HandleFunc("/delete", delete)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
