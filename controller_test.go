package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/murtaza-udaipurwala/rest-api"
	"github.com/murtaza-udaipurwala/rest-api/mocks"
)

var _ = Describe("Controller", func() {
    dbMock := mocks.Service{}
    controller := &main.Controller{&dbMock}

    Describe("create user", func () {
        BeforeEach(func() {
            dbMock.On("Create", mock.Anything).Return(nil)
        })

        It("should create a user", func () {
            user := main.User {
                ID: "",
                Username: "Bob",
                Password: "secret",
            }
            err := controller.S.Create(user)
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Describe("get all users", func () {
        BeforeEach(func() {
            dbMock.On("GetAll", mock.Anything).Return(nil)
        })

        It("should return all users", func () {
            users := make([]main.User, 0)
            err := controller.S.GetAll(&users)
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Describe("get user by ID", func () {
        BeforeEach(func() {
            dbMock.On("Get", mock.Anything, mock.Anything).Return(nil)
        })

        It("should return a user", func () {
            id, _ := primitive.ObjectIDFromHex("61dc3aa24a3dbcf214a3fa78")
            err := controller.S.Get(id, nil)
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Describe("update user", func () {
        BeforeEach(func() {
            dbMock.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
        })

        It("should update the user", func () {
            id, _ := primitive.ObjectIDFromHex("61dc3aa24a3dbcf214a3fa78")
            err := controller.S.Update(id, "", "")
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Describe("delete user", func () {
        BeforeEach(func() {
            dbMock.On("Delete", mock.Anything).Return(nil)
        })

        It("should delete the user", func () {
            id, _ := primitive.ObjectIDFromHex("61dc3aa24a3dbcf214a3fa78")
            err := controller.S.Delete(id)
            Expect(err).NotTo(HaveOccurred())
        })
    })
})
