package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	PhoneNo  string
}

type UserType struct {
	gorm.Model
	Uid         User
	Name        string
	Description string
}

type Shop struct {
	gorm.Model
	Name        string
	Uid         User
	Tid         UserType
	Description string
}
