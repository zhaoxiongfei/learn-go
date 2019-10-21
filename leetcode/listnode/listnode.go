package listnode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s []string
	var curr = l
	for curr != nil {
		s = append(s, strconv.Itoa(curr.Val))
		curr = curr.Next
	}

	return strings.Join(s, " -> ")
}

func Create(values []int) *ListNode {
	pre := &ListNode{0, nil}
	var curr = pre
	for i := range values {
		curr.Next = &ListNode{values[i], nil}
		curr = curr.Next
	}

	return pre.Next
}
