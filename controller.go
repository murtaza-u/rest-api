package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
    S Service
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
    body := new(User)
    err := ctx.BodyParser(body)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    body.ID = ""

    err = c.S.Create(*body)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            ctx.Status(fiber.StatusBadRequest)
        } else {
            ctx.Status(fiber.StatusInternalServerError)
        }
        return err
    }

    return ctx.Status(fiber.StatusCreated).SendString("User created successfully")
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
    users := make([]User, 0)
    err := c.S.GetAll(&users)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    return ctx.Status(fiber.StatusOK).JSON(&users)
}

func (c *Controller) Get(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    _id, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    user := new(User)
    err = c.S.Get(_id, user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            ctx.Status(fiber.StatusBadRequest)
        } else {
            ctx.Status(fiber.StatusInternalServerError)
        }
        return err
    }

    return ctx.Status(fiber.StatusOK).JSON(user)
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    _id, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    body := new(User)
    err = ctx.BodyParser(body)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    user := new(User)
    err = c.S.Get(_id, user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            ctx.Status(fiber.StatusBadRequest)
        } else {
            ctx.Status(fiber.StatusInternalServerError)
        }
        return err
    }

    username := body.Username
    passwod := body.Password

    if len(username) == 0 {
        username = user.Username
    }

    if len(passwod) == 0 {
        passwod = user.Password
    }

    err = c.S.Update(_id, username, passwod)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    return ctx.Status(fiber.StatusOK).SendString("User updated successfully")
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
    id := ctx.Params("id")
    _id, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        ctx.Status(fiber.StatusInternalServerError)
        return err
    }

    err = c.S.Delete(_id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            ctx.Status(fiber.StatusBadRequest)
        } else {
            ctx.Status(fiber.StatusInternalServerError)
        }

        return err
    }

    return ctx.Status(fiber.StatusOK).SendString("User deleted successfully")
}
