/*---------------------------------------------------------------------------------
  Description
  Declaration Array and Slice
  Initialization Array and Slice
  Array literal with indices
  Default value Array and Slice
  Copy And Modifying Arrays And Slice (by value and reference)
  Modifying Arrays(by value) And Slice(by reference) in Function
  Delete Items Arrays And Slice
  Add item Array and Slice
  Iterating Arrays And Slice
  Nested Arrays And Slice (Multi Dimensional)
  Comparison Arrays And Slice
  Length and Capacity of Array And Slice
  Convert Slice to Array Pointer
  Convert Slice to Array

  Filter Array and Slice Elements
  Find in Array and Slice
  Reverse Array and Slice
  Memory Optimisation

  Defference between array and slice
---------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//---------------------------------------------------------------------------------
	// Description

	//----------------------------------- Array

	//1- Array are aggregate types.
	//2- The type `[n]T` is an array of `n` values of type `T`, fixed size.
	//3- Arrays passed by value.
	//4- if you pass an array to a function, it will receive a copy of the array.not a pointer to it.
	//5- The size of an array is part of its type.
	//6-

	//----------------------------------- Slice

	//1- Slices are reference types.
	//2- The type `[]T` is a slice with elements of type `T`, dynamically-sized.
	//3- Slices passed by reference.
	//4- What append does is `append` the elements to the end of the slice and return the result.
	//5- slice are dynamic data structures that grow as values are added.
	//5-

	//---------------------------------------------------------------------------------
	// Declaration Array and Slice

	//----------------------------------- Slice

	//Array := [5]int{76, 77, 78, 79, 80, 50}

	//var Slice1 []int = Array[1:4]
	//Slice2 := Array[:3]
	//Slice3 := Array[3:]

	//var Slice4 []string
	//func make([]T, len, cap) []T

	//Slice5 := []int{}
	//Slice6 := []int{8, 7, 6, 9, 9}
	//Slice7 := []int     //ERROR

	//----------------------------------- Array

	//var array [4]int
	//Array1 := [6]int{12, 4, 5, 6, 8, 2}
	//Array2 := [4]int{1, 2}
	//Array3 := [...]int{12, 78, 50, 44, 12, 6}
	//Array4 := [5]int     //ERROR
	//Array5 := [4]int{12, 4, 5, 6, 8, 9, 0}    //ERROR

	//---------------------------------------------------------------------------------
	// Initialization Array and Slice

	//----------------------------------- Array

	//var array [5]int = [5]int{11, 22, 33, 44, 55}
	//var array1 = [5]int{11, 22, 33, 44, 55}
	//var array2 = [...]int{11, 22, 33, 44, 55}  //the array length is determined by the number of initializers.

	//var array3 [5]int = [3]int{1, 2, 3}    //Error
	//var array4 [4]int = []int{1, 2, 3}     //Error
	//var array [5]int{1, 2, 3, 4, 5}        //Error

	//----------------------------------- Slice

	//Slice := []string{"a", "b", "c", "d"}

	//---------------------------------------------------------------------------------
	// Array literal with indices

	// array := [...]int{0: 11, 9: 99}
	// fmt.Println(array)

	//---------------------------------------------------------------------------------
	// Default value Array and Slice

	// var slice []string
	// var array [3]int

	// fmt.Printf("Slice = %d\n", slice)
	// fmt.Printf("Array = %d\n", array)

	//---------------------------------------------------------------------------------
	// Copy And Modifying Arrays and Slice (by value and reference)

	//array := [...]int{1, 2, 3, 4, 5}

	//----------------------------------- Slice

	//S1 := array[:]     //Creates slice (by reference)
	//S2 := array[0:5]   //Creates slice (by reference)

	//----------------------------------- Array

	//A1 := array    //Copy Array (by value)
	//A2 := &array   //Copy Pointer Address Array (by reference)

	//-----------------------------------

	//array[0] = 50
	//fmt.Println("S1 : ", S1)

	//---------------------------------------------------------------------------------
	// Modifying Arrays And Slice in Function (by value and reference)

	//Array := [...]int{1, 2, 3, 4, 5}
	//or
	//Slice := []int{1, 2, 3, 4, 5}

	//----------------------------------- Slice

	// fmt.Println("before passing to function ", Slice)
	// modifyBySlice(Slice)   //by reference
	// fmt.Println("after passing to function ", Slice)

	//----------------------------------- Array

	// fmt.Println("before passing to function ", Array)
	// modifyByArray(Array)   //by value
	// fmt.Println("after passing to function ", Array)

	//---------------------------------------------------------------------------------
	// Delete Items Arrays And Slice

	//----------------------------------- Slice

	//Slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	//fmt.Println("Slice: ", RemoveIndex(Slice, 5))

	//----------------------------------- Array

	//Array := [7]int{10, 20, 30, 40, 50, 60, 70}
	//fmt.Println("Array: ", RemoveIndex(Array, 5))  //Error

	//---------------------------------------------------------------------------------
	// Add item Array and Slice

	//----------------------------------- Add slice
	//var slice []string
	//or
	//slice := []string{}

	//slice = append(slice, "string1")
	//slice = append(slice, "string3", "string4")

	//slice[0] = "salam"

	//fmt.Println(slice[0])

	//------------------------------------ Add Array
	//var arry [2]int
	//arry[0] = 100
	//arry[1] = 200

	//------------------------------------- Add slice
	//slice := []string{"Ferrari", "Honda", "Ford", "kamil"}

	//fmt.Println("slice:", slice, "old length", len(slice), " capacity", cap(slice))   //capacity of slice is 4
	//slice2 = append(slice, "Toyota")
	//fmt.Println("slice2:", slice2, "new length", len(slice2), " capacity", cap(slice2))   //capacity of cars is doubled to 8

	//------------------------------------- Add slice
	// It is also possible to append one slice to another slice using the ... operator.

	//slice1 := []string{"potatoes", "tomatoes", "brinjal"}
	//slice2 := []string{"oranges", "apples"}

	//slice3 := append(slice1, slice2...)
	//fmt.Println("food:", slice3)

	//------------------------------------- Add slice
	//slice := append([]byte("Hello "), "world!"...)

	//fmt.Println("result:", slice)

	//---------------------------------------------------------------------------------
	// Iterating Arrays And Slice

	//----------------------------------- (for)
	//array := [...]float64{67.7, 89.8, 21, 78, 10.7}

	//for i := 0; i < len(array); i++ {    //looping from 0 to the length of the array
	//	fmt.Printf("%d th element of array is %.2f\n", i, a[i])
	//}

	//----------------------------------- (for range)
	// array := [...]float64{67.7, 89.8, 21, 78}

	// sum := 0.0
	// for i, v := range array {    //range returns both the index and value
	// 	fmt.Printf("%d the element of array is %.2f\n", i, v)
	// 	sum += v
	// }
	// fmt.Println("\n sum of all elements of array", sum)

	//----------------------------------- (for range)
	//array := [...]float64{67.7, 89.8, 21, 78}

	//for _, v := range array {     //ignores index
	//	fmt.Println("item : ", v)
	//}

	//----------------------------------- (for range)
	// Array := [5]int{10, 20, 30, 40, 50}

	// j := 0
	// for range Array {
	// 	fmt.Println(Array[j])
	// 	j++
	// }

	//---------------------------------------------------------------------------------
	// Nested Arrays And Slice (Multi Dimensional)

	//Slice := [][]string{{"C", "C++"}, {"JavaScript"}, {"Go", "Rust"}}
	// //or
	// Slice2 := [][]int{
	// 	[]int{1, 2, 3, 4, 5},
	// 	[]int{6, 7, 8, 9, 10},
	// 	[]int{11, 12, 13, 14, 15},
	// }

	//for _, v1 := range Slice {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%s ", v2)
	//	}
	//	fmt.Printf("\n")
	//}

	//---------------------------------------------------------------------------------
	// Comparison Arrays And Slice

	//----------------------------------- Expl_1
	//a := [3]int{1, 2, 3}
	//b := [3]int{1, 2, 3}
	//c := [3]int{4, 2, 3}
	//d := a
	//var e [3]int

	//if a == b {
	//	fmt.Println("a == b")
	//}
	//if a == c {
	//	fmt.Println("a == c")
	//}
	//if a == d {
	//	fmt.Println("a == d")
	//}
	//if a == e {
	//	fmt.Println("a == e")
	//}

	//----------------------------------- Expl_2
	//a := [3]int{5, 78, 8}
	//var b [3]int
	//var c [5]int

	//b = a   //ok
	//c = a   //Error (not possible since [3]int and [5]int are distinct types)
	//fmt.Println(b)

	//---------------------------------------------------------------------------------
	// Length and Capacity of Array And Slice

	//----------------------------------- Expl_1
	// var Slice []string
	// a = append(Slice, "string1")
	// a = append(Slice, "string2")

	// fmt.Println("length of Slice is", len(Slice))

	//------------------------------------ Expl_2
	// array := [...]string{"grape", "mango", "melon", "apple", "chikoo", "loko"}
	// fmt.Printf("length %d capacity %d\n", len(array), cap(array))

	// Slice := array[1:3]
	// fmt.Printf("length of slice %d capacity %d\n", len(Slice), cap(Slice))
	// fmt.Println("Slice:=>  ", Slice)

	// Slice2 := Slice[:cap(Slice)]
	// fmt.Println("After re-slicing length", len(Slice2), "and capacity", cap(Slice2))
	// fmt.Println("Slice2:=>  ", Slice2)

	//---------------------------------------------------------------------------------
	// Convert Slice to Array Pointer

	// type S []int
	// type A [2]int
	// type P *A

	// var x []int
	// var y = make([]int, 0)
	// var x0 = (*[0]int)(x)   //okay, x0 == nil
	// var y0 = (*[0]int)(y)   //okay, y0 != nil
	// _, _ = x0, y0

	// var z = make([]int, 3, 5)
	// var _ = (*[3]int)(z)   //okay
	// var _ = (*[2]int)(z)   //okay
	// var _ = (*A)(z)        //okay
	// var _ = P(z)           //okay

	// var w = S(z)
	// var _ = (*[3]int)(w)   //okay
	// var _ = (*[2]int)(w)   //okay
	// var _ = (*A)(w)        //okay
	// var _ = P(w)           //okay

	//var _ = (*[4]int)(z)    //will panic

	//---------------------------------------------------------------------------------
	// Convert Slice to Array

	// var s = []int{0, 1, 2, 3}
	// var a = []int(s[1:])
	// s[2] = 9

	// fmt.Println(s)    //[0 1 9 3]
	// fmt.Println(a)    //[1 2 3]

	//_ = [3]int(s[:2])  //panic

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	// Filter Array and Slice Elements

	// array := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	// fmt.Printf(":2 %v\n", array[:2])
	// fmt.Printf("1:3 %v\n", array[1:3])
	// fmt.Printf("2: %v\n", array[2:])

	// fmt.Printf("Last element: %v\n", array[len(array)-1])
	// fmt.Println(array[0:len(array)])
	// fmt.Printf("Last two elements: %v\n", array[len(array)-2:len(array)])

	//---------------------------------------------------------------------------------
	// Find in Array and Slice

	//item := "Iran"
	//array := [...]string{"India", "Canada", "Japan", "Iran", "Germany", "Italy"}

	//for i, j := range array {
	//	if j == item {
	//		fmt.Println("find index :", i)
	//	}
	//}

	//---------------------------------------------------------------------------------
	// Reverse Array and Slice

	//array := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	//for i := len(array) - 1; i >= 0; i-- {
	//	fmt.Println(array[i])
	//}

	//---------------------------------------------------------------------------------
	// Memory Optimisation

	//This way we can use the new slice and the original array can be garbage collected.

	//countriesNeeded := countries()
	//fmt.Println(countriesNeeded)

	//---------------------------------------------------------------------------------
	// Defference between array and slice

	//------------------------------------ Array

	//---------- Declaration and Initialization
	//var b [5]int
	//b := [2]string{"Penn", "Teller"}
	//b := [...]string{"Penn", "Teller"}  //defaults to the number of elements

	//---------- Append Array(Error)
	// arraySpace := [...]int{1, 2, 3, 4, 5}  //Error
	// arraySpace := [5]int{1, 2, 3, 4, 5}    //Error
	// arraySpace = append(arraySpace, "7")

	//---------- by value and function

	//------------------------------------ Slice

	//----------- Declaration and Initialization
	//var a []string
	//letters := []string{"a", "b", "c", "d"}

	//A slice can be created with the built-in function called make, which has the signature
	//func make([]T, len, cap) []T
	// dynamicSlice := make([]int, 2, 5) // make(datatype, length, capacity)
	// fmt.Println(dynamicSlice )        // [0 0]
	// fmt.Println(len(dynamicSlice ))   // 2
	// fmt.Println(cap(dynamicSlice ))   // 5

	//This is also the syntax to create a slice given an array:
	//array := [3]string{"hello", "my", "word"}
	//s := array[:]   //a slice referencing the storage of array
	//s := array[1:3]
	//s := array[2:]

	//----------- Append() to slices
	// var slice1 = []int{1, 2, 3, 4}
	// slice2 := []int{6,7,8}

	// slice1 = append(slice1, 5, 6, 7)
	// slice1 = append(slice1, slice2 ...)

	// sliceSpace := []string{}
	// sliceSpace = append(sliceSpace, "personA")

	//---------- by reference and function

	//---------------------------------------------------------------------------------
}

//---------------------------------------------------------------------------------
// Modifying Arrays And Slice in Function (by value and reference)

func modifyBySlice(num []int) {
	num[0] = 55
	fmt.Println("in function: ", num)
}

func modifyByArray(num [5]int) {
	num[0] = 55
	fmt.Println("in function: ", num)
}

//---------------------------------------------------------------------------------
// Delete item

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

//---------------------------------------------------------------------------------
// Memory Optimisation

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

//---------------------------------------------------------------------------------
