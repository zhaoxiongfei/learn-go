package main

import "fmt"

func main() {
	fmt.Println("hello world")
	sentence := "Luke, I'm not your Father"
	fmt.Println(sentence[:10])
	for i := 0; i < 10; i += 1 {
		fmt.Println("i =", i)
	}

	kvs := map[string]string{
		"name":    "Redstone",
		"website": "http://xiongfei.me",
	}

	for key, value := range kvs {
		fmt.Println(key, value)
	}

	Post := map[string]interface{}{
		"Id":         300,
		"Title":      "Moving from NodeJS to GO",
		"Author":     "Christopher Ganga",
		"Difficulty": "beginner",
		"Start":      3.47,
		"Followers":  []string{"Redstone", "Open-node"},
	}

	fmt.Println(Post)

	fmt.Println(Post["Id"])
	fmt.Println(Post["Title"])
	fmt.Println(Post["Followers"])

	type Person struct {
		ID      int
		Name    string
		Age     int
		Parents []int
	}

	p := Person{
		ID:      30,
		Name:    "Redstone",
		Age:     40,
		Parents: []int{20, 21},
	}

	fmt.Println(p)
	fmt.Println(p.ID)
	fmt.Println(p.Name)
	fmt.Println(p.Parents)
}
