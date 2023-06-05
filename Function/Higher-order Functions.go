/*---------------------------------------------------------------------------------
  Higher-order functions
  Passing functions as arguments to other functions
  Returning functions from other functions
--------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//---------------------------------------------------------------------------------
	// Higher-order functions

	// a function which does at least one of the following
	//1- takes one or more functions as arguments.
	//2- returns a function as its result.

	//---------------------------------------------------------------------------------
	// Passing functions as arguments to other functions

	// f := func(a, b int) int {
	// 	return a + b
	// }

	// SimpleOne(f)

	//---------------------------------------------------------------------------------
	// Returning functions from other functions

	// s := SimpleTwo()
	// fmt.Println(s(60, 7))

	//---------------------------------------------------------------------------------

}

//---------------------------------------------------------------------------------
// Passing functions as arguments to other functions

func SimpleOne(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

//---------------------------------------------------------------------------------
// Returning functions from other functions

func SimpleTwo() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

//---------------------------------------------------------------------------------
