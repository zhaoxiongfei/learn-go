package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var sets = make(map[byte]bool)
	var max, m int
	for i := 0; i < len(s); i++ {
		for sets[s[i]] {
			sets[s[i-max]] = false
			max--
		}
		sets[s[i]] = true
		max++
		if m < max {
			m = max
		}
	}
	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
