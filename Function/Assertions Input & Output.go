/*---------------------------------------------------------------------------------
  Return interface() in function
  Passing interface() in function
--------------------------------------------------------------------------------*/

package main

import "fmt"

type Person struct {
	Name string
	Tell int
}

//-------------------------------------------------------------------------------

func main() {

	//-------------------------------------------------------------------------------
	// Return interface() in function

	// a := TestOne()
	// p := a.(Person)
	// n := a.(Person).Name
	// t := a.(Person).Tell
	// fmt.Println(n, t)

	//-------------------------------------------------------------------------------
	// Passing interface() in function

	// TestTwo("salam donya")

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Return interface() in function

func TestOne() interface{} {
	person := Person{
		Name: "ali",
		Tell: 1123,
	}
	return person
}

//-------------------------------------------------------------------------------
// Passing interface() in function

func TestTwo(n interface{}) {
	fmt.Println(n)
}

//-------------------------------------------------------------------------------
