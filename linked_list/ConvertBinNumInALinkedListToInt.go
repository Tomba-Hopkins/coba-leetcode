package linkedlist

// Question :

// Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.
// Return the decimal value of the number in the linked list.
// The most significant bit is at the head of the linked list.
// Example 1:
// Input: head = [1,0,1]
// Output: 5
// Explanation: (101) in base 2 = (5) in base 10
// Example 2:
// Input: head = [0]
// Output: 0
// Constraints:
// The Linked List is not empty.
// Number of nodes will not exceed 30.
// Each node's value is either 0 or 1.

// Answer :

import (
	"fmt"
)

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 1, 0, 1 -> r = 0 -> r * 2 + head.val
// -> index1: 0 * 2 + 1
// -> index2: 1 * 2 + 0
// -> index3: 2 * 2 + 1 -> 5
func ConvertBinNumInALinkedListToInt(head *ListNode1) int {
	r := 0
	for head != nil {
		r = r*2 + head.Val
		head = head.Next
	}
	return r
}

func buatinNode(arr []int) *ListNode1 {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode1{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode1{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printNodeNyah(h *ListNode1) {
	for h != nil {
		fmt.Print(h.Val, " ")
		h = h.Next
	}
	fmt.Println()
}

// func main() {
// 	a := []int{1, 0, 1}
// 	nod := buatinNode(a)
// 	printNodeNyah(nod)
// 	fmt.Println(ConvertBinNumInALinkedListToInt(nod))
// }
