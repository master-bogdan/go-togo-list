package todo_handlers

import (
	"context"
	"go-todo-list/src/db"
	todo_model "go-todo-list/src/todo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []todo_model.Todo

	var cursor, err = db.Repository().Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo todo_model.Todo

		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)
}
