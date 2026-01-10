package linkedlist

// Question :

// You are given the head of a linked list, which contains a series of integers separated by 0's. The beginning and end of the linked list will have Node.val == 0.
// For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. The modified list should not contain any 0's.
// Return the head of the modified linked list.
// Example 1:
// Input: head = [0,3,1,0,4,5,2,0]
// Output: [4,11]
// Explanation:
// The above figure represents the given linked list. The modified list contains
// - The sum of the nodes marked in green: 3 + 1 = 4.
// - The sum of the nodes marked in red: 4 + 5 + 2 = 11.
// Example 2:
// Input: head = [0,1,0,3,0,2,2,0]
// Output: [1,3,4]
// Explanation:
// The above figure represents the given linked list. The modified list contains
// - The sum of the nodes marked in green: 1 = 1.
// - The sum of the nodes marked in red: 3 = 3.
// - The sum of the nodes marked in yellow: 2 + 2 = 4.
// Constraints:
// The number of nodes in the list is in the range [3, 2 * 105].
// 0 <= Node.val <= 1000
// There are no two consecutive nodes with Node.val == 0.
// The beginning and end of the linked list have Node.val == 0.

// Answer :

// pseudocode dari chatgpt
// current = head
// sum = 0
// tail = dummy
// while current != null:
//     if current.val == 0:
//         if sum > 0:
//             tail.next = new Node(sum)
//             tail = tail.next
//             sum = 0
//     else:
//         sum += current.val
//     current = current.next
// return dummy.next

import "fmt"

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func MergesNodesInBetweenZero(head *ListNode3) *ListNode3 {
	curr := head
	sum := 0
	dummy := &ListNode3{}
	tail := dummy
	for curr != nil {
		if curr.Val == 0 {
			if sum > 0 {
				tail.Next = &ListNode3{Val: sum}
				tail = tail.Next
				sum = 0
			}
		} else {
			sum += curr.Val
		}
		curr = curr.Next
	}
	return dummy.Next
}

func buatinNodeDong(arr []int) *ListNode3 {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode3{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode3{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printLahHeran(head *ListNode3) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

// func main() {
// 	a := []int{0, 3, 1, 0, 4, 5, 2, 0}
// 	h := buatinNodeDong(a)
// 	printLahHeran(h)
// 	printLahHeran(MergesNodesInBetweenZero(h))
// }
