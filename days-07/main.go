package main

import (
	"fmt"
	"strings"
)

// 统计字符串中单词的数量
func wordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, v := range strings.Fields(s) {
		ans[v]++
	}
	return ans
}

func main() {
	fmt.Println(wordCount("hello world hahha hello world world"))
}
