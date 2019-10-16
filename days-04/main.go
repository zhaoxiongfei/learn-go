package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(math.Pi)
	defer fmt.Println("world")

	for i := 0; i < 10; i += 1 {
		defer fmt.Println(i)
	}

	fmt.Println("hello")

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(p)

	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		v4 = &Vertex{1, 2} // 类型为 *Vertex
	)

	fmt.Println(v1, v2, v3, v4)
	var arr1 = []int{1, 2, 3, 4, 5}
	fmt.Println("len(arr1)", len(arr1))

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("len(arr2)", len(arr2))

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q ==", q)
	fmt.Println("q[1:4] ==", q[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("q[:3] ==", q[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("q[4:] ==", q[4:])
}
