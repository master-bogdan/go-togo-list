package todo_handlers

import (
	"context"
	"go-todo-list/src/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *fiber.Ctx) error {
	var id = c.Params("id")
	var objectId, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	var filter = bson.M{"_id": objectId}

	_, err = db.Repository().DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}
