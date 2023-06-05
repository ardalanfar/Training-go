/*---------------------------------------------------------------------------------
  Switch Syntax
  Switch Example
  Default case (else)
  Multiple expressions in case
  Expressionless switch (expression is omitted)
  Fallthrough (next)
  Breaking switch
  Breaking the label outer for loop
---------------------------------------------------------------------------------*/

package main

func main() {

	//---------------------------------------------------------------------------------
	// Switch Syntax

	// switch statement; main expression{
	// case expression 1,.... :
	// case expression 2,.... :
	// default:
	// }

	//---------------------------------------------------------------------------------
	// Switch Example
	// the first one that equals the switch expression triggers execution of the statements of the associated case,

	//finger := 3
	//fmt.Printf("Finger %d is ", finger)

	//switch finger {
	//case 1:
	//	fmt.Println("Thumb")
	//case 2, 3:
	//	fmt.Println("Index")
	//case 3 ,4 ,5:
	//	fmt.Println("Middle")
	//}

	//---------------------------------------------------------------------------------
	// Default case (else)

	//switch finger := 6; finger {
	//case 1:
	//	fmt.Println("Thumb")
	//case 2:
	//	fmt.Println("Index")
	//case 3:
	//	fmt.Println("Middle")
	//default: //default case
	//	fmt.Println("incorrect finger number")
	//}

	//---------------------------------------------------------------------------------
	// Multiple expressions in case

	//letter := "i"
	//fmt.Printf("Letter %s is a ", letter)

	//switch letter {
	//case "a", "e", "i", "o", "u": //multiple expressions in case
	//	fmt.Println("vowel")
	//default:
	//	fmt.Println("not a vowel")
	//}

	//---------------------------------------------------------------------------------
	// Expressionless switch (expression is omitted)

	//num := 75

	//switch {  //expression is omitted
	//case num >= 0 && num <= 50:
	//	fmt.Printf("%d is greater than 0 and less than 50", num)
	//case num >= 51 && num <= 100:
	//	fmt.Printf("%d is greater than 51 and less than 100", num)
	//case num >= 101:
	//	fmt.Printf("%d is greater than 100", num)
	//}

	//---------------------------------------------------------------------------------
	// Fallthrough (next)

	// switch num := 25; {
	// case num < 50:
	// 	fmt.Printf("%d is lesser than 50\n", num)
	// 	fallthrough
	// case num > 100:
	// 	fmt.Printf("%d is greater than 100\n", num)
	// case num < 60:
	// 	fmt.Println("salam")
	// }

	//---------------------------------------------------------------------------------
	// Breaking switch

	// switch num := 6; {
	// case num < 50:
	// 	if num < 0 {
	// 		fmt.Printf("break")
	// 		break
	// 	}
	// 	fmt.Printf("%d is lesser than 50\n", num)
	// 	fallthrough
	// case num > 100:
	// 	fmt.Printf("%d is lesser than 100\n", num)
	// 	fallthrough
	// case num < 200:
	// 	fmt.Printf("%d is lesser than 200\n", num)
	// }

	//---------------------------------------------------------------------------------
	// Breaking the label outer for loop

	// randloop:
	// 	for {
	// 		switch i := rand.Intn(100); {
	// 		case i%2 == 0:
	// 			fmt.Printf("Generated even number %d", i)
	// 			break randloop
	// 		}
	// 	}

	//---------------------------------------------------------------------------------

}
