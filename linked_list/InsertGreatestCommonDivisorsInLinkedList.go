package linkedlist

// Question :

// Given the head of a linked list head, in which each node contains an integer value.
// Between every pair of adjacent nodes, insert a new node with a value equal to the greatest common divisor of them.
// Return the linked list after insertion.
// The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.
// Example 1:
// Input: head = [18,6,10,3]
// Output: [18,6,6,2,10,1,3]
// Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes (nodes in blue are the inserted nodes).
// - We insert the greatest common divisor of 18 and 6 = 6 between the 1st and the 2nd nodes.
// - We insert the greatest common divisor of 6 and 10 = 2 between the 2nd and the 3rd nodes.
// - We insert the greatest common divisor of 10 and 3 = 1 between the 3rd and the 4th nodes.
// There are no more adjacent nodes, so we return the linked list.
// Example 2:
// Input: head = [7]
// Output: [7]
// Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes.
// There are no pairs of adjacent nodes, so we return the initial linked list.
// Constraints:
// The number of nodes in the list is in the range [1, 5000].
// 1 <= Node.val <= 1000

// Answer :
import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func InsertGreatestCommonDivisorsInLinkedList(head *ListNode2) *ListNode2 {
	curr := head
	for curr != nil && curr.Next != nil {
		g := greatestCommonDivisor(curr.Val, curr.Next.Val)
		newHead := &ListNode2{Val: g}
		newHead.Next = curr.Next
		curr.Next = newHead
		curr = newHead.Next
	}
	return head
}

func greatestCommonDivisor(a, b int) int { // rumus dari gpt
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func buatinNode1(arr []int) *ListNode2 {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode2{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode2{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printDong(head *ListNode2) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {

	a := []int{18, 6, 10, 3}
	n := buatinNode1(a)
	printDong(InsertGreatestCommonDivisorsInLinkedList(n))

}
