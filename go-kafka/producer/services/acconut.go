package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) (string, error)
	DepositFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommand) error
	ClosedAccount(command commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(producer EventProducer) AccountService {
	return &accountService{producer}
}

func (s *accountService) OpenAccount(command commands.OpenAccountCommand) (ID string, err error) {
	if command.AccountHandler == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}
	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolders: command.AccountHandler,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}
	log.Printf("^#v", event)
	return event.ID, s.eventProducer.Produce(event)
}

func (s *accountService) DepositFund(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("^#v", event)
	return s.eventProducer.Produce(event)
}

func (s *accountService) WithdrawFund(command commands.WithdrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}
	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}
	log.Printf("^#v", event)
	return s.eventProducer.Produce(event)
}

func (s *accountService) ClosedAccount(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}
	event := events.ClosedAccountEvent{
		ID: command.ID,
	}
	log.Printf("^#v", event)
	return s.eventProducer.Produce(event)
}
