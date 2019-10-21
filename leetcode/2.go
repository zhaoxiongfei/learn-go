package main

import (
	"fmt"

	. "./listnode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 存放结果
	var ans = ListNode{0, nil}
	// 存放进位倍数 1, 10, 100, 1000, ....
	n1, n2, c := l1, l2, &ans
	for n1 != nil || n2 != nil {
		n := c.Next
		if n == nil {
			n = &ListNode{0, nil}
		}
		if n1 != nil {
			n.Val += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			n.Val += n2.Val
			n2 = n2.Next
		}
		c.Next = n
		if 9 < n.Val {
			n.Next = &ListNode{1, nil}
			n.Val -= 10
		}
		c = c.Next
	}
	return ans.Next
}

func main() {
	var l1 = Create([]int{2, 4, 3})
	var l2 = Create([]int{5, 6, 4})
	fmt.Println(addTwoNumbers(l1, l2))
}
