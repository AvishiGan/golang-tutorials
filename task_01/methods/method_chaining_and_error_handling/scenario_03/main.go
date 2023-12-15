package main

import (
	"fmt"
	"net/mail"
)

type User struct{}

func (u User) validateEmail(email string) bool {
	// _ -> blank identifier
	// Used when want to ignore the value being returned
	// In this case the Address object returned by ParseAddress
	_, err := mail.ParseAddress(email)

	// No error
	return err == nil
}

func main() {
	user := User{}

	fmt.Println(user.validateEmail("avishiganepola@gmail.com"))
	fmt.Println(user.validateEmail("avishiganepolagmail.com"))
}
