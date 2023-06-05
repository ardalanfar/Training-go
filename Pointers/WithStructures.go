/*---------------------------------------------------------------------------------
  Initialization
  Edit Struct Pointer By Fuction
  Struct Field Addressability
  Composite Literals Are Unaddressable But Can Take Addresses
  In Field Selectors, Dereferences of Receivers Can Be Implicit
---------------------------------------------------------------------------------*/

package main

import "fmt"

//----------------------------------------------------------------------------
// Edit Struct Pointer By Fuction

type myName struct {
	Name string
}

//----------------------------------------------------------------------------

func main() {

	//----------------------------------------------------------------------------
	// Initialization

	// emp8 := &myName{
	// 	Name: "Sam",
	// }

	// fmt.Println("First Name:", (*emp8).Name)
	// fmt.Println("First Name:", emp8.Name)
	// fmt.Println("emp8:", emp8)
	// fmt.Println("emp8:", *emp8)
	// fmt.Println("emp8.name:", *&emp8.Name)

	//----------------------------------------------------------------------------
	// Edit Struct Pointer By Fuction

	// structure := myName{Name: "MyName ali"}
	// structureFunction(&structure)
	// fmt.Println(structure)

	//----------------------------------------------------------------------------
	// Struct Field Addressability

	// type Book struct {
	// 	Pages int
	// }
	// var book = Book{} // book is addressable
	// p := &book.Pages
	// *p = 123

	// fmt.Println(book)       //{123}
	// fmt.Println(book.Pages) //123

	//----------------------------------------------------------------------------
	// Composite Literals Are Unaddressable But Can Take Addresses

	// type Book struct {
	// 	Pages int
	// }
	// // Book{100} is unaddressable but can taken address.
	// p := &Book{100} // <=> tmp := Book{100}; p := &tmp
	// p.Pages = 200

	// fmt.Println(p)
	// fmt.Println(p.Pages)

	//----------------------------------------------------------------------------
	// In Field Selectors, Dereferences of Receivers Can Be Implicit

	// type Book struct {
	// 	pages int
	// }
	// book1 := &Book{100} // book1 is a struct pointer
	// book2 := new(Book)  // book2 is another struct pointer
	// // Use struct pointers as structs.
	// book2.pages = book1.pages
	// fmt.Println(book2.pages)

	// // This last line is equivalent to the above line.
	// // In other words, if the receiver is a pointer,
	// // it will be implicitly dereferenced.
	// (*book2).pages = (*book1).pages
	// fmt.Println(book2.pages)

	//----------------------------------------------------------------------------

}

//----------------------------------------------------------------------------
// Edit Struct Pointer by fuction

func structureFunction(e *myName) {
	fmt.Println(e, *e, e.Name, (e).Name, &e.Name, *&e.Name)
	e.Name = "mamad"
}

//----------------------------------------------------------------------------
