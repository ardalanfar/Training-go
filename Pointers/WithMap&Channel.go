/*---------------------------------------------------------------------------------
  Description
  Edit map by function
---------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//----------------------------------------------------------------------------
	// Description

	//1- Maps and channels are already references to the internal structure.
	//2- Consequently, a function/method that accepts a map or a channel can modify it, even if the parameter is not a pointer type.
	//3-
	
	//----------------------------------------------------------------------------
	// Edit map by function

	// myMap := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	// editMaps(myMap)
	// fmt.Println(myMap)

	//----------------------------------------------------------------------------

}

//----------------------------------------------------------------------------
// Edit map by function

func editMaps(c map[string]int) {
	c["d"] = 4
	fmt.Println(c)
}

//----------------------------------------------------------------------------
