/*---------------------------------------------------------------------------------
  Description
  Implement Methods (Value Receivers)
  Pointer Receivers VS Value Receivers (Methods)
  Value Receivers in methods VS Value Arguments in functions
  Pointer receivers in methods VS Pointer arguments in functions
  Methods with non-struct receivers
  Methods of anonymous struct fields
---------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"math"
)

func main() {

	//----------------------------------------------------------------------------
	// Description

	//1- A method is a function with a special receiver argument.
	//2- The receiver has access to non-exported (private) fields of type.
	//3- You can declare a method on non-struct types.
	//4- A function can be used as first-order objects and can be passed around while methods cannot.
	//5- There can exist different methods with the same name with a different receiver, but there cannot exist two different functions with the same name in the same package.
	//6-

	//----------------------------------------------------------------------------
	// Implement Methods (Value Receivers)

	//------------------------------------ Expl_1
	// r := Rectangle{
	// 	length: 10,
	// 	width:  5,
	// }
	// fmt.Printf("Area of rectangle %d\n", r.Area())

	// c := Circle{
	// 	radius: 12,
	// }
	// fmt.Printf("Area of circle %f", c.Area())

	//------------------------------------ Expl_2
	//myStr := MyString("OLLEH")
	//fmt.Println(myStr.reverse())

	//----------------------------------------------------------------------------
	// Pointer Receivers VS Value Receivers (Methods)
	//1- Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.
	//2- Pointers receivers can also be used in places where it's expensive to copy a data structure.
	//3- In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.
	//4-

	//------------------------------------
	//e := Employee2{
	//	name: "Mark Andrew",
	//	age:  50,
	//}

	//------------------------------------ Value Receivers (by value)
	//fmt.Printf("Employee name before change: %s", e.name)
	//e.changeName("Michael Andrew")
	//fmt.Printf("\nEmployee name after change: %s", e.name)

	//------------------------------------ Pointer Receivers (by reference)
	//fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//(&e).changeAge(51)
	//or
	//e.changeAge(51)
	//fmt.Printf("\nEmployee age after change: %d", e.age)

	//----------------------------------------------------------------------------
	// Value Receivers in methods VS Value Arguments in functions
	//1- When a function has a value argument, it will accept only a value argument.
	//2- When a method has a value receiver, it will accept both pointer and value receivers.
	//3-

	//r := rectangle{
	//	length: 10,
	//	width:  5,
	//}

	//area(r)
	//r.area()

	//p := &r  //Pointer to r

	//area(p)  //Error cannot use p (type *rectangle)
	//p.area() //Calling pointer receiver

	//----------------------------------------------------------------------------
	// Pointer receivers in methods VS Pointer arguments in functions
	//1- When a function has a pointer argument, it will accept only a pointer argument.
	//2- When a method has a pointer receiver, it will accept both pointer and value receivers.
	//3-

	//r := rectangle{
	//	length: 10,
	//	width:  5,
	//}

	//p := &r  //Pointer to r

	//perimeter(p)
	//p.perimeter()

	//perimeter(r)  //Error
	//r.perimeter() //Calling value receiver

	//----------------------------------------------------------------------------
	// Methods with non-struct receivers

	// num1 := myInt(5)
	// num2 := myInt(10)
	// sum := num1.add(num2)

	// fmt.Println("Sum is", sum)

	//----------------------------------------------------------------------------
	// Methods of anonymous struct fields

	// values := employee{
	// 	post: "HR",
	// 	eid:  4567,
	// 	details: details{
	// 		name:    "Sumit",
	// 		age:     28,
	// 		gender:  "Male",
	// 		psalary: 890,
	// 	},
	// }

	// fmt.Println("Name: ", values.name)
	// fmt.Println("Age: ", values.age)
	// fmt.Println("Post: ", values.post)
	// fmt.Println("Employee id: ", values.eid)

	// fmt.Println("Total Salary: ", values.promotmethod(30))

	//----------------------------------------------------------------------------
}

//----------------------------------------------------------------------------
// Implement Methods (Value Receivers)

//--------------------------------- Expl_1
type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//--------------------------------- Expl_2
type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//----------------------------------------------------------------------------
// Pointer Receivers VS Value Receivers (Methods)

type Employee2 struct {
	name string
	age  int
}

//Method with value receiver
func (e Employee2) changeName(newName string) {
	e.name = newName
}

//Method with pointer receiver
func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

//----------------------------------------------------------------------------
// Value Receivers in methods VS Value Arguments in functions

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

//----------------------------------------------------------------------------
// Pointer receivers in methods vs Pointer arguments in functions

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

//----------------------------------------------------------------------------
// Methods with non-struct receivers

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

//----------------------------------------------------------------------------
// Methods of anonymous struct fields

type employee struct {
	post string
	eid  int
	details
}

type details struct {
	name    string
	age     int
	gender  string
	psalary int
}

func (d details) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

//----------------------------------------------------------------------------
