package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
}

//Constructor
func New(first string, last string) User {
	e := User{
		firstName: first,
		lastName:  last,
	}
	return e
}

//Method
func (u User) LeavesRemaining() {
	fmt.Printf("%s %s has \n", u.firstName, u.lastName)
}

func main() {
	user := New("Sam", "Adolf")
	user.LeavesRemaining()
}
