/*---------------------------------------------------------------------------------
  Discretion
  Type switch
--------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"reflect"
)

func main() {

	//-------------------------------------------------------------------------------
	// Discretion

	//It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.

	//1- can not use of .(type) outside type switch.
	//2-

	//-------------------------------------------------------------------------------
	// Type switch

	//----------------------------------- Expl_1
	// findType("Naveen")
	// findType(77)
	// findType(89.98)

	//----------------------------------- Expl_2
	//findType2("Naveen")

	//----------------------------------- Expl_3
	// p := Person2{
	// 	name: "Naveen",
	// 	age:  25,
	// }

	// findType2(p)

	//-------------------------------------------------------------------------------
}

//-------------------------------------------------------------------------------
// Type switch

//----------------------------------- Expl_1

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

//----------------------------------- Expl_2, Expl_3

type Describer interface {
	Describe()
}

type Person2 struct {
	name string
	age  int
}

func (p Person2) Describe() {
	fmt.Printf("%s is %d years old \n", p.name, p.age)
}

func findType2(i interface{}) {
	fmt.Println(reflect.TypeOf(i))

	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

//-------------------------------------------------------------------------------
