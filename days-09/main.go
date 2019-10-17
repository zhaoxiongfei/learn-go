package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyInt64 int64

func (f MyInt64) Abs() uint64 {
	if 0 <= f {
		return uint64(f)
	}
	return uint64(0 - f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	var a, b MyInt64
	a = -32
	b = 64
	fmt.Println("a =", a, "a.Abs =", a.Abs())
	fmt.Println("b =", b, "b.Abs =", b.Abs())
}
