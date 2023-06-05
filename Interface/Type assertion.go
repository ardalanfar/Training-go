/*---------------------------------------------------------------------------------
  Discretion
  Type assertion (syntax i.(T))
--------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//-------------------------------------------------------------------------------
	// Discretion

	//1- Type assertion is used to extract the underlying value of the interface.
	//2- i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
	//3-

	//-------------------------------------------------------------------------------
	// Type assertion (syntax i.(T))

	//------------------------------------ Expl_1
	//String

	// var s interface{} = 56
	// Assert(s)

	//------------------------------------ Expl_2
	//Int

	// var i interface{} = "Steven Paul"
	// Assert(i)

	//------------------------------------ Expl_3
	//Struct

	// type person struct {
	// 	name   string
	// 	family string
	// }

	// new := person{
	// 	name:   "ali",
	// 	family: "salimi",
	// }

	// var i interface{} = new
	// val, ok := i.(person)
	// name := i.(person).name

	// fmt.Println(val, ok)
	// fmt.Println(val.family, ok)
	// fmt.Println(name)

	//-------------------------------------------------------------------------------
}

//-------------------------------------------------------------------------------
// Type assertion (syntax i.(T))

//------------------------------------ (Expl_1, Expl_2)

func Assert(i interface{}) {
	//val = i.(int)  //PANIC
	//val, ok := i.(float64)  //OUTPUT: 0, false
	//val, ok := i.(string)
	val, ok := i.(int)

	fmt.Println(val, ok)
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

//-------------------------------------------------------------------------------
