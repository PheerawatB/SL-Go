package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	couter := 0
	app.Get("/api", func(c *fiber.Ctx) error {
		couter++

		if couter >= 5 && couter <= 10 {
			time.Sleep(time.Second)
		}
		msg := fmt.Sprintf("counter: %d", couter)
		fmt.Println(msg)
		return c.SendString(msg)
	})

	app.Listen(":3030")
}
