/*---------------------------------------------------------------------------------
  Implementing interfaces using (pointer receivers vs value receivers)
--------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//-------------------------------------------------------------------------------
	// Implementing interfaces using (pointer receivers vs value receivers)

	//------------------------------------- value receivers
	// var d1 Describer
	// p1 := Person{"Sam", 25}

	// d1 = p1
	// d1.Describe()

	//------------------------------------- pointer receivers
	// var d2 Describer
	// p2 := Person{"James", 32}

	// d2 = &p2
	// d2.Describe()

	//------------------------------------- (pointer receivers & value receivers)
	// var d3 Describer

	// adress := Address{"Washington", "USA"}
	// person := Person{"ahmad", 20}

	//----------------- (pointer receivers)
	// //d3 = adress  //Eror(cannot use a value receivers)
	// d3 = &adress   //True
	// d3.Describe()

	//----------------- (value receivers)
	// d3 = person   //True
	// d3 = &person  //True
	// d3.Describe()

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Implementing interfaces using (pointer receivers vs value receivers)

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s \n", a.state, a.country)
}

//-------------------------------------------------------------------------------
