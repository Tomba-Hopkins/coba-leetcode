package linkedlist

import "fmt"

// Question :
// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.
// Example 2:
// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
// Constraints:
// The number of nodes in the list is in the range [1, 100].
// 1 <= Node.val <= 100

// Answer :

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleOfTheLinkedList(head *ListNode) *ListNode {
	slow, fast := head, head // dia bukan array, dia ga pake index2, tapi posisi node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow // 1 2 3 4 5 -> fast: 1, 3, 5 (abis) -> slow: 1, 2, 3 (ambil sisanya yg belum di lewati: 3, 4, 5)
}

func bikinListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head

}

func printLis(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// func main() {
// 	h := bikinListNode([]int{1, 2, 3, 4, 5})
// 	printLis(h)
// 	mid := MiddleOfTheLinkedList(h)
// 	printLis(mid)
// }
