/*---------------------------------------------------------------------------------
  Examples variadic function
  Examples Slice arguments
  When we can achieve the same functionality using slices?
  Passing a slice to a variadic function
  Modifying a slice inside a variadic function (by reference)
---------------------------------------------------------------------------------*/

package main

import (
	"fmt"
)

func main() {

	//-------------------------------------------------------------------------------
	// Examples variadic function

	//FindVariadic(89, 89, 90)
	//FindVariadic(45, 56, 67, 45, 90, 109)
	//FindVariadic(78, 38, 56, 98)
	//FindVariadic(87)

	//-------------------------------------------------------------------------------
	// Examples Slice arguments

	// FindSlice(89, []int{89, 90, 95})
	// FindSlice(45, []int{56, 67, 45, 90, 109})
	// FindSlice(78, []int{38, 56, 98})
	// FindSlice(87, []int{})

	//-------------------------------------------------------------------------------
	// When we can achieve the same functionality using slices?

	//1- There is no need to create a slice during each function call.
	//2- I personally feel that the program with variadic functions is more readable than the once with slices.
	//3- This line can just be FindVariadic(87) when variadic function is used. but in slice aguman should send a slice empty FindSlice(87, []int{}).
	//4-

	//-------------------------------------------------------------------------------
	// Passing a slice to a variadic function
	//the slice is directly passed to the function without a new slice being created.

	// nums := []int{89, 90, 95}
	// FindVariadic(89, nums...)

	//-------------------------------------------------------------------------------
	// Modifying a slice inside a variadic function (by reference)

	//-----------------------------------
	// str := []string{"hello", "world"}
	// Change(str...)
	// fmt.Println(str)

	//-----------------------------------
	// str := []string{"hello", "world"}
	// ChangeAndAppend(str...)
	// fmt.Println(str)

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Examples variadic function

func FindVariadic(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//-------------------------------------------------------------------------------
// Examples Slice arguments

func FindSlice(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//-------------------------------------------------------------------------------
// Modifying a slice inside a variadic function (by reference)

func Change(s ...string) {
	s[0] = "Go"
}

func ChangeAndAppend(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

//-------------------------------------------------------------------------------
