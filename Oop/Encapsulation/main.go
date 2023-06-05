/*---------------------------------------------------------------------------------
Encapsulation : by method
---------------------------------------------------------------------------------*/

package main

import "fmt"

type Person struct {
	Name string // exported field
	Age  int    // non-exported field
}

func main() {
	me := Person{Name: "mohoamad"}

	me.SetAndGetAge(36)
	me.GeTAge() //Error

	fmt.Println(me.Name)
	fmt.Println(me.Age) //Error
}

func (p Person) SetAndGetAge(age int) {
	p.Age = age
	fmt.Println(p.Age)
	fmt.Println(p.Name)
}

func (p Person) GeTAge() {
	fmt.Println(p.Age)
	fmt.Println(p.Name)
}
