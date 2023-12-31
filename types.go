package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateAccountUpdateRequest struct {
	ID      int   `json:"id"`
	Balance int64 `json:"balance"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Number     int64     `json:"number"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"created_at"`
	InsertedAt time.Time `json:"inserted_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now(),
	}
}
