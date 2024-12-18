package todo_handlers

import (
	"context"
	"go-todo-list/src/db"
	todo_model "go-todo-list/src/todo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(c *fiber.Ctx) error {
	var todo = new(todo_model.Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	var result, err = db.Repository().InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}
