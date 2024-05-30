package main

import "fmt"

func main() {
	values := []interface{}{
		7,
		"hello",
		false,
		make(chan int),
	}

	// Iterates over slice and compares each value to determine its type
	for _, v := range values {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")
		case chan int:
			fmt.Println("int channel")
		default:
			fmt.Println("unknown type")
		}
	}
}
