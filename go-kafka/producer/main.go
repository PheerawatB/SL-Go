package main

import (
	"producer/controllers"
	"producer/services"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewOpenAccountController(accountService)

	app := fiber.New()
	app.Post("/openAccounts", accountController.OpenAccount)
	app.Post("/depositFund", accountController.DepositFunds)
	app.Post("/withdrawFund", accountController.WithdrawFunds)
	app.Post("/closeAccount", accountController.CloseAccount)
	app.Listen(":8080")
}
