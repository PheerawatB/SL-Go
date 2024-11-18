package main

import (
	"client/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()

	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	playerClient := services.NewPlayerClient(cc)
	playerServer := services.NewPlayerService(playerClient)

	// //err = playerServer.Hello("Frank")
	// // err = playerServer.Avg(1, 2, 3, 4, 5)
	// err = playerServer.Sum(1, 22, 33, 5, 4)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		// Call the Hello function from the gRPC client
		result, err := playerServer.CallHello(name)
		if err != nil {
			log.Printf("Failed to call Hello: %v", err)
			return c.Status(500).SendString("Failed to call Hello")
		}

		return c.JSON(fiber.Map{
			"message": result,
		})
	})

	// Start Fiber API server at port
	log.Fatal(app.Listen(":8081"))

}
