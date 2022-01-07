package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "testing"
)

func TestCreate(t *testing.T) {
    // dummy data
    user := User {
        Name: "Murtaza Udaipurwala",
        Dob: "2002-12-09",
        Address: "India",
        Description: "Hey, I am a computer genius",
    }

    // preparing our json request body
    var buff bytes.Buffer
    if err := json.NewEncoder(&buff).Encode(user); err != nil {
        log.Panic(err)
    }

    // instantiating a new http request
    r, err := http.NewRequest("POST", "http://localhost:5000/create", &buff)
    if err != nil {
        log.Panic(err)
    }

    // setting appropriate headers
    r.Header.Set("Content-Type", "application/json")

    // Dispatching request
    client := &http.Client{}
    response, err := client.Do(r)
    if err != nil {
        log.Panic(err)
    }
    defer response.Body.Close()

    // assertions
    status := response.StatusCode
    if status != http.StatusCreated {
        t.Errorf("Handler returned incorrect status code: %v. Expected: %v", status, http.StatusCreated)
    }
}

func TestGet(t *testing.T) {
    // instantiating a new http request
    r, err := http.NewRequest("GET", "http://localhost:5000/get/1", nil)
    if err != nil {
        log.Panic(err)
    }

    // Dispatching request
    client := &http.Client{}
    response, err := client.Do(r)
    if err != nil {
        log.Panic(err)
    }
    defer response.Body.Close()

    // assertions
    status := response.StatusCode
    if status != http.StatusOK {
        t.Errorf("Handler returned incorrect status code: %v. Expected: %v", status, http.StatusOK)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        t.Fail()
    }
    t.Log(string(body))
}

func TestUpdate(t *testing.T) {
    // dummy data
    user := User {
        Description: "Hey, I am a math genius",
    }

    // preparing our json request body
    var buff bytes.Buffer
    if err := json.NewEncoder(&buff).Encode(user); err != nil {
        log.Panic(err)
    }

    // instantiating a new http request
    r, err := http.NewRequest("PUT", "http://localhost:5000/update/1", &buff)
    if err != nil {
        log.Panic(err)
    }

    // setting appropriate headers
    r.Header.Set("Content-Type", "application/json")

    // Dispatching request
    client := &http.Client{}
    response, err := client.Do(r)
    if err != nil {
        log.Panic(err)
    }
    defer response.Body.Close()

    // assertions
    status := response.StatusCode
    if status != http.StatusOK {
        t.Errorf("Handler returned incorrect status code: %v. Expected: %v", status, http.StatusOK)
    }
}

func TestDelete(t *testing.T) {
    // instantiating a new http request
    r, err := http.NewRequest("DELETE", "http://localhost:5000/delete/1", nil)
    if err != nil {
        log.Panic(err)
    }

    // setting appropriate headers
    r.Header.Set("Content-Type", "application/json")

    // Dispatching request
    client := &http.Client{}
    response, err := client.Do(r)
    if err != nil {
        log.Panic(err)
    }
    defer response.Body.Close()

    // assertions
    status := response.StatusCode
    if status != http.StatusOK {
        t.Errorf("Handler returned incorrect status code: %v. Expected: %v", status, http.StatusOK)
    }
}
