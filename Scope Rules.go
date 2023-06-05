/*---------------------------------------------------------------------------------
  Global variables (out main())
  Local variables (in main())
  LifeTime
--------------------------------------------------------------------------------*/
package main

// Global variables (out main())
var a int

func main() {

	//---------------------------------------------------------------------------------
	// Local variables (in main())
	// var m int

	// Actual initialization
	// a = 10

	// Short variable declaration
	// v := 2
	// fmt.Println(a, m, v)

	//---------------------------------------------------------------------------------
	// LifeTime

	//----------------------------------- Expl_1 (block)
	// x := 10
	// for i := 0; i < x; i++ {
	// 	x := 1
	// 	if x == 1 {
	// 		x := x + 1
	// 		fmt.Printf("%x", x)
	// 	}
	// }

	//----------------------------------- Expl_2 (function)

	// var cwd string
	// func init() {
	//      cwd, err := os.Getwd()  // compile error: unused: cwd
	// }

	//---------------------------------------------------------------------------------
}
