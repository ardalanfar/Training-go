/*---------------------------------------------------------------------------------
  Description
  Creating pointer using new()
  Default value pointer using new()
  Creating slice pointer using new()
  Creating map pointer using new()
  Creating channel pointer using new()
--------------------------------------------------------------------------------*/

package main

type emp struct {
	name string
	id   int
}

func main() {

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	// Description

	//1- new(T) always returns a *T pointing to an initialised T.
	//2- Using new() to construct a pointer to a slice, map, or channel.
	//3- new() only returns pointers to initialised memory.
	//4- it returns a pointer to a newly allocated zero value of type T.
	//5- A built-in function that allocates memory, but it does not initialize the memory.
	//6-

	//---------------------------------------------------------------------------------
	// Creating pointer using new()

	// size := new(int)
	// fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)

	//*size = 85
	//fmt.Println("New size value is", *size)

	//---------------------------------------------------------------------------------
	// Default value pointer using new()

	//p := new(emp)           //type *emp

	//fmt.Println(p == nil)   //false
	//fmt.Println(*p == nil)  //true
	//fmt.Println(p)   //Default value

	//---------------------------------------------------------------------------------
	// Creating slice pointer using new()

	//s := new([]string)       //slice

	//fmt.Println(len(*s))     //0
	//fmt.Println(*s == nil)   //true

	//---------------------------------------------------------------------------------
	// Creating map pointer using new()

	//m := new(map[string]int)   //map

	//fmt.Println(m == nil)      //false
	//fmt.Println(*m == nil)     //true

	//---------------------------------------------------------------------------------
	// Creating channel pointer using new()

	//c := new(chan int)      //channel

	//fmt.Println(c == nil)   //false
	//fmt.Println(*c == nil)  //true

	//---------------------------------------------------------------------------------

}
