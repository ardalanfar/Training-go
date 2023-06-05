/*---------------------------------------------------------------------------------
  IF statement syntax
  IF else statement
  IF ... else if ... else statement
  IF with assignment
  IF with label(Goto)
---------------------------------------------------------------------------------*/

package main

func main() {

	//----------------------------------------------------------------------------
	// IF statement syntax

	// num := 10
	// if num % 2 == 0 {
	// 	fmt.Println("The number", num, "is even")
	// }

	//----------------------------------------------------------------------------
	// IF else statement

	//num := 11
	//if num % 2 == 0 {
	//	fmt.Println("the number", num, "is even")
	//} else {
	//	fmt.Println("the number", num, "is odd")
	//}

	//----------------------------------------------------------------------------
	// IF ... else if ... else statement

	// num := 99
	// if num <= 50 {
	// 	fmt.Println(num, "is less than or equal to 50")
	// } else if num >= 51 && num <= 100 {
	// 	fmt.Println(num, "is between 51 and 100")
	// } else {
	// 	fmt.Println(num, "is greater than 100")
	// }

	//----------------------------------------------------------------------------
	// IF with assignment

	//if assignment-statement; condition {
	//}

	//----------------------------------- Expl_1
	//if num := 10; num % 2 == 0 {
	//	fmt.Println(num,"is even")
	//}  else {
	//	fmt.Println(num,"is odd")
	//}

	//----------------------------------- Expl_2
	//if x := f(); x <= y {
	//	return x
	//}

	//----------------------------------------------------------------------------
	// IF with label(Goto)

	// 	num := 2
	// 	if num < 3 {
	// 		goto nxt
	// 	}
	// 	fmt.Println("bigger")
	// nxt:
	// 	fmt.Println("smaller")

	//----------------------------------------------------------------------------
}
