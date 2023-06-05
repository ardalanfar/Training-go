/*---------------------------------------------------------------------------------
  Description
  Declaration and Initialization
  Default value
  Copy And Modifying Maps (by reference)
  Copy And Modifying Maps in Function (by reference)
  Delete Items
  Add Items
  Iterating
  Nested Map
  Comparison Maps
  Length of Map

  Retrieving value for a key from a map
  Checking if a key exists
  Map and structs
---------------------------------------------------------------------------------*/
package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {

	//-------------------------------------------------------------------------------
	// Description

	//1- map is a reference to a hash table.
	//2- It is a collection of elements that group in key-value pairs.
	//3- maps are dynamic data structures that grow as values are added.
	//4- A `map` maps keys to values, `map[KeyType]ValueType`.
	//5-

	//-------------------------------------------------------------------------------
	// Declaration and Initialization

	//var m1 map[string]int
	//m2 := map[string]int      //Error

	// var m3 map[string]int = map[string]int{
	// 	"one": 1,
	// 	"tow": 2,
	// }

	//m4 := map[string]int{
	//	"steve": 12000,
	//	"jamie": 15000,
	//}

	//-------------------------------------------------------------------------------
	// Default value

	//var m1 map[string]int                 //nil map
	//m2 := make(map[string]int)            //Empty map
	//m3 := make(map[string]float64, 100)   //Preallocate room for 100 entries
	//var m4 = make(map[string]int)         //Empty map

	//fmt.Println(m1, m2, m3, m4)

	//-------------------------------------------------------------------------------
	// Copy And Modifying Maps (by reference)

	// employeeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// 	"mike":  9000,
	// }

	// modified := employeeSalary
	// employeeSalary["mike"] = 18000

	// fmt.Println(modified)
	// fmt.Println(employeeSalary)

	//-------------------------------------------------------------------------------
	// Copy And Modifying Maps in Function (by reference)

	// employeeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// 	"mike":  9000,
	// }

	// fmt.Println("before passing to function ", employeeSalary)
	// Modify(employeeSalary)
	// fmt.Println("after passing to function ", employeeSalary)

	//-------------------------------------------------------------------------------
	// Delete Items

	//employeeSalary := map[string]int{
	//	"steve": 12000,
	//	"jamie": 15000,
	//	"mike":  9000,
	//}

	//fmt.Println("map before deletion", employeeSalary)
	//delete(employeeSalary, "steve")
	//fmt.Println("map after deletion", employeeSalary)

	//-------------------------------------------------------------------------------
	// Add Items

	//----------------------------------- Expl_1
	// employeeSalary := make(map[string]int)

	// employeeSalary["steve"] = 12000
	// employeeSalary["jamie"] = 15000

	//----------------------------------- Expl_2
	// employeeSalary2 := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// }

	// employeeSalary2["mike"] = 9000
	//----------------------------------- Expl_3
	// var m map[string]int

	// m["ali"] = 11 //Error

	//----------------------------------- Expl_4
	//var employee = make(map[string]int)

	//employee["Mark"] = 10
	//employee["Sandy"] = 20

	//-------------------------------------------------------------------------------
	// Iterating

	// employeeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// 	"mike":  9000,
	// }

	//----------------------------------- Expl_1

	//for key, value := range employeeSalary {
	//	 fmt.Println("Key:", key, "=>", "Element:", element)
	//}

	//----------------------------------- Expl_2

	//for item := range employeeSalary {
	//	fmt.Println(item, "=>", employeeSalary[item])
	//}

	//----------------------------------- Expl_3

	// for range employeeSalary {
	// 	fmt.Println(employeeSalary["mike"])
	// }

	//-------------------------------------------------------------------------------
	// Nested Map

	//----------------------------------- Expl_1
	// var x = map[string]map[string]string{}

	// x["fruits"] = map[string]string{}
	// x["colors"] = map[string]string{}
	// x["fruits"]["a"] = "apple"
	// x["fruits"]["b"] = "banana"
	// x["colors"]["r"] = "red"
	// x["colors"]["b"] = "blue"
	// fmt.Println(x)

	//----------------------------------- Expl_2
	// nested := map[int]map[string]string{
	// 	1: {
	// 		"a": "Apple",
	// 		"b": "Banana",
	// 		"c": "Coconut",
	// 	},
	// 	2: {
	// 		"a": "Tea",
	// 		"b": "Coffee",
	// 		"c": "Milk",
	// 	},
	// 	3: {
	// 		"a": "Italian Food",
	// 		"b": "Indian Food",
	// 		"c": "Chinese Food",
	// 	},
	// }

	//fmt.Println(nested)

	//----------------------------------- Expl_3
	//var nested2 = map[string]map[string]string{
	//	"a": map[string]string{
	//		"w": "x"},
	//	"b": map[string]string{
	//		"w": "x"},
	//	"c": map[string]string{
	//		"w": "x"},
	//	"d": map[string]string{},
	//}

	//fmt.Println(nested2)

	//----------------------------------- Expl_4
	// Nested Map and slice

	// var data = map[string][]string{}
	// data["a"] = append(data["a"], "x")
	// data["b"] = append(data["b"], "x")
	// data["c"] = append(data["c"], "x")

	// fmt.Println(data)

	//-------------------------------------------------------------------------------
	// Comparison Maps
	//Map can only be compared to nil

	//map1 := map[string]int{
	//	"one": 1,
	//	"two": 2,
	//}
	//map2 := map1
	//map3 := map[string]int{
	//	"one": 2,
	//	"two": 3,
	//}
	//var map4 map[string]int

	//if map4 == nil {
	//	fmt.Println("moosavi (map4 == nil)")
	//} else {
	//	fmt.Println("moosavi nist")
	//}

	//-------------------------------------------------------------------------------
	// Length of Map

	//employeeSalary := map[string]int{
	//	"steve": 12000,
	//	"jamie": 15000,
	//}

	//fmt.Println("length is", len(employeeSalary))

	//-------------------------------------------------------------------------------
	//-------------------------------------------------------------------------------
	//-------------------------------------------------------------------------------
	// Retrieving value for a key from a map

	// employeeSalary := map[string]int{
	// 	"jamie": 15000,
	// 	"mike":  9000,
	// }

	// salary, ok := employeeSalary["jamie"]
	// fmt.Println("Salary of", salary)

	//-------------------------------------------------------------------------------
	// Checking if a key exists

	//employeeSalary := map[string]int{
	//	"steve": 12000,
	//	"jamie": 15000,
	//}

	//value, ok := employeeSalary["steve"]
	//if ok == true {
	//	fmt.Println("Salary of steve is :", value, "value ok :", ok)
	//	return
	//}

	//-------------------------------------------------------------------------------
	// Map and structs

	//emp1 := employee{
	//	salary:  12000,
	//	country: "USA",
	//}

	//emp2 := employee{
	//	salary:  14000,
	//	country: "Canada",
	//}

	//emp3 := employee{
	//	salary:  13000,
	//	country: "India",
	//}

	//employeeInfo := map[string]employee{
	//	"Steve": emp1,
	//	"Jamie": emp2,
	//	"Mike":  emp3,
	//}

	//for name, info := range employeeInfo {
	//	fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	//}

	//-------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------
// Copy And Modifying Maps in Function

func Modify(mp map[string]int) {
	mp["mike"] = 18000
	fmt.Println("in function: ", mp)
}

//-------------------------------------------------------------------------------
