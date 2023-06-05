/*---------------------------------------------------------------------------------
  Description
  Pointer with Method (Pointer receiver vs. value receiver)
---------------------------------------------------------------------------------*/

package main

import "fmt"

//----------------------------------------------------------------------------
// Pointer with Method

type myName struct {
	Name string
}

//----------------------------------------------------------------------------

func main() {

	//----------------------------------------------------------------------------
	// Description

	// Use a pointer receiver :
	//1 Your struct is heavy
	//2 You want to modify the receiver
	//3 Your struct contains a synchronization primitive

	// Use a value receiver :
	//1 When your struct is small
	//2 When you do not intend to modify the receiver
	//3 When the receiver is a map, a func, a chan, a slice, a string, or an interface value (because internally it’s already a pointer)
	//4 When your other receivers are pointers

	//----------------------------------------------------------------------------
	// Pointer with Method (Pointer receiver vs. value receiver)

	//theStructure := myName{Name: "MyName"}

	//----------------------------------- by value(no pointer)
	// theStructure.noPointer()
	// fmt.Println(theStructure.Name)

	//----------------------------------- by reference(with pointer)
	// theStructure.withPointer()
	// fmt.Println(theStructure.Name)

	//-----------------------------------
	// myName{}.noPointer_p()      //True
	// //myName{}.withPointer_p()  //Error(cannot call pointer method on myStructure{})
	// ms := &myName{}    //True
	// ms.withPointer_p() //True

	//----------------------------------------------------------------------------

}

//----------------------------------------------------------------------------
// Pointer with Method

func (ms myName) noPointer() {
	ms.Name = "xxxx"
	fmt.Println(ms.Name)
}

func (ms *myName) withPointer() {
	ms.Name = "yyyy"
	fmt.Println(ms.Name)
}

func (ms myName) noPointer_p() {
	fmt.Println("Not a pointer")
}

func (ms *myName) withPointer_p() {
	fmt.Println("a pointer")
}

//----------------------------------------------------------------------------
