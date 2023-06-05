/*---------------------------------------------------------------------------------
  Description
  Declaring pointers
  Default value of a pointer
  Dereferencing a pointer
  Pass value pointer to a function
  Return value pointer from a function
  Go does not support pointer arithmetic
  Double Pointer
---------------------------------------------------------------------------------*/

package main

//----------------------------------------------------------------------------
// Structure with pointer

type myName struct {
	Name string
}

//----------------------------------------------------------------------------

func main() {

	//----------------------------------------------------------------------------
	// Description

	//1- Use `&` to get the memory address of a non-pointer variable, called address-of operator.
	//2- And `*` to get the value of a pointer, which is called dereferencing the pointer.
	//3- Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are `nil`.
	//4-

	//----------------------------------------------------------------------------
	// Declaring pointers

	//----------------------------------- Expl_1
	//var p1 *int
	//p2 := &T{}

	// fmt.Printf("p1 address: %p\n", p1)
	// fmt.Printf("p2 address: %p\n", p2)

	//----------------------------------- Expl_2
	// a := 255
	// b := *a  //Error
	// b := &a  //True
	// var c *int = &a

	// fmt.Println("Type of a is ", *b)
	// fmt.Printf("Type of a is %T\n", c)
	// fmt.Println("address of b is", c)
	// fmt.Println("pointer of d is", *c)
	// fmt.Println("Type of Cart is ", *b)

	//----------------------------------------------------------------------------
	// Default value of a pointer

	// var p1 *int
	// fmt.Printf("p1 address: %v\n", p1) //nil

	//----------------------------------------------------------------------------
	// Dereferencing a pointer

	// b := 255
	// a := &b

	// fmt.Println("address of b is", a)
	// fmt.Println("value of b is", *a)
	// *a++
	// fmt.Println("new value of b is", b)

	//----------------------------------------------------------------------------
	// Pass value pointer to a function

	// a := 58

	// change(&a)
	// fmt.Println("value after function call", a)

	//----------------------------------------------------------------------------
	// Return value pointer from a function

	// d := hello()
	// fmt.Println("Value of d", *d)
	//or
	// fmt.Println("Value of d", *hello())

	//----------------------------------------------------------------------------
	// Go does not support pointer arithmetic

	// b := [...]int{109, 110, 111}
	// p := &b
	// p++ //Error

	//----------------------------------------------------------------------------
	// Double Pointer

	//----------------------------------- Expl_1
	// var V int = 100
	// var pt1 *int = &V
	// var pt2 **int = &pt1

	// fmt.Println("The Value of Variable V is = ", V)
	// fmt.Println("Address of variable V is = ", &V)

	// fmt.Println("The Value of pt1 is = ", pt1)
	// fmt.Println("Address of pt1 is = ", &pt1)

	// fmt.Println("The value of pt2 is = ", pt2)
	// fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)
	// fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)

	//----------------------------------- Expl_2
	// var v int = 100
	// var pt1 *int = &v
	// var pt2 **int = &pt1

	// fmt.Println("The Value of Variable v is = ", v)

	// *pt1 = 200
	// fmt.Println("Value stored in v after changing pt1 = ", v)

	// **pt2 = 300
	// fmt.Println("Value stored in v after changing pt2 = ", v)

	//----------------------------------------------------------------------------

}

//----------------------------------------------------------------------------
// Pass value pointer to a function

func change(val *int) {
	*val = 55
}

//----------------------------------------------------------------------------
// Return value pointer from a function

func hello() *int {
	i := 5
	return &i
}

//----------------------------------------------------------------------------
