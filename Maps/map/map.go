package main

import "fmt"

func main() {
	// Creating an empty map using make
	m := make(map[string]int)

	// Adding key-value pairs to the map
	m["apple"] = 10
	m["banana"] = 20
	m["orange"] = 15

	// Accessing values from the map
	fmt.Println("Value of apple:", m["apple"])
	fmt.Println("Value of banana:", m["banana"])
	fmt.Println("Value of orange:", m["orange"])

	// Modifying values in the map
	m["apple"] = 12
	fmt.Println("Modified value of apple:", m["apple"])

	// Deleting a key-value pair from the map
	delete(m, "banana")
	fmt.Println("After deleting banana:", m)

	// Checking if a key exists in the map
	if _, ok := m["orange"]; ok {
		fmt.Println("Orange exists in the map")
	} else {
		fmt.Println("Orange does not exist in the map")
	}

	// Iterating over the map
	fmt.Println("Iterating over the map:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
