package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Redstone", 39}
	z := Person{"Garbin", 35}
	v := []int{2, 3, 4, 5}
	// v 并没有实现 String 也可以通过 fmt.Println 输出，为什么？
	// Println 不是单一的一种模式，里面有很多方式将不同类型的变量转换为 string
	fmt.Println(a, z, v)
}
