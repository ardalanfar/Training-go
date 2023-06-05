/*---------------------------------------------------------------------------------
  Description
  Accessing bytes and characters of a string
  Rune()
  Rune (Accessing bytes and characters of a string)
  String are Mutable by rune
  Revstring string by rune
  Unique String by rune

  Byte()

  Iterate String using Range Loop
  String Comparison
  String Concatenation
  String are Immutable
---------------------------------------------------------------------------------*/

package main

import (
	"fmt"
)

func main() {

	//---------------------------------------------------------------------------------
	// Description

	//1- Strings in Go are value types, not pointers as is the case with C strings
	//2- Strings can be converted to byte slices and back again
	//3-

	//---------------------------------------------------------------------------------
	// Accessing bytes and characters of a string

	//----------------------------------- //Expl_1
	// name := "Hello World"
	// printChars(name)
	// fmt.Printf("\n")
	// printBytes(name)

	//----------------------------------- //Expl_2 (Problem)
	// name := "Señor"
	// fmt.Printf("String: %s\n", name)
	// printChars(name)
	// fmt.Printf("\n")
	// printBytes(name)

	//---------------------------------------------------------------------------------
	// Rune()

	// str1 := "salam"
	// str2 := "سلام"
	// str3 := "ABC€"
	// str4 := "Señor"

	// r := []rune(str1)
	// fmt.Println(reflect.TypeOf(r).Kind())
	// fmt.Printf("%v\n", r) //default format
	// fmt.Printf("%U\n", r) //Unicode format
	// fmt.Printf("%x\n", r) //base 16(hexadecimal)

	// strR := string(r)
	// fmt.Println(strR)

	//---------------------------------------------------------------------------------
	// Rune (Accessing bytes and characters of a string)

	//----------------------------------- //Expl_1
	// name := "Hello World"
	// printCharsRune(name)
	// fmt.Printf("\n")
	// printBytes(name)
	// fmt.Printf("\n\n")

	//----------------------------------- //Expl_2 (True)
	// name := "Señor"
	// fmt.Printf("String: %s\n", name)
	// printCharsRune(name)
	// fmt.Printf("\n")
	// printBytes(name)

	//---------------------------------------------------------------------------------
	// String are Mutable by rune

	// a := "hello"
	// r := []rune(a)
	// r[0] = 's'
	// fmt.Println("res:", a)
	// fmt.Println("res:", string(r))

	//---------------------------------------------------------------------------------
	// Revstring string by rune

	// str := "hello word"
	// run := []rune(str)

	// for i := len(run) - 1; i >= 0; i-- {
	// 	fmt.Printf("%c", run[i])
	// }

	//---------------------------------------------------------------------------------
	// Unique String by rune

	//str := "hello word"
	//run := []rune(str)

	//for i := 0; i <= len(run)-1; i++ {
	//	res := true
	//	for j := i + 1; j <= len(run)-1; j++ {
	//		if run[i] == run[j] {
	//			res = false
	//			break
	//		}
	//	}
	//	if res != false {
	//		fmt.Printf("%c", run[i])
	//	} else {
	//		continue
	//	}
	//}

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	// Byte()

	//------------------------------------ Expl_1
	//byteSlice1 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}     //UTF-8 Encoded hex bytes of the string
	//byteSlice2 := []byte{67, 97, 102, 195, 169}           //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}

	//fmt.Println(string(byteSlice1))

	//------------------------------------ Expl_2
	//var ch1 byte = 65         //decimal
	//var ch2 byte = 0b1000001  //Binaria
	//var ch3 byte = 0o101      //Octadecimal
	//var ch4 byte = 0x41       //hexadecimal

	//fmt.Println(string(ch1))

	//------------------------------------ Expl_3
	//a := "ABC€"
	//b := []byte(a)
	//fmt.Println(b)   //[65 66 67 226 130 172]
	//fmt.Println(string(b))
	//fmt.Printf("%s", a)

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------
	// Iterate String using Range Loop

	// str1 := "salam"
	// str2 := "سلام"
	// str3 := "ABC€"
	// str4 := "Señor"

	// for index, charcter := range str4 {
	// 	fmt.Printf("index:%d - charcter:%c\n", index, charcter)
	// }

	//---------------------------------------------------------------------------------
	// String Comparison

	//string1 := "Go"
	//string2 := "Go"

	//if string1 == string2 {
	//	fmt.Printf("%s and %s are equal\n", str1, str2)
	//	return
	//}
	//fmt.Printf("%s and %s are not equal\n", str1, str2)

	//---------------------------------------------------------------------------------
	// String Concatenation

	// string1 := "Go"
	// string2 := "is awesome"
	// result1 := string1 + " " + string2
	// result2 := fmt.Sprintf("%s %s", string1, string2)
	// fmt.Println(result2)

	//---------------------------------------------------------------------------------
	// String are Immutable

	//a := "hello"
	//a[0] = "s"
	//fmt.Println(a) //Error (String are Immutable)

	//---------------------------------------------------------------------------------

}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsRune(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}
