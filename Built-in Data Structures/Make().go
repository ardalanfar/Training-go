/*---------------------------------------------------------------------------------
  Description
  Creating Slice using make()
  Creating Map using make()
  Creating Channel using make()
  Copy by function Copy() and use make()
--------------------------------------------------------------------------------*/

package main

func main() {

	//---------------------------------------------------------------------------------
	// Description

	//1- built-in function that is used to initialize (slices, maps, channels).
	//2- unlike the new() function, make() does not return a pointer.
	//3-

	//---------------------------------------------------------------------------------
	// Creating Slice using make()

	//var d = make([]string, 5)
	//a := make([]string, 3)
	//b := make([]int, 5, 5)
	//c := make([]string, len(b))
	//d := make([]int)   //Error
	//fmt.Println(c)

	//------------------------------------
	// Creating 2Dimensional slice using make()

	// i, j := 3,3
	// a := make([][]int, i)
	// for i := range a {
	// 	a[i] = make([]int, j)
	// }

	//---------------------------------------------------------------------------------
	// Creating Map using make()

	//var a = make(map[string]int)
	//b := make(map[string]int)
	//c := make(map[string]int, 3)

	//---------------------------------------------------------------------------------
	// Creating Channel using make()

	//var a = make(chan string)
	//b := make(chan int)

	//---------------------------------------------------------------------------------
	// Copy by function Copy() and use make()

	// a := make([]string, 3)
	// b := make([]string, 3)
	//or
	// var a []string   //slice
	// var b []string   //slice
	//or
	// var a [3]string  //array  //Error
	// var b [3]string  //array  //Error

	// a = append(a, "d", "e")
	// copy(a, b)  //Copy (arguments to copy must be slices)
	// fmt.Println("copy:", a)

	//---------------------------------------------------------------------------------

}
