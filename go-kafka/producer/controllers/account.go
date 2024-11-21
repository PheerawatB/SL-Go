package controllers

import (
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	service services.AccountService
}

func NewOpenAccountController(service services.AccountService) accountController {
	return accountController{service: service}
}

func (c accountController) OpenAccount(ctx *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}
	if err := ctx.BodyParser(&command); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	id, err := c.service.OpenAccount(command)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "open account successfully",
		"id":      id,
	})
}

func (c accountController) DepositFunds(ctx *fiber.Ctx) error {
	command := commands.DepositFundCommand{}
	if err := ctx.BodyParser(&command); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := c.service.DepositFund(command)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "deposit fund successfully",
	})
}

func (c accountController) WithdrawFunds(ctx *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}
	if err := ctx.BodyParser(&command); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := c.service.WithdrawFund(command)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "withdraw fund successfully",
	})
}

func (c accountController) CloseAccount(ctx *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}
	if err := ctx.BodyParser(&command); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := c.service.ClosedAccount(command)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "close account successfully",
	})
}
