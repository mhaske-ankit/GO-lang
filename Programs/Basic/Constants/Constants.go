/*Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.*/

package main

import "fmt"

const ab = "Ganesh"
const date = 14.2

func main() {
	const new = "everyone"
	fmt.Println("Hello", new)
	fmt.Println("Happy", ab, "chatturti", date)

}
