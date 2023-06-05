/*---------------------------------------------------------------------------------
  Description
  Declaration Struct & short hand declaration
  Initialization
  Default value
  Field tags
  Copy And Modifying Structs (by value)
  Copy And Modifying Structs in Function (by value)
  Nested Structs
  Promoted Fields
  Comparison Structs
  Anonymous fields
  Creating Anonymous structs

  Length of Structs  *
  Delete Items  *
  Add Items  *
  Iterating  *

  Accessing & Modify fields of a struct
 ---------------------------------------------------------------------------------*/

package main

import "fmt"

//----------------------------------------------------------------------------
// Declaration Structs & short hand declaration

type Employee struct {
	firstName string
	lastName  string
	age       int
}

type Employee2 struct {
	firstName, lastName string
	age                 int
}

//----------------------------------------------------------------------------
// Default value

type Person3 struct {
	Name    *string
	Family  string
	Tell    *int
	Address int
	Man     bool
	*Sex
}

type Sex struct {
	Man   *string
	Zan   string
	Bache bool
	Sen   int
}

//----------------------------------------------------------------------------
// Nested structs

//----------------------------------- Exmp_1

type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address
}

//----------------------------------- Exmp_2

type Salary struct {
	Basic, HRA, TA float64
}

type Employee4 struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

//----------------------------------------------------------------------------
// Promoted fields

type details struct {
	name   string
	age    int
	gender string
}

type student struct {
	branch string
	year   int
	details
}

//----------------------------------------------------------------------------
// Comparison Structs

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

//----------------------------------------------------------------------------
// Anonymous fields

type Employee3 struct {
	string
	int
	bool
}

//----------------------------------------------------------------------------

