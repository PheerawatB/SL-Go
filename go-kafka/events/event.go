package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
	reflect.TypeOf(DepositFundEvent{}).Name(),
	reflect.TypeOf(WithdrawFundEvent{}).Name(),
	reflect.TypeOf(ClosedAccountEvent{}).Name(),
}

type OpenAccountEvent struct {
	ID             string
	AccountHolders string
	AccountType    int
	OpeningBakance float64
}

type DepositFundEvent struct {
	ID     string
	Amount float64
}

type WithdrawFundEvent struct {
	ID     string
	Amount float64
}

type ClosedAccountEvent struct {
	ID string
}
