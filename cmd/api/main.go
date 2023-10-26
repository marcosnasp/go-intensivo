package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosnasp/go-intensivo/internal/entity"
)

const server_port = 8080

func main() {

	app := fiber.New()
	
	fmt.Printf("Server is running on port: %d\n", server_port)

	app.Get("/order", Order)
	app.Post("/order", Order)

	app.Listen(fmt.Sprintf(":%v", server_port))
}

func Order(c *fiber.Ctx) error {
	order, err := entity.NewOrder("123", 1000, 10)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).Accepts("application/json")
		return c.JSON(err.Error())
	}
	order.CalculateFinalPrice()
	c.Status(fiber.StatusOK).Accepts("application/json")
	c.JSON(order)
	return nil
}