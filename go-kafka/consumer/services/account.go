package services

import (
	"comsumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) *accountEventHandler {
	return &accountEventHandler{accountRepo: accountRepo}
}

func (h *accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name(): //OpenAccountEvent
		event := &events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println("Error unmarshalling event:", err)
			return
		}
		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolders,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBakance,
		}
		err = h.accountRepo.Save(bankAccount)
		if err != nil {
			log.Println("Error saving account:", err)
			return
		}
	case reflect.TypeOf(events.DepositFundEvent{}).Name(): //DepositFundEvent
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println("Error unmarshalling event:", err)
			return
		}
		account, err := h.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println("Error finding account:", err)
			return
		}
		account.Balance += event.Amount
		err = h.accountRepo.Save(account)
		if err != nil {
			log.Println("Error saving account:", err)
			return
		}
	case reflect.TypeOf(events.WithdrawFundEvent{}).Name(): //WithdrawFundEvent
		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println("Error unmarshalling event:", err)
			return
		}
		account, err := h.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println("Error finding account:", err)
			return
		}
		account.Balance -= event.Amount
		err = h.accountRepo.Save(account)
		if err != nil {
			log.Println("Error saving account:", err)
			return
		}
	case reflect.TypeOf(events.ClosedAccountEvent{}).Name(): //ClosedAccountEvent
		event := &events.ClosedAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println("Error unmarshalling event:", err)
			return
		}
		err = h.accountRepo.Delete(event.ID)
		if err != nil {
			log.Println("Error deleting account:", err)
			return
		}

	default:
		log.Println("no event handler:", topic)
	}
}
