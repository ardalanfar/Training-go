/*---------------------------------------------------------------------------------
  Closures function
--------------------------------------------------------------------------------*/
package main

func main() {
	//-------------------------------------------------------------------------------
	// Closures function
	//1- A closure is a function that references variables outside of its scope.
	//2- A closure can outlive the scope in which it was created. Thus it can access variables within that scope, even after the scope is destroyed.
	//3-

	//----------------------------------- Expl_1
	// a := 5

	// func() {
	// 	fmt.Println("a =", a)
	// }()

	//----------------------------------- Expl_2
	// a := appendStr()
	// b := appendStr()

	// fmt.Println(a("World"))
	// fmt.Println(b("Everyone"))

	// fmt.Println(a("Gopher"))
	// fmt.Println(b("!"))

	//----------------------------------- Expl_3
	// counter := newCounter()

	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())
	//fmt.Println(newCounter()) //Address

	//----------------------------------- Expl_4
	// pos, neg := adder(), adder()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(pos(i),neg(-2*i))
	// }

	//---------------------------------------------------------------------------------
}

//---------------------------------------------------------------------------------
// Closures function

//----------------------------------- Expl_2
func appendStr() func(b string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

//----------------------------------- Expl_3
func newCounter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//----------------------------------- Expl_4
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//---------------------------------------------------------------------------------
