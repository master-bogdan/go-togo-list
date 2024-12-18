package main

import (
	"go-todo-list/src/config"
	"go-todo-list/src/todo"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var PORT = config.LoadConfig().PORT
	var app = fiber.New()

	todo.SetupRoutes(app)

	log.Fatal(app.Listen(":" + PORT))
}
