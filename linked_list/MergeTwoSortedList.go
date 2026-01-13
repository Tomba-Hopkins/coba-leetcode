package linkedlist

// Question :

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:
// Input: list1 = [], list2 = []
// Output: []
// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]
// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

// Answer :
import "fmt"

type ListNode6 struct {
	Val  int
	Next *ListNode6
}

// pseudocode from gpt
// buat dummy node
// tail = dummy
// while list1 != null AND list2 != null:
//     if list1.val <= list2.val:
//         tail.next = list1
//         list1 = list1.next
//     else:
//         tail.next = list2
//         list2 = list2.next
//     tail = tail.next
// if list1 masih ada:
//     tail.next = list1
// if list2 masih ada:
//     tail.next = list2
// return dummy.next

func MergeTwoSortedList(list1 *ListNode6, list2 *ListNode6) *ListNode6 {
	dummy := &ListNode6{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return dummy.Next
}

func buatinAkuNode(a []int) *ListNode6 {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode6{Val: a[0]}
	curr := head
	for i := 1; i < len(a); i++ {
		curr.Next = &ListNode6{Val: a[i]}
		curr = curr.Next
	}
	return head
}

func buatinPrint(h *ListNode6) {
	for h != nil {
		fmt.Print(h.Val, " ")
		h = h.Next
	}
	fmt.Println()
}

// func main() {
// 	a1, a2 := []int{1, 2, 4}, []int{1, 3, 4}
// 	h1, h2 := buatinAkuNode(a1), buatinAkuNode(a2)
// 	buatinPrint(MergeTwoSortedList(h1, h2))
// }
