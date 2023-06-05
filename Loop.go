/*---------------------------------------------------------------------------------
   For syntax
   For (break, continue, while, Forever)
   For Range
   Multiple initialisation and condition and increment (for)
   Nested loops (for)
   For and label (Goto)
   Label(Goto) loop
   Break and continue statements with label(Goto)
---------------------------------------------------------------------------------*/

package main

func main() {

	//---------------------------------------------------------------------------------
	// For syntax

	//for initialisation; condition; increment {
	//}

	//---------------------------------------------------------------------------------
	// For (break, continue, while, Forever)

	//-----------------------------------
	// with initialize

	//for i:=0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	//-----------------------------------
	// without initialize

	// var i = 0
	// var j = 10

	// for i < j {
	// 	fmt.Println(i)
	// 	i++
	// }

	//----------------------------------- (break)
	//for i:=0; i < 5; i++ {
	//	if i == 3 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	//----------------------------------- (continue)
	//for i := 1; i <= 10; i++ {
	//	if i % 2 == 0 {
	//		continue
	//	}
	//	fmt.Printf("%d ", i)
	//}

	//----------------------------------- (while)
	// i := 0
	// for i <= 10 {   //initialisation and increment are omitted (while)
	// 	fmt.Printf("%d ", i)
	// 	i += 2
	// }

	//----------------------------------- (Forever)
	// for {
	// }

	//----------------
	//for true {
	//	fmt.Printf("This loop will run forever.\n");
	//}

	//---------------------------------------------------------------------------------
	// For Range
	// range on (arrays , slices , map , string)

	//-----------------------------------
	// arrays and slices

	//nums := []int{1, 2, 3, 4}

	//for index, val := range nums {
	//	if val == 3 {
	//		fmt.Println("index array:", index)
	//	}
	//}

	//-----------------------------------
	// map

	// kvs := map[string]string{"a": "apple", "b": "banana"}

	// for key, val := range kvs {
	// 	fmt.Printf("%s -> %s\n", key, val)
	// }

	//----------------
	// `range` can also iterate over just the keys of a map.

	// kvs := map[string]string{"a": "apple", "b": "banana"}

	// for key := range kvs {
	// 	fmt.Println("key:", key)
	// }

	//-----------------------------------
	// strings

	//for index, char := range "golang" {
	//	fmt.Println(index, char)
	//}

	//-----------------------------------
	// blank identifier

	//for _, val := range nums {
	//	if val == 3 {
	//		fmt.Println("Value:", val)
	//	}
	//}

	//-----------------------------------
	// without initialize

	// Array := [5]int{10, 20, 30, 40, 50}

	// j := 0
	// for range Array {
	// 	fmt.Println(Array[j])
	// 	j++
	// }

	//---------------------------------------------------------------------------------
	// Multiple initialisation and condition and increment (for)

	//----------------------------------- Expl_1
	//for i := 10; i <= 10 && i <= 19; i ++{
	//	fmt.Printf(i)
	//}

	//----------------------------------- Expl_2
	//for i, j := 10, 1; j <= 10 && i <= 19; j, i = j+1, i+1 {
	//	fmt.Printf(i,j)
	//}

	//---------------------------------------------------------------------------------
	// Nested loops (for)

	//----------------------------------- Expl_1
	//num := 5
	//for i := 0; i < num; i++ {
	//	for j := 0; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//----------------------------------- Expl_2
	//for i := 0; i < 3; i++ {
	//	for j := 1; j < 4; j++ {
	//		fmt.Printf("i = %d , j = %d\n", i, j)
	//		if i == j {
	//			break
	//		}
	//	}
	//}

	//----------------------------------- Expl_3
	//names := [2]string{"ali", "mamad"}
	//fruits := [3]string{"apple", "orange", "banana"}

	//for i:=0; i < len(names); i++ {
	//	for j:=0; j < len(fruits); j++ {
	//		fmt.Println(names[i],fruits[j])
	//	}
	//}

	//--------------------------------------------------------------------------------
	// For and label (Goto)

	// outer:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 4; j++ {
	// 			fmt.Printf("i = %d , j = %d\n", i, j)
	// 			if i == j {
	// 				break outer
	// 			}
	// 		}
	// 	}

	//---------------------------------------------------------------------------------
	// Label(Goto) loop

	// 	var i int
	// loop:
	// 	if i < 3 {
	// 		fmt.Println("looping")
	// 		i++
	// 		goto loop
	// 	}
	// 	fmt.Println("done")

	//---------------------------------------------------------------------------------
	// Break and continue statements with label(Goto)

	// for i := 90; i < 100; i++ {
	// 	n := FindSmallestPrimeLargerThan(i)
	// 	fmt.Print("The smallest prime number larger than ")
	// 	fmt.Println(i, "is", n)
	// }

	//---------------------------------------------------------------------------------

}

//---------------------------------------------------------------------------------
// Break and continue statements with label(Goto)

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}

//---------------------------------------------------------------------------------
