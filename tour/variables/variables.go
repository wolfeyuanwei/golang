// variables.go
package main

import (
	"fmt"
)

//var c, python, java bool = true, true, false

func main() {
	var i, j int = 10, 2
	k := 3
	c, python, java := true, true, "yes!" //:=结构不能使用在函数外；
	fmt.Println(i, j, k, c, python, java)
}
