package main

import "fmt"

func main() {
	// Declare a variable 'num' and initialize it with a value
	num := 42
	fmt.Println("Value of num:", num) // Output: Value of num: 42

	// Declare a pointer variable 'ptr' of type int
	var ptr *int

	// Assign the memory address of 'num' to 'ptr'
	ptr = &num

	// Print the memory address of 'num' and the value stored at that address
	fmt.Println("Memory address of num:", ptr)    // Output: Memory address of num: 0x123456 (example address)
	fmt.Println("Value pointed to by ptr:", *ptr) // Output: Value pointed to by ptr: 42

	// Modify the value of 'num' indirectly using the pointer
	*ptr = 100

	// Print the updated value of 'num'
	fmt.Println("Updated value of num:", num) // Output: Updated value of num: 100
}
