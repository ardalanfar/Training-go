/*---------------------------------------------------------------------------------
 Composition by embedding structs (Inheritance)
 Example_1
---------------------------------------------------------------------------------*/

package main

import "fmt"

func main() {

	//-------------------------------------------------------------------------------
	// Composition by embedding structs (Inheritance)

	// author11 := author0{
	// 	"Naveen",
	// 	"Ramanathan",
	// 	"Golang Enthusiast",
	// }
	// blogPost11 := blogPost0{
	// 	"Inheritance in Go",
	// 	"Go supports composition instead of inheritance",
	// 	author11,
	// }

	// blogPost11.details0()

	// -------------------------------------------------------------------------------
	// Example_1

	// author1 := author{
	// 	"Naveen",
	// 	"Ramanathan",
	// 	"Golang Enthusiast",
	// }
	// blogPost1 := blogPost{
	// 	"Inheritance in Go",
	// 	"Go supports composition instead of inheritance",
	// 	author1,
	// }
	// blogPost2 := blogPost{
	// 	"Struct instead of Classes in Go",
	// 	"Go does not support classes but methods can be added to structs",
	// 	author1,
	// }
	// blogPost3 := blogPost{
	// 	"Concurrency",
	// 	"Go is a concurrent language and not a parallel one",
	// 	author1,
	// }

	// w := website{
	// 	blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	// }

	// w.contents()

	//-------------------------------------------------------------------------------
}

//-------------------------------------------------------------------------------
// Composition by embedding structs (Inheritance)

type author0 struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost0 struct {
	title   string
	content string
	author0
}

func (a author0) fullName0() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (b blogPost0) details0() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName0())
	fmt.Println("Bio: ", b.bio)
}

//-------------------------------------------------------------------------------
// Example_1

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (p blogPost) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

//-------------------------------------------------------------------------------
