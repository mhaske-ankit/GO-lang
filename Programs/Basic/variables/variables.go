package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)

	//Variables declared without an explicit initial value are given their zero value.
	var a int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", a, f, b, s)
}
