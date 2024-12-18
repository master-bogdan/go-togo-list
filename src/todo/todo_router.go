package todo

import (
	todo_handlers "go-todo-list/src/todo/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/todos", todo_handlers.GetTodos)
	app.Post("/api/todos", todo_handlers.CreateTodo)
	app.Patch("/api/todos/:id", todo_handlers.UpdateTodo)
	app.Delete("/api/todos/:id", todo_handlers.DeleteTodo)
}
