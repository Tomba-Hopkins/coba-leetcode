package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// pseudocode from gpt
// function MaxTwinSumOfAnLinkedList(head):
//     // cari tengah
//     slow = head
//     fast = head
//     while fast != null and fast.next != null:
//         slow = slow.next
//         fast = fast.next.next
//     // reverse setengah kedua
//     prev = null
//     curr = slow
//     while curr != null:
//         next = curr.next
//         curr.next = prev
//         prev = curr
//         curr = next
//     // hitung max twin sum
//     maxSum = 0
//     left = head
//     right = prev
//     while right != null:
//         maxSum = max(maxSum, left.val + right.val)
//         left = left.next
//         right = right.next
//     return maxSum

func MaxTwinSumOfAnLinkedList(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode = nil
	curr := slow

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	maxSum, left, right := 0, head, prev
	for right != nil {
		maxSum = max(maxSum, left.Val+right.Val)
		left = left.Next
		right = right.Next
	}
	return maxSum
}

func buatinNodeSambilLetdown(a []int) *ListNode {
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

func main() {
	a := []int{4, 2, 2, 3}
	h := buatinNodeSambilLetdown(a)
	fmt.Println(MaxTwinSumOfAnLinkedList(h))

}
