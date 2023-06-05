/*---------------------------------------------------------------------------------
  TypeOf() Vs ValueOf()
  Reflect.Kind()
  NumField() Vs Field()
  Reflect.Elem()
  Int() Vs String()
  Reflect.Copy()
  Reflect.DeepEqual()
  Reflect.Swapper()
  Reflect.Indirect()
  Reflect.Interface()
  Reflect.MakeSlice()
  Reflect.MakeMap()
  Reflect.MakeChan()
  Reflect.MakeFunc()
---------------------------------------------------------------------------------*/
//Link:
//https://go.dev/blog/laws-of-reflection
//---------------------------------------------------------------------------------

package main

type order struct {
	ordId      int
	customerId int
}

//---------------------------------------------------------------------------------
// Reflect.DeepEqual()

type mobile struct {
	price float64
	color string
}

//---------------------------------------------------------------------------------

func main() {

	//---------------------------------------------------------------------------------
	// TypeOf() Vs ValueOf()

	//str := "hello"
	// ord := order{
	// 	ordId:      456,
	// 	customerId: 56,
	// }

	// fmt.Println("Type: ", reflect.TypeOf(str))
	// fmt.Println("Value: ", reflect.ValueOf(str))

	//---------------------------------------------------------------------------------
	// Reflect.Kind()

	//str := "hello"
	// ord := order{
	// 	ordId:      456,
	// 	customerId: 56,
	// }

	//k := reflect.TypeOf(ord).Kind()
	//or
	// typ := reflect.TypeOf(ord)
	// k := typ.Kind()

	//fmt.Println("Kind: ", k)

	//---------------------------------------------------------------------------------
	// NumField() Vs Field()

	// ord := order{
	// 	ordId:      456,
	// 	customerId: 56,
	// }

	// if reflect.ValueOf(ord).Kind() == reflect.Struct {
	// 	val := reflect.ValueOf(ord)
	// 	fmt.Println("Number of fields", val.NumField())
	// 	for i := 0; i < val.NumField(); i++ {
	// 		fmt.Printf("Field:%d type:%T value:%v\n", i, val.Field(i), val.Field(i))
	// 	}
	// }

	//---------------------------------------------------------------------------------
	// Reflect.Elem()

	//sNumber := []int{1, 2, 3, 4}
	// mInt := map[int]string{1: "A", 2: "B", 3: "C"}

	// fmt.Println("Elem: ", reflect.TypeOf(mInt).Elem())

	//---------------------------------------------------------------------------------
	// Int() Vs String()

	// a := 56
	// x := reflect.ValueOf(a).Int()
	// fmt.Printf("type:%T value:%v\n", x, x)

	// b := "Naveen"
	// y := reflect.ValueOf(b).String()
	// fmt.Printf("type:%T value:%v\n", y, y)

	//---------------------------------------------------------------------------------
	// Reflect.Copy()

	// destination := reflect.ValueOf([]string{"A", "B", "C"})
	// source := reflect.ValueOf([]string{"D", "E", "F"})

	// // Copy() function is used and it returns the number of elements copied
	// counter := reflect.Copy(destination, source)
	// fmt.Println(counter)

	// fmt.Println(source)
	// fmt.Println(destination)

	//---------------------------------------------------------------------------------
	// Reflect.DeepEqual()

	//// DeepEqual is used to check two slices are equal or not
	// s1 := []string{"A", "B", "C", "D", "E"}
	// s2 := []string{"D", "E", "F"}
	// result := reflect.DeepEqual(s1, s2)
	// fmt.Println(result)

	//// DeepEqual is used to check two arrays are equal or not
	// n1 := [5]int{1, 2, 3, 4, 5}
	// n2 := [5]int{1, 2, 3, 4, 5}
	// result = reflect.DeepEqual(n1, n2)
	// fmt.Println(result)

	//// DeepEqual is used to check two structures are equal or not
	// m1 := mobile{500.50, "red"}
	// m2 := mobile{400.50, "black"}
	// result = reflect.DeepEqual(m1, m2)
	// fmt.Println(result)

	//---------------------------------------------------------------------------------
	// Reflect.Swapper()

	// theList := []int{1, 2, 3, 4, 5}
	// swap := reflect.Swapper(theList)
	// fmt.Printf("Original Slice :%v\n", theList)

	//// Swapper() function is used to swaps the elements of slice
	// swap(1, 3)
	// fmt.Printf("After Swap :%v\n", theList)

	//// Reversing a slice using Swapper() function
	// for i := 0; i < len(theList)/2; i++ {
	// 	swap(i, len(theList)-1-i)
	// }
	// fmt.Printf("After Reverse :%v\n", theList)

	//---------------------------------------------------------------------------------
	// Reflect.Indirect()

	// val1 := []int{1, 2, 3, 4}
	// val2 := reflect.ValueOf(&val1)
	// fmt.Println("&val2 type : ", val2.Kind())

	// indirectI := reflect.Indirect(val2)
	// fmt.Println("indirectI  type : ", indirectI.Kind())
	// fmt.Println("indirectI  value : ", indirectI)

	//---------------------------------------------------------------------------------
	// Reflect.Interface()

	// t := reflect.TypeOf(5)

	// arr := reflect.ArrayOf(4, t)
	// inst := reflect.New(arr).Interface().(*[4]int)

	// for i := 1; i <= 4; i++ {
	// 	inst[i-1] = i * i
	// }

	// fmt.Println(inst)

	//---------------------------------------------------------------------------------
	// Reflect.MakeSlice()

	// var str []string
	// strType := reflect.ValueOf(&str)
	// newSlice := reflect.MakeSlice(reflect.Indirect(strType).Type(), 10, 15)

	// fmt.Println("Kind :", newSlice.Kind())
	// fmt.Println("Length :", newSlice.Len())
	// fmt.Println("Capacity :", newSlice.Cap())

	//---------------------------------------------------------------------------------
	// Reflect.MakeMap()

	// var str map[string]string
	// strType := reflect.ValueOf(&str)
	// newMap := reflect.MakeMap(reflect.Indirect(strType).Type())

	// fmt.Println("Kind :", newMap.Kind())

	//---------------------------------------------------------------------------------
	// Reflect.MakeChan()

	// var str chan string
	// strType := reflect.ValueOf(&str)
	// newChannel := reflect.MakeChan(reflect.Indirect(strType).Type(), 512)

	// fmt.Println("Kind :", newChannel.Kind())
	// fmt.Println("Capacity :", newChannel.Cap())

	//---------------------------------------------------------------------------------
	// Reflect.MakeFunc()

	// type Sum func(int64, int64) int64
	// t := reflect.TypeOf(Sum(nil))
	// fmt.Println("Type T :", t)

	// mul := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
	// 	a := args[0].Int()
	// 	b := args[1].Int()
	// 	return []reflect.Value{reflect.ValueOf(a + b)}
	// })

	// fn, ok := mul.Interface().(Sum)
	// if !ok {
	// 	return
	// }
	// fmt.Println(fn(5, 6))

	//---------------------------------------------------------------------------------

}
