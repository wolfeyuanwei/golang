package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.1415

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	var i int
	var ff float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, ff, b, s)

	var x, y int = 3, 4
	ff = math.Sqrt(float64(x*x + y*y))
	var z int = int(ff)
	fmt.Println(x, y, z, ff)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}
