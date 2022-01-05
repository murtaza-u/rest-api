package main

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// initialise DB
func initDB() (db *gorm.DB) {
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Panic(err)
    }

    // auto migrate tables, columns, rows
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

    // creates a new row in the database
    db.Create(&user)

    return nil
}

func getUserByID(id string) (User, error) {
    var user User
    db.Where("id = ?", id).Find(&user)

    // incase the user does not exist, return an error
    if len(user.Name) == 0 {
        return user, errors.New("user not found")
    }
    return user, nil
}

func deleteUser(id string) error {
    user, err := getUserByID(id)
    if err != nil {
        return err
    }

    // deletes specified row from the database
    db.Delete(&user, user.ID)
    return nil
}
