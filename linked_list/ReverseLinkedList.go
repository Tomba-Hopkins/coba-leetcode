package linkedlist

// Question :

// Given the head of a singly linked list, reverse the list, and return the reversed list.
// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
// Example 2:
// Input: head = [1,2]
// Output: [2,1]
// Example 3:
// Input: head = []
// Output: []
// Constraints:
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000
// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

// Answer :
import "fmt"

type ListNode5 struct {
	Val  int
	Next *ListNode
}

// pseudocode from gpt :
// function ReverseLinkedList(head):
//     prev = null
//     curr = head
//     while curr != null:
//         next = curr.next
//         curr.next = prev
//         prev = curr
//         curr = next
//     return prev

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func buatinNodeLahYaelahGimanaSih(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	curr := head
	for i := 1; i < len(a); i++ {
		curr.Next = &ListNode{Val: a[i]}
		curr = curr.Next
	}
	return head
}

func hadehPrintCoba(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5}
// 	h := buatinNodeLahYaelahGimanaSih(a)
// 	hadehPrintCoba(ReverseLinkedList(h))
// }
