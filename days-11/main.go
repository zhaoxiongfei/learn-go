package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if 0 <= f {
		return float64(f)
	}
	return float64(0 - f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Pi)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了Abser
	fmt.Println("a =", a, "a.Abs() =", a.Abs())

	a = &v // v *Vertex 实现了Abser
	fmt.Println("a =", a, "a.Abs() =", a.Abs())

	/*
		a = v // v Vertex 没有实现了Abser, 执行会报错
		fmt.Println("a =", a, "a.Abs() =", a.Abs())
	*/
}
