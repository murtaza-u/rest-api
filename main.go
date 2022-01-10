package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
    app := fiber.New()

    ctx := context.Background()
    db := Connect()
    s := InitServe(ctx, db)

    controller := &Controller{s}

    app.Post("/user", controller.Create)
    app.Get("/user", controller.GetAll)
    app.Get("/user/:id", controller.Get)
    app.Put("/user/:id", controller.Update)
    app.Delete("/user/:id", controller.Delete)

    log.Fatal(app.Listen(":5000"))
}
