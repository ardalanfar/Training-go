/*---------------------------------------------------------------------------------
  Discretion
  Implementing an interface (Example_1)
  Implementing an interface (Example_2)
  Interface with function
--------------------------------------------------------------------------------*/

package main

import (
	"fmt"
)

func main() {

	//-------------------------------------------------------------------------------
	// Discretion

	//1- An interface type is defined as a set of method signatures.
	//2- A value of interface type can hold any value that implements those methods.
	//3- To implement an interface in Go, we just need to implement all the methods in the interface.
	//4-

	//-------------------------------------------------------------------------------
	// Implementing an interface (Example_1)

	//----------------------------------- Implement_1
	//WithOut New Function

	// Person := Person{
	// 	name: "Naveen",
	// 	age:  50,
	// }

	// fmt.Println(Person.NameWorker("asghari")) //False
	// fmt.Println(Person.AgeWorker())           //False

	// var abs Worker = Person
	// fmt.Println(abs.NameWorker("asghari")) //True
	//fmt.Println(abs.AgeWorker()) //Error

	//----------------------------------- Implement_2
	//With New Function

	// person := NewPerson()
	// fmt.Println(person.NameWorker("asghari"))
	//fmt.Println(person.AgeWorker()) //Error

	//-------------------------------------------------------------------------------
	// Implementing an interface (Example_2)

	// pemp1 := Permanent{
	// 	empId:    1,
	// 	basicpay: 5000,
	// 	pf:       20,
	// }
	// cemp1 := Contract{
	// 	empId:    3,
	// 	basicpay: 3000,
	// }

	// employees := []SalaryInterface{pemp1, cemp1}
	// totalExpense(employees)

	//-------------------------------------------------------------------------------
	// Interface with function

	// var in Show

	// in.Showuser()  //Error
	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Implementing an interface (Example_1)

type Worker interface {
	NameWorker(family string) string
}

type Person struct {
	name string
	age  int
}

func NewPerson() Worker {
	return Person{
		name: "ahamd",
		age:  20,
	}
}

func (p Person) NameWorker(family string) string {
	return p.name
}

func (p Person) AgeWorker() int {
	return p.age
}

//-------------------------------------------------------------------------------
// Implementing an interface (Example_2)

type SalaryInterface interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryInterface) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

//-------------------------------------------------------------------------------
// Interface with function

type Show interface {
	Showuser()
}

func Showuser() {
	fmt.Println("ali is user system")
}

//-------------------------------------------------------------------------------
