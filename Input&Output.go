/*---------------------------------------------------------------------------------
  fmt.Print()
  fmt.Printf()

  input(fmt.Scan())
  input(fmt.Scanf())

  fmt.Sprint()
  fmt.Sprintf()
---------------------------------------------------------------------------------*/

package main

var name string
var age int

func main() {

	//----------------------------------------------------------------------------
	// fmt.Print()
	// Print(a ...any) (n int, err error)

	//var a, b string = "Hello", "World"
	//fmt.Print(a)
	//fmt.Print(b, "\n")

	//----------------------------------------------------------------------------
	// fmt.Printf()

	//fmt.Printf("%s", "salam")

	// %s – string.
	// %d – formats base 10 numbers.
	// %b – Format base 2 numbers (binary values).
	// %o – Formats base 8 numbers.
	//

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// input(fmt.Scan())
	//The Scan() function receives the user input in raw format as space-separated values and stores them in the variables. Newlines count as spaces.

	//var a int
	//fmt.Print("Type a number: ")
	//fmt.Scan(&a)
	//fmt.Println("Your number is:", a)

	//----------------------------------------------------------------------------
	// input(fmt.Scanf())
	//The Scanf() function receives the inputs and stores them based on the determined formats for its arguments.

	//----------------------------------- Expl_1
	//fmt.Printf("Write your name : ")
	//fmt.Scanf("%s", &name)
	//fmt.Printf("Hello my name is %s", name)

	//----------------------------------- Expl_2
	//fmt.Print("Enter your name & age: ")
	//fmt.Scanf("%s %d", &name, &age)
	//fmt.Printf("%s is %d years old\n", name, age)

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// fmt.Sprint()
	//func Sprint(a ...any) string

	// var name, age = "Kim", 22
	// s := fmt.Sprint(name, " is ", age, " years old.\n")

	// fmt.Println(s)
	//----------------------------------------------------------------------------
	// fmt.Sprintf()

	// const name, age = "Kim", 22
	// s := fmt.Sprintf("%s is %d years old.\n", name, age)

	// fmt.Println(s)
	//----------------------------------------------------------------------------

}
