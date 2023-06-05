/*---------------------------------------------------------------------------------
  General syntax
  string to int (Atoi)
  int/int32/int64 to string (Itoa)
  int/int32/int64 to string (FormatInt)
  string to int32/int64 (ParseInt)
  string to boolean (ParseBool)
  string to float64 (ParseBool)
---------------------------------------------------------------------------------*/

package main

var (
	num int     = 22
	str string  = "Hello"
	flt float32 = 10.33
	bl  bool    = true
)

func main() {

	//---------------------------------------------------------------------------------
	// General syntax

	// fmt.Println(float32(num)) //convert int to float32
	// fmt.Println(float64(num)) //convert int to float64

	// //fmt.Println(int(str))     //cannot convert
	// //fmt.Println(string(num))  //cannot convert

	// fmt.Println(int(flt))    //convert float to int
	// fmt.Println(int32(flt))  //convert float to int32
	// fmt.Println(int64(flt))  //convert float to int64

	//---------------------------------------------------------------------------------
	// string to int (Atoi)
	// func Atoi(s string) (int, error)

	// var str2 string = "1234"
	// i, _ := strconv.Atoi(str2)  //convert str to int

	// fmt.Println(i)
	// fmt.Println(reflect.TypeOf(i)) //int

	//---------------------------------------------------------------------------------
	// int/int32/int64 to string (Itoa)
	// func Itoa(i int) string

	// var i int = 1234
	// var i32 int32 = 1234
	// var i64 int64 = 42353245

	//str1 := strconv.Itoa(i)          //convert int to str
	//str2 := strconv.Itoa(int(i32))   //convert int32 to str
	//str3 := strconv.Itoa(int(i364))  //convert int64 to str

	// fmt.Println(str1)
	// fmt.Println(reflect.TypeOf(str1)) //string

	//---------------------------------------------------------------------------------
	// int/int32/int64 to string (FormatInt)
	// func FormatInt(i int64, base int) string

	// var i int = 345345
	// var i32 int32 = 346543523
	// var i64 int64 = 1234

	// str1 := strconv.FormatInt(int64(i), 10)   //convert int to str
	// str2 := strconv.FormatInt(int64(i32), 10) //convert int32 to str
	// str3 := strconv.FormatInt(i64, 10)        //convert int64 to str

	// fmt.Println(str1)
	// fmt.Println(reflect.TypeOf(str1)) //string

	//---------------------------------------------------------------------------------
	// string to int32/int64 (ParseInt)
	//func ParseInt(s string, base int, bitSize int)

	// var str string = "1234"
	// i32, _ := strconv.ParseInt(str, 10, 32) //convert str to int32
	// i64, _ := strconv.ParseInt(str, 10, 64) //convert str to int64

	// fmt.Println(i32)
	// fmt.Println(reflect.TypeOf(i32)) //int64

	//---------------------------------------------------------------------------------
	// string to boolean (ParseBool)
	// func ParseBool(str string) (bool, error)

	// strbool := "false"
	// strb, _ := strconv.ParseBool(strbool)  //convert str to boolean

	// fmt.Println(strb)
	// fmt.Println(reflect.TypeOf(strb)) //bool

	//---------------------------------------------------------------------------------
	// string to float64 (ParseBool)
	// ParseFloat(s string, bitSize int) (float64, error)

	// strflt := "6.7"
	// strf32, _ := strconv.ParseFloat(strflt, 32) //convert str to float32
	// strf64, _ := strconv.ParseFloat(strflt, 64) //convert str to float64

	// fmt.Println(strf32)
	// fmt.Println(reflect.TypeOf(strf32)) //float

	//---------------------------------------------------------------------------------
}
