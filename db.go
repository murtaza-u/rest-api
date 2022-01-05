package main

import (
	"errors"
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

func getUserByName(name string) (User, error) {
    var user User
    db.Where("name = ?", name).Find(&user)
    if len(user.Name) == 0 {
        return user, errors.New("user not found")
    }
    return user, nil
}