func main() {

	//---------------------------------------------------------------------------------
	// Description

	//1- Structs are aggregate types.
	//2- Structs is a collection of disparate data types.
	//3-

	//----------------------------------------------------------------------------
	// Declaration Struct & short hand declaration

	// emp1 := Employee{
	// 	firstName: "Sam",
	// 	lastName:  "Anderson",
	// 	age:       25,
	// }

	// emp2 := Employee{"Thomas", "Paul", 800}
	// var emp3 = Employee2{"Thomas", "Paul", 800}

	// var emp4 = Employee2{"Thomas", "Paul"}    //Eror

	// var emp5 = Employee{
	// 	firstName: "Sam",
	// 	lastName:  "Anderson",
	// }

	//var emp6 = Employee{}  //zero valued struct
	//var emp7 Employee      //zero valued struct
	//emp8 := Employee{}     //zero valued struct

	// fmt.Println("Employee 1", emp1)
	// fmt.Println("Employee 2", emp2)
	// fmt.Println("Employee 3", emp3)
	// fmt.Println("Employee 5", emp5)

	//----------------------------------------------------------------------------
	// Initialization

	// var emp Employee       //zero valued struct
	// var emp1 = Employee{}  //zero valued struct
	// emp2 := Employee{}     //zero valued struct

	//emp.age = 20
	//emp.lastName = "asghari"
	//emp.firstName = "ali"

	//fmt.Println(emp)

	//----------------------------------------------------------------------------
	// Default value

	//var prman Person3
	//var prman = Person3{}
	//prman := Person3{}

	//fmt.Println("First Name prman:", prman)

	//----------------------------------------------------------------------------
	// Field tags

	// type Post struct {
	// 	Title  string `json:"title" myfmt:"s1"`
	// 	Author string `json:"author,omitempty" myfmt:"s2"`
	// 	Pages  int    `json:"pages,omitempty" myfmt:"n1"`
	// 	X, Y   bool   `myfmt:"b1"`
	// }

	//----------------------------------------------------------------------------
	// Copy And Modifying Structs (by value)

	// p1 := Employee{"Thomas", "Paul", 800}
	// p2 := p1    //A copy of the struct `p1` is assigned to `p2`

	// p2.age = 15
	// fmt.Println("\n After Modifying p2:")
	// fmt.Println("p1 = ", p1)
	// fmt.Println("p2 = ", p2)

	//----------------------------------------------------------------------------
	// Copy And Modifying Structs in Function (by value)

	// p1 := Employee{"Thomas", "Paul", 800}

	// fmt.Println("before passing to function ", p1)
	// Modify(p1)
	// fmt.Println("after passing to function ", p1)

	//----------------------------------------------------------------------------
	// Nested Structs

	//----------------------------------- Exmp_1
	//Struct and Struct

	// p := Person{
	// 	name: "Naveen",
	// 	age:  50,
	// 	address: Address{
	// 		city:  "Chicago",
	// 		state: "Illinois",
	// 	},
	// }

	// fmt.Println("Name:", p.name)
	// fmt.Println("City:", p.address.city)
	// fmt.Println("State:", p.address.state)

	//----------------------------------- Exmp_2
	//Struct and Slice

	// e := Employee4{
	// 	FirstName: "Mark",
	// 	LastName:  "Jones",
	// 	Email:     "mark@gmail.com",
	// 	Age:       25,
	// 	MonthlySalary: []Salary{
	// 		Salary{
	// 			Basic: 15000.00,
	// 			HRA:   5000.00,
	// 			TA:    2000.00,
	// 		},
	// 		Salary{
	// 			Basic: 16000.00,
	// 			HRA:   5000.00,
	// 			TA:    2100.00,
	// 		},
	// 	},
	// }

	// fmt.Println(e.FirstName, e.LastName)
	// fmt.Println(e.Age)
	// fmt.Println(e.Email)
	// fmt.Println(e.MonthlySalary[0])
	// fmt.Println(e.MonthlySalary[1])

	//----------------------------------------------------------------------------
	// Promoted fields

	// values := student{
	// 	branch: "CSE",
	// 	year:   2010,
	// 	details: details{
	// 		name:   "Sumit",
	// 		age:    28,
	// 		gender: "Male",
	// 	},
	// }

	// fmt.Println("Name: ", values.name)
	// fmt.Println("Age: ", values.age)
	// fmt.Println("Gender: ", values.gender)

	// fmt.Println("Year: ", values.year)
	// fmt.Println("Branch : ", values.branch)

	//----------------------------------------------------------------------------
	// Comparison Structs

	//----------------------------------- Exmp_1
	// name1 := name{
	// 	firstName: "Steve",
	// 	lastName:  "Jobs",
	// }

	// name2 := name{
	// 	firstName: "Steve",
	// 	lastName:  "Jobs",
	// }

	// if name1 == name2 {
	// 	fmt.Println("name1 and name2 are equal")
	// } else {
	// 	fmt.Println("name1 and name2 are not equal")
	// }

	// name3 := name{
	// 	firstName: "Steve",
	// 	lastName:  "Jobs",
	// }

	// name4 := name{
	// 	firstName: "Steve",
	// }

	// if name3 == name4 {
	// 	fmt.Println("name3 and name4 are equal")
	// } else {
	// 	fmt.Println("name3 and name4 are not equal")
	// }

	//----------------------------------- Exmp_2
	// image1 := image{
	// 	data: map[int]int{
	// 		0: 155,
	// 	}}

	// image2 := image{
	// 	data: map[int]int{
	// 		0: 155,
	// 	}}

	// if image1 == image2 { //Error
	// 	fmt.Println("image1 and image2 are equal")
	// }

	//----------------------------------------------------------------------------
	// Anonymous fields

	// p1 := Employee3{
	// 	string: "naveen",
	// 	int:    50,
	// 	bool:   true,
	// }

	// fmt.Println(p1.string)
	// fmt.Println(p1.int)
	// fmt.Println(p1.bool)

	//----------------------------------------------------------------------------
	// Creating anonymous structs

	// emp3 := struct {
	// 	firstName string
	// 	age       int
	// }{
	// 	firstName: "Andrea-h",
	// 	age:       31,
	// }

	// fmt.Println("Employee 3", emp3)

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// Accessing & Modify fields of a struct

	// emp6 := Employee{
	// 	firstName: "Sam",
	// 	age:       55,
	// }

	// fmt.Println("Age:", emp6.age)
	// emp6.age = 65
	// fmt.Printf("New Salary: $%d", emp6.age)

	//----------------------------------------------------------------------------

}

//---------------------------------------------------------------------------------
// Copy And Modifying Arrays in Function

func Modify(e Employee) {
	e.age = 200
	fmt.Println("in function: ", e)
}

//---------------------------------------------------------------------------------
