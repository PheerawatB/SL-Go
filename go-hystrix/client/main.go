package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api", api)
	app.Listen(":3031")
}

func init() {
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 2,
		SleepWindow:            15000,
		ErrorPercentThreshold:  50,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":3330", hystrixStreamHandler)
}

func api(c *fiber.Ctx) error {
	output := make(chan string)
	errorChan := make(chan error)

	hystrix.Go("api", func() error {
		res, err := http.Get("http://localhost:3030/api")
		if err != nil {
			errorChan <- err
			return nil
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			errorChan <- err
			return nil
		}
		msg := string(data)
		fmt.Println(msg)
		output <- msg
		return nil
	}, func(err error) error {
		errorChan <- err
		return nil
	})

	select {
	case out := <-output:
		return c.SendString(out)
	case err := <-errorChan:
		return c.Status(500).SendString(err.Error())
	}
}
