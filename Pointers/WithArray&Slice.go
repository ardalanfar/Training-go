/*---------------------------------------------------------------------------------
  Composite Literals Are Unaddressable but Can Be Taken Addresses
  Edit Slice and Array Pointer by fuction
  Edit Slice and Array by fuction (without pointer)
---------------------------------------------------------------------------------*/

package main

func main() {

	//----------------------------------------------------------------------------
	// Composite Literals Are Unaddressable but Can Be Taken Addresses

	// pSlice := &[]string{"break", "continue"}
	// pArray := &[...]bool{false, true, true, false}

	// fmt.Printf("%T\n", pSlice)  //*[]string
	// fmt.Printf("%T\n", pArray)  //*[4]bool

	//----------------------------------------------------------------------------
	// Edit Slice and Array Pointer by fuction

	// a := [3]int{89, 90, 91}
	// modify(&a)
	// fmt.Println("Array", a)

	//----------------------------------------------------------------------------
	// Edit Slice and Array by fuction (without pointer)

	//----------------------------------- Slice (by reference)

	// EUcountries1 := []string{"Austria", "Belgium"}
	// CountriesSlice(EUcountries1)
	// fmt.Println(EUcountries1)

	//----------------------------------- Array (by value)

	// EUcountries2 := [2]string{"Austria", "Belgium"}
	// CountriesArray(EUcountries2)
	// fmt.Println(EUcountries2)

	//----------------------------------------------------------------------------

}

//----------------------------------------------------------------------------
// Edit Slice and Array Pointer by fuction

func modify(arr *[3]int) {
	arr[0] = 90
	//or
	//(*arr)[0] = 90
}

//----------------------------------------------------------------------------
// Edit Slice and Array by fuction (without pointer)

func CountriesSlice(countries []string) {
	//countries = append(countries, "iran")
	countries[0] = "iran"
}

func CountriesArray(countries [2]string) {
	//countries = append(countries, "iran")  //error
	countries[0] = "iran"
}

//----------------------------------------------------------------------------
