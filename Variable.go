/*---------------------------------------------------------------------------------
  Declaring Variable
  Declaring Constants

  Default value
  Again declaration value
  Static Type to change value
  Overwrite variable
  Swapping values
--------------------------------------------------------------------------------*/

package main

import (
	_ "math"
	_ "reflect" //(Typeof)
	_ "unsafe"  //(Sizeof)
)

//-------------------------------------------------------------------------------
// Declaring Variable

//------------------------------------
// Group Declaring (only in main function)

// var (
// 	name  string
// 	id    int
// 	flt   float32
// 	byt   byte
// 	bol   bool
// 	cmp   complex64
// 	run   rune
// 	d, e, f int
// 	h2, i2 int = 7, 8
// 	number = 21
// 	ft     = 21.10
// 	str    = "salam"
// 	bl     = true
// 	cmplx  = 23 + 34i
// )

//------------------------------------
// Var

//var a int

//------------------------------------
// Var and Initialize

//var str string = "shorts"

//------------------------------------
// Var Multi-Variable Initialize

//var width, height int
//var width, height int = 5, 6
//var h, i, g = 3, 4, "foo"

//------------------------------------
// Short variable declarations
//<name> := <initializer>

// a := 10
// a, b := 10, 20

//-------------------------------------------------------------------------------
// Declaring Constants

//------------------------------------
// Constants

// const PRODUCT string = "Canada"
// const Y = 55

// //const PRICE int  //Error

//------------------------------------
// Multi-Constants

// const X, Z = 20, "salam"

//------------------------------------
// Group Declaring

// const (
// 	PRODUCT2 = "Mobile"
// 	QUANTITY = 50
// 	PRICE    = 50.50
// 	STOCK    = true
// )

//------------------------------------
// Constant Generator iota
// The constant generator iota is used for declaring a sequence of related values that use incrementing numbers without the need to explicitly type each one of them.

// const (
// 	Zero int = iota
// 	One
// 	Two
// )

// EQUAL TO
// const (
// 	Zero = 0
// 	One  = 1
// 	Two  = 2
// )

//-------------------------------------------------------------------------------

func main() {

	//-----------------------------------------------------------------------------
	// Default value
	// Local and global variables are initialized to their default value, which is 0.

	// var num int
	// var flt float32
	// var bol bool
	// var str string

	// fmt.Println("Default int:", num, " Default bool:", bol, " Default float:", flt, " Default string:", str)

	//----------------------------------------------------------------------------
	// Again declaration value

	//x := 10    //Declare variable x
	//fmt.Println("x is", x)

	//x := 20   //Error again Declaration value
	//x = 30    //True
	//fmt.Println("x is", x)

	//-----------------------------------------------------------------------------
	// Static Type to change value

	//var a int
	//b := 10

	//b = 13.5  //Error
	//a = 14.5  //Error
	//a = 20    //True

	//-----------------------------------------------------------------------------
	// Overwrite variable

	// var (
	// 	param string
	// )

	// param = "invalid"
	// if 2 > 0 {
	// 	param = "valid"
	// } else if 2 > 1 {
	// 	param = "valid ^ 2"
	// }
	// param = "invalid"

	// fmt.Println(param)

	//-----------------------------------------------------------------------------
	// Swapping values

	// i := 0
	// j := 1

	// i, j = j, i

	//-----------------------------------------------------------------------------

}
