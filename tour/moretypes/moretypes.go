package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x*y - 1)
		}
		a[x] = b
	}

	return a
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s == ", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	fmt.Println("s[1:4] == ", s[1:4])
	fmt.Println("s[:3] == ", s[:3])
	fmt.Println("s[4:] == ", s[4:])

	aa := make([]int, 5)
	printSlice("aa", aa)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if nil == z {
		fmt.Println("z is nil!")
	}

	aa = append(aa, 1)
	printSlice("aa", aa)

	aa = append(aa, 2)
	printSlice("aa", aa)

	aa = append(aa, 3, 4, 5)
	printSlice("aa", aa)

	var bb []int
	bb = append(bb, 1)
	printSlice("bb", bb)

	bb = append(bb, 2)
	printSlice("bb", bb)

	bb = append(bb, 3, 4, 5)
	printSlice("bb", bb)

	bb = append(bb, 6, 7)
	printSlice("bb", bb)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	pic.Show(Pic)
}
