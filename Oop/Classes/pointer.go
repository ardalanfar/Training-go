package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

//Constructor
func New(id int, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

//Method
func (u *User) PrintName() {
	fmt.Println(u.name, u.id)
}

func main() {
	user := New(10, "ahmad")
	user.PrintName()
}
