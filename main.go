package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hellow world...ghe mng jai shree ram")

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading.env file")
	}
	PORT := os.Getenv("PORT")
	todos := []Todo{} // Changed from todo to todos to refer to the slice of todos
	//Get Todos
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})
	//Create Todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {

		todo := &Todo{}
		fmt.Println("are ky ahe", todo)
		err := c.BodyParser(todo)
		if err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}
		fmt.Println("after parsing", todo)
		todo.ID = len(todos) + 1     // Increment ID based on the number of todos
		todos = append(todos, *todo) // Append the todo to the todos slice
		fmt.Println("after parsing", todos)
		return c.Status(201).JSON(todo)

	})

	//Update the Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})
	//Delete Todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")
		fmt.Println("Delete Todo", id)

		for i, todon := range todos {
			if fmt.Sprint(todon.ID) == id {
				fmt.Println("Delete Todo", id)
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "true"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})

	log.Fatal(app.Listen(":" + PORT))
}
