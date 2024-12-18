package todo_handlers

import (
	"context"
	"go-todo-list/src/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *fiber.Ctx) error {
	var id = c.Params("id")
	var objectId, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	var filter = bson.M{"_id": objectId}
	var update = bson.M{"$set": bson.M{"completed": true}}

	_, err = db.Repository().UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}
