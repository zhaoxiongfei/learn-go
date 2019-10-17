package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println("before scale: v =", v, "abs =", v.Abs())
	v.Scale(5)
	fmt.Println("after scale: v =", v, "abs =", v.Abs())
}
