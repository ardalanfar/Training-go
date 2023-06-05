/*---------------------------------------------------------------------------------
Variables

IF statement
IF ... else if constructs should end with "else" clauses
Control flow statements "if", "for" and "switch" should not be nested too deeply

Switch statements should have "default" clauses
Switch case clauses should not have too many "lines"
Switch statements should not have too many "case" clauses
Switch statements should not be nested

Receiver functions (use method)
Functions should not be empty
Functions should not have identical implementations
Functions should not have too many parameters
Return early In Function
Clean Functions

The name of the interface better "er" end with

Error management (error handling)
Redundant pairs of parentheses should be removed
Nested blocks of code should not be left empty
Boolean checks should not be inverted
Duplicate code
Track uses of "TODO" tags
---------------------------------------------------------------------------------*/

package main

func main() {
	//----------------------------------------------------------------------------
	//Variables

	//---------------------------- bad
	// var (
	// 	a = 10
	// 	b = 20
	// )

	// var (
	// 	c = 30
	// )

	//---------------------------- Solution

	// var (
	// 	a = 10
	// 	b = 20
	// 	c = 30
	// )

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// IF statement

	//---------------------------- bad
	// f, contains := factory[string(token)]
	// if contains {
	// 	// Do something
	// }

	//---------------------------- Solution
	// if f, contains := factory[sToken]; contains {
	// 	// Do something
	// }

	//----------------------------------------------------------------------------
	// IF ... else if constructs should end with "else" clauses

	//---------------------------- bad
	// if x == 0 {
	// 	doSomething()
	// } else if x == 1 {
	// 	doSomethingElse()
	// }

	//---------------------------- Solution
	// if x == 0 {
	// 	doSomething()
	// } else if x == 1 {
	// 	doSomethingElse()
	// } else {
	// 	return errors.New("unsupported int")
	// }

	//----------------------------------------------------------------------------
	// Control flow statements "if", "for" and "switch" should not be nested too deeply

	//---------------------------- bad
	// if condition1 { // Compliant - depth = 1
	// 	/* ... */
	// 	if condition2 { // Compliant - depth = 2
	// 		/* ... */
	// 		for i := 1; i <= 10; i++ { // Compliant - depth = 3, not exceeding the limit
	// 			/* ... */
	// 			if condition4 { // Noncompliant - depth = 4
	// 				if condition5 { // Depth = 5, exceeding the limit, but issues are only reported on depth = 4
	// 					/* ... */
	// 				}
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// Switch statements should have "default" clauses

	//---------------------------- bad
	// switch simpleToken.token {
	// case tokenTitle:
	// 	msg.Title = value
	// case tokenAdep:
	// 	msg.Adep = value
	// case tokenAltnz:
	// 	msg.Alternate = value
	// 	// Other cases
	// }

	//---------------------------- Solution
	// switch simpleToken.token {
	// case tokenTitle:
	// 	msg.Title = value
	// case tokenAdep:
	// 	msg.Adep = value
	// case tokenAltnz:
	// 	msg.Alternate = value
	// // Other cases
	// default:
	// 	log.Errorf("unexpected token type %v", simpleToken.token)
	// 	return Message{}, fmt.Errorf("unexpected token type %v", simpleToken.token)
	// }

	//----------------------------------------------------------------------------
	// Switch case clauses should not have too many "lines"

	//---------------------------- bad
	// func foo(tag int) {
	// 	switch tag {
	// 	case 0:
	// 		methodCall1()
	// 		methodCall2()
	// 		methodCall3()
	// 		methodCall4()
	// 	case 1:
	// 		bar()
	// 	}
	// }

	//---------------------------- Solution
	// func foo(tag int) {
	// 	switch tag {
	// 	case 0:
	// 		executeAll()
	// 	case 1:
	// 		bar()
	// 	}
	// }

	// func executeAll() {
	// 	methodCall1()
	// 	methodCall2()
	// 	methodCall3()
	// 	methodCall4()
	// }

	//----------------------------------------------------------------------------
	// Switch statements should not have too many "case" clauses

	//1-When switch statements have large sets of case clauses, it is usually an attempt to map two sets of data. A real map structure would be more readable and maintainable, and should be used instead.

	//----------------------------------------------------------------------------
	// Switch statements should not be nested

	//---------------------------- bad
	// func foo(x,y int) {
	// 	switch x {
	// 	case 0:
	// 		switch y { // Noncompliant; nested switch
	// 		// ...
	// 		}
	// 	case 1:
	// 		// ...
	// 	default:
	// 		// ...
	// 	}
	// }

	//---------------------------- Solution
	// func foo(x,y int) {
	// 	switch x {
	// 	case 0:
	// 		bar(y)
	// 	case 1:
	// 		// ...
	// 	default:
	// 		// ...
	// 	}
	// }

	// func bar(y int) {
	// 	switch y {
	// 	// ...
	// 	}
	// }

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// Receiver functions (use method)

	//---------------------------- bad
	// func IsUpperLevel(m Message) bool {
	// 	for _, r := range m.RoutePoints {
	// 		if r.FlightLevel > upperLevel {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	//---------------------------- Solution
	// func (m *Message) IsUpperLevel() bool {
	// 	for _, r := range m.RoutePoints {
	// 		if r.FlightLevel > upperLevel {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	//----------------------------------------------------------------------------
	// Functions should not be empty

	//---------------------------- bad
	// func doNothing() { // Noncompliant
	//
	// }

	//---------------------------- Solution
	// func doNothing() {
	// 	// Do nothing because of X and Y.
	// }

	//----------------------------------------------------------------------------
	// Functions should not have identical implementations

	//---------------------------- bad
	// func fun1() (x, y int) {
	// 	a, b := 1, 2
	// 	b, a = a, b
	// 	return a, b
	// }

	// func fun2() (x, y int) {  // Noncompliant; fun1 and fun2 have identical implementations
	// 	a, b := 1, 2
	// 	b, a = a, b
	// 	return a, b
	// }

	//---------------------------- Solution
	// func fun1() (x, y int) {
	// 	a, b := 1, 2
	// 	b, a = a, b
	// 	return a, b
	// }

	// func fun2() (x, y int) {  // Compliant
	// 	return fun1()
	// }

	//----------------------------------------------------------------------------
	// Functions should not have too many parameters

	//---------------------------- bad
	// func foo(p1 int, p2 int, p3 int, p4 int, p5 int , p6 int) { // Noncompliant
	// 	// ...
	// }

	//---------------------------- Solution
	// func foo(p1 int, p2 int, p3 int, p4 int) {
	// 	// ...
	// }

	//----------------------------------------------------------------------------
	// Return early In Function

	//---------------------------- bad
	// func do(n int) bool {
	// 	if n > 12 {
	// 	    return false
	// 	} else {
	// 	    return true
	// 	}
	//   }

	//---------------------------- Solution
	// func do(n int) bool {
	// 	if n > 12 {
	// 	    return false
	// 	}
	// 	return true
	//   }

	//----------------------------------------------------------------------------
	// Clean Functions

	//1- First parameter argument function context type
	//2- End parameter argument function Error type
	//3- Avoid writing longer functions ,Smaller functions are better ,Longer functions can be hard to read
	//4- Try to keep fewer function parameters
	//5- Good function names are better than comments
	//6-

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// The name of the interface better "er" end with

	//---------------------------- Solution
	// type Adder interface {
	// }

	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	// Error management (error handling)

	//---------------------------- bad
	//preprocessed, _ := preprocess(string)

	//---------------------------- Solution
	// preprocessed, err := preprocess(bytes)
	// if err != nil {
	// 	return Message{}, err
	// }

	//----------------------------------------------------------------------------
	// Redundant pairs of parentheses should be removed

	//---------------------------- bad
	// func foo(a bool, y int) int {
	// 	x := (y / 2 + 1)   //Compliant even if the parenthesis are ignored by the compiler

	// 	if a && ((x+y > 0)) {  // Noncompliant
	// 	  //...
	// 	}

	// 	return ((x + 1))  // Noncompliant
	// }

	//---------------------------- Solution
	// func foo(a bool, y int) int {
	// 	x := (y / 2 + 1)

	// 	if a && (x+y > 0) {
	// 	  //...
	// 	}

	// 	return (x + 1)
	// }

	//----------------------------------------------------------------------------
	// Nested blocks of code should not be left empty

	//---------------------------- bad
	// func compute(a int, b int) {
	// 	sum := a + b
	// 	if  sum > 0 { } // Noncompliant; empty on purpose or missing piece of code?
	// 	fmt.Println("Result:", sum)
	// }

	//---------------------------- Solution
	// func compute(a int, b int) {
	// 	sum := a + b
	// 	if  sum > 0 {
	// 		fmt.Println("Positive result")
	// 	}
	// 	fmt.Println("Result:", sum)
	// }

	//----------------------------------------------------------------------------
	// Boolean checks should not be inverted

	//---------------------------- bad
	// if ( !(a == 2)) {...}  // Noncompliant
	// boolean b = !(i < 10);  // Noncompliant

	//---------------------------- Solution
	// if (a != 2) { ...}
	// boolean b = (i >= 10);

	//----------------------------------------------------------------------------
	// Duplicate code

	//---------------------------- bad
	// fmt.println("What would you like for breakfast?")
	// breakfast := fmt.Scan()
	// fmt.printf("One {breakfast} coming up What would you like for lunch?")
	// lunch := fmt.Scan()
	// fmt.printf("One {lunch} coming up What would you like for dinner?")
	// dinner := fmt.Scan()
	// fmt.printf("One {dinner} coming up")

	//---------------------------- Solution
	// func ask_meal(meals_of_the_day array) string {
	//      fmt.printf("What would you like to eat for {meal_of_the_day}")
	//      meal := fmt.Scan()
	//      return fmt.Sprintf("One {meal} coming up")
	// }

	// meals_of_the_day := ["breakfast", "lunch", "dinner"]
	// for meal range meals_of_the_day{
	//     ask_meal(meal)
	// }

	//----------------------------------------------------------------------------
	// Track uses of "TODO" tags

	//TODO tags are commonly used to mark places where some more code is required,

	// func foo() {
	// 	// TODO
	// }

	//----------------------------------------------------------------------------g
}
