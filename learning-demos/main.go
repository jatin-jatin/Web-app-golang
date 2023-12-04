package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Jatin",
		LastName:    "Lachhwani",
		PhoneNumber: "9999888888",
	}
	log.Println(user)
}
