/*---------------------------------------------------------------------------------
  Discretion
  Example_1
  Example_2
  Example_3
--------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//-------------------------------------------------------------------------------
	// Discretion

	//1- Use Embedded Interface: Call  method interface in other  method interface
	//2-

	//-------------------------------------------------------------------------------
	// Example_1
	//Although go does not offer inheritance, it is possible to create a new interfaces by embedding other interfaces.

	// emp := Employee{
	// 	firstName:   "Naveen",
	// 	basicPay:    5000,
	// 	pf:          200,
	// 	totalLeaves: 30,
	// 	leavesTaken: 5,
	// }
	// var employeOperate EmployeeOperations = emp

	// employeOperate.DisplaySalary()
	// fmt.Println("\nLeaves left =", employeOperate.CalculateLeavesLeft())

	//-------------------------------------------------------------------------------
	// Example_2

	// shirtItem := NewItme("NICKI")
	// observerFirst := Customer{email: "FIRST@gmail.com"}

	// shirtItem.register(observerFirst)
	//-------------------------------------------------------------------------------
	// Example_3

	// strSlice := []string{"ali", "mamad", "mohsen"}

	// Users := NewGetUsers(strSlice)
	// iterate := Users.GetGiv()
	// fmt.Println(iterate.GiveName())

	//-------------------------------------------------------------------------------
}

//-------------------------------------------------------------------------------
// Example_1

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s has salary $%d", e.firstName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

//-------------------------------------------------------------------------------
// Example_2

//-----------------------------------

type Observer interface {
	update(string)
}

type Customer struct {
	email string
}

func (c Customer) update(itemName string) {
	fmt.Printf("Sending to: %s for item: %s\n", c.email, itemName)
}

//-----------------------------------

type Subject interface {
	register(o Observer)
}

type Item struct {
	name string
}

func NewItme(txt string) Subject {
	return Item{
		name: txt,
	}
}

func (i Item) register(o Observer) {
	o.update(i.name)
}

//-------------------------------------------------------------------------------
// Example_3

//-----------------------------------

type Get interface {
	GetGiv() Give
}

type UserCollection struct {
	users []string
}

func (u UserCollection) GetGiv() Give {
	return UserIterator{
		users: u.users,
	}
}

func NewGetUsers(input []string) Get {
	return UserCollection{
		users: input,
	}
}

//-----------------------------------

type Give interface {
	GiveName() string
}

type UserIterator struct {
	users []string
}

func (u UserIterator) GiveName() string {
	return u.users[0]
}

//-------------------------------------------------------------------------------
