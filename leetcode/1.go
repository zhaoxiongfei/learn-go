package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		key := target - num
		if j, ok := numsMap[key]; ok {
			return []int{i, j}
		}
		numsMap[num] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
