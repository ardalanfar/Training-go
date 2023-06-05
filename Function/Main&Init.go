/*---------------------------------------------------------------------------------
  Description
  Main
  Init
---------------------------------------------------------------------------------*/

package main

import "fmt"

//----------------------------------------------------------------------------
// Description

//1- The Go language reserve two functions for special purpose and the functions are main() and init() function.
//2- init() function is just like the main function, does not take any argument not return anything.
//3- This function is present in every package and this function is called when the package is initialized.
//4- The main purpose of the init() function is to initialize the global variables that cannot be initialized in the global context.
//5- so it does not depend to main() function.
//6- An init function is a function used to initialize the state of an application.
//7-

//----------------------------------------------------------------------------
// Main

func main() {
	fmt.Println("main")
}

//----------------------------------------------------------------------------
// Init

func init() {
	fmt.Println("init")
}

//----------------------------------------------------------------------------
