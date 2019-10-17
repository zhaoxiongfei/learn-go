package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() uint64 {
	var c uint64 // 上一个值
	var n uint64 // 当前值
	c = 1
	n = 0
	return func() uint64 {
		t := c
		c = n
		n += t
		return n
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println("pos:", pos(i), "neg:", neg(-2*i))
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("Fibonacci", i, "=", fib())
	}
}
