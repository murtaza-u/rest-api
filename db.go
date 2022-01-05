package main

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

func initDB() (db *gorm.DB) {
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Panic(err)
    }
    db.AutoMigrate(&User{})
    return db
}

func createUser(name, dob, address, description string) error {
    user := User{
        Name: name,
        Dob: dob,
        Address: address,
        Description: description,
    }
    db.Create(&user)

    return nil
}
