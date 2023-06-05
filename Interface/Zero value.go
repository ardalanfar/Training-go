/*---------------------------------------------------------------------------------
  Zero value of Interface
--------------------------------------------------------------------------------*/
package main

func main() {

	//-------------------------------------------------------------------------------
	// Zero value of Interface
	//1- The zero value of a interface is nil.
	//2- If we try to call a method on the nil interface, the program will panic since the nil interface neither has a underlying value nor a concrete type.
	//3-

	// var d1 Describer2

	// if d1 == nil {
	// 	fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	// }

	// d1.Describe() //Error

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Zero value of Interface

type Describer2 interface {
	Describe()
}

//-------------------------------------------------------------------------------
