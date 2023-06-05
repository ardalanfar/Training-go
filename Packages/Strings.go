/*---------------------------------------------------------------------------------
  ToUpper, ToLower, Title
  String Search (hasprefix,hassub)
  Length
  String Manipulation (join,replacall,replace)
  Indexing
  Repeat
  Comparison (EqualFold,compare)
  Contains , count , ContainsRune
  Trim
  Split
  Json
--------------------------------------------------------------------------------*/

package main

func main() {

	//----------------------------------------------------------------------------
	// ToUpper, ToLower, Title

	// str := "Hello worD"

	// fmt.Println(strings.ToLower(str))
	// fmt.Println(strings.ToUpper(str))
	// fmt.Println(strings.Title(str))

	//----------------------------------------------------------------------------
	// String Search (hasprefix,hassub)

	// str := "Sammy Shark"

	// fmt.Println(strings.HasPrefix(str, "Sammy"))
	// fmt.Println(strings.HasSuffix(str, "Sharkk"))

	//----------------------------------------------------------------------------
	// Length

	// openSource := "Sammy contributes to open source."
	// fmt.Println(len(openSource))

	//----------------------------------------------------------------------------
	// String Manipulation (join,replacall,replace)

	// fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ","))

	// str := "Sammy has a has balloon."
	// s := strings.Split(str, " ")
	// fmt.Println(s[0])

	//fmt.Println(strings.ReplaceAll(balloon, "has", "had"))

	//fmt.Println(strings.Replace(balloon, "has", "had", 1))

	//----------------------------------------------------------------------------
	// Indexing

	// msg := "I saw a fox in the forest. The fox had brown fur. I like foxes."

	// fmt.Println(strings.Index(msg, "fox"))
	// fmt.Println(strings.LastIndex(msg, "fox"))

	//----------------------------------------------------------------------------
	// Repeat

	// str := "falcon"
	// fmt.Println(strings.Repeat(str+" ", 5))

	//----------------------------------------------------------------------------
	// Comparison (EqualFold,compare)

	// w1 := "falcon"
	// w2 := "Falcon"

	// if strings.Compare(w1, w2) == 0 {

	// 	fmt.Println("The words are equal")
	// } else {

	// 	fmt.Println("The words are not equal")
	// }

	// if strings.EqualFold(w1, w2) {

	// 	fmt.Println("The words are equal")
	// } else {

	// 	fmt.Println("The words are not equal")
	// }

	//----------------------------------------------------------------------------
	// Contains , count , ContainsRune

	// str := "Sammy Shark"
	// str2 := "a blue 🐋"
	// r := '🐋'

	// fmt.Println(strings.Contains(str, "Sh"))
	// fmt.Println(strings.Count(str, "Sh"))
	// fmt.Println(strings.ContainsRune(str2, r))

	//----------------------------------------------------------------------------
	// Trim

	// str := ".an old falcon!"
	// cutset := ".!"
	// fmt.Println(strings.Trim(str, cutset))
	// fmt.Println(strings.TrimLeft(str, cutset))
	// fmt.Println(strings.TrimRight(str, cutset))

	// str2 := "\t\tand old falcon\n"
	// fmt.Println(strings.TrimSpace(str2))

	// str3 := "--and old falcon--"
	// fmt.Println(strings.TrimPrefix(str3, "--"))
	// fmt.Println(strings.TrimSuffix(str3, "--"))

	//----------------------------------------------------------------------------
	// Split

	// str := "3,4,5,6,7,8,9,10,11"
	// fmt.Println(strings.Split(str, "--"))

	//----------------------------------------------------------------------------
	// Json

	// words := []string{"an", "old", "falcon", "in", "the", "sky"}

	// fmt.Println(strings.Join(words, " "))

	//----------------------------------------------------------------------------
}
