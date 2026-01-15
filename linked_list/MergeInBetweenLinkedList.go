package linkedlist

// Question :

// You are given two linked lists: list1 and list2 of sizes n and m respectively.
// Remove list1's nodes from the ath node to the bth node, and put list2 in their place.
// The blue edges and nodes in the following figure indicate the result:
// Build the result list and return its head.
// Example 1:
// Input: list1 = [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
// Output: [10,1,13,1000000,1000001,1000002,5]
// Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.
// Example 2:
// Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
// Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
// Explanation: The blue edges and nodes in the above figure indicate the result.
// Constraints:
// 3 <= list1.length <= 104
// 1 <= a <= b < list1.length - 1
// 1 <= list2.length <= 104

// Answer :

import "fmt"

type ListNode7 struct {
	Val  int
	Next *ListNode7
}

// pseudocode from gpt
// function MergeInBetweenLinkedList(list1, a, b, list2):
//     curr = list1
//     index = 0
//     # 1. cari node ke-(a-1)
//     while index < a - 1:
//         curr = curr.next
//         index++
//     prevA = curr
//     # 2. lanjut cari node ke-(b+1)
//     while index < b + 1:
//         curr = curr.next
//         index++
//     afterB = curr
//     # 3. sambung prevA ke list2
//     prevA.next = list2
//     # 4. cari tail list2
//     tail = list2
//     while tail.next != null:
//         tail = tail.next
//     # 5. sambung tail list2 ke afterB
//     tail.next = afterB
//     return list1

func MergeInBetweenLinkedList(list1 *ListNode7, a int, b int, list2 *ListNode7) *ListNode7 {
	curr := list1
	index := 0

	// atur index yg dibuang
	for index < a-1 {
		curr = curr.Next
		index++
	}

	prevA := curr

	for index < b+1 {
		curr = curr.Next
		index++
	}

	afterB := curr

	// sambung ke list2
	prevA.Next = list2

	tail := list2
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = afterB
	return list1
}

func buatinNodeGatauLah(a []int) *ListNode7 {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode7{Val: a[0]}
	curr := head

	for i := 1; i < len(a); i++ {
		curr.Next = &ListNode7{Val: a[i]}
		curr = curr.Next
	}
	return head
}

func printGatauLah(head *ListNode7) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// func main() {
// 	a1, a2 := []int{10, 1, 13, 6, 9, 5}, []int{1000000, 1000001, 1000002}
// 	h1, h2 := buatinNodeGatauLah(a1), buatinNodeGatauLah(a2)
// 	printGatauLah(MergeInBetweenLinkedList(h1, 3, 4, h2))
// }
