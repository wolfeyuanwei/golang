package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1  = Vertex{1, 2}
	v2  = Vertex{X: 1}
	v3  = Vertex{}
	pv1 = &Vertex{1, 2}
)

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println("i = ", i)

	p = &j
	*p = *p / 37
	fmt.Println("j = ", j)

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 5
	fmt.Println("v.x = ", v.X)

	pv := &v
	pv.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, pv1, v2, v3)
}
