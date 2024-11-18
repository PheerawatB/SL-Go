package services

import "errors"

var (
	ErrZeroAmount = errors.New("Amount must be greater than zero")
	ErrRepository = errors.New("Repository error")
)
