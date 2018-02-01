package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World!")

	fmt.Print("Hello ")

	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done!!!")
}