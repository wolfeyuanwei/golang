// split.go
package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 2 / 6
	y = sum - x
	return x, y
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(split(17))
}
