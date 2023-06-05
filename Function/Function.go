/*---------------------------------------------------------------------------------
  Syntax
  Function types
  Named return values
  User defined function types
--------------------------------------------------------------------------------*/

package main

//-------------------------------------------------------------------------------
// Syntax

//func function_name(arguman type) return type {
//  function body
//}

//-------------------------------------------------------------------------------

func main() {

	//-------------------------------------------------------------------------------
	// Function types

	//func ()
	//func (x int) int
	//func (a, _ int, z float32) bool
	//func (a, b int, z float32) (bool)
	//func (prefix string, values ...int)
	//func (a, b int, z float64, opt ...interface {}) (success bool)
	//func (int, int, float64) (float64, *[]int)
	//func (n int)
	//func (p *T)
	//func (p *T) *T

	//-------------------------------------------------------------------------------
	// Named return values

	//area := ExampleOne(1.8, 2.3)
	//fmt.Println("Area %f Perimeter %f", area)

	//---------------------------------------------------------------------------------
	// User defined function types

	// type funcSum func(a int, b int) int

	// var a funcSum = func(a int, b int) int {
	// 	return a + b
	// }

	// s := a(5, 6)
	// fmt.Println("Sum", s)

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Named return values

func ExampleOne(length, width float64) (area float64) {
	area = length * width
	//return area //True
	return //1-no explicit return value and //2-no need declaration variable area
}

//-------------------------------------------------------------------------------
