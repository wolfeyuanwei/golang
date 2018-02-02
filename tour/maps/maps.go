package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	c := strings.Fields(s)
	for _, v := range c {
		words[v] += 1
	}
	return words
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"google":    Vertex{37.4202, -122.08408},
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		out := sum1 + sum2
		sum1 = sum2
		sum2 = out
		return out
	}
}

func main() {
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{40.68433, -74.39967}
	//fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	mm := make(map[string]int)
	mm["Answer"] = 42
	fmt.Println("The value: ", mm["Answer"])
	mm["Answer"] = 48
	fmt.Println("The Value: ", mm["Answer"])
	delete(mm, "Answer")
	fmt.Println("The Value: ", mm["Answer"])

	v, ok := mm["Answer"]
	fmt.Println("The Value: ", v, "Present?", ok)

	wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
