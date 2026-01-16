package linkedlist

// Question :

// You are given two integers m and n, which represent the dimensions of a matrix.
// You are also given the head of a linked list of integers.
// Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix. If there are remaining empty spaces, fill them with -1.
// Return the generated matrix.
// Example 1:
// Input: m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
// Output: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
// Explanation: The diagram above shows how the values are printed in the matrix.
// Note that the remaining spaces in the matrix are filled with -1.
// Example 2:
// Input: m = 1, n = 4, head = [0,1,2]
// Output: [[0,1,2,-1]]
// Explanation: The diagram above shows how the values are printed from left to right in the matrix.
// The last space in the matrix is set to -1.
// Constraints:
// 1 <= m, n <= 105
// 1 <= m * n <= 105
// The number of nodes in the list is in the range [1, m * n].
// 0 <= Node.val <= 1000

// Answer :

type ListNode8 struct {
	Val  int
	Next *ListNode8
}

// pseudo from gpt
// function SpiralMatrix4(m, n, head):
//     create matrix[m][n] filled with -1

//     top = 0
//     bottom = m - 1
//     left = 0
//     right = n - 1

//     while head is not null AND top <= bottom AND left <= right:

//         // 1. ke kanan
//         for col from left to right:
//             if head is null: break
//             matrix[top][col] = head.value
//             head = head.next
//         top = top + 1

//         // 2. ke bawah
//         for row from top to bottom:
//             if head is null: break
//             matrix[row][right] = head.value
//             head = head.next
//         right = right - 1

//         // 3. ke kiri
//         for col from right down to left:
//             if head is null: break
//             matrix[bottom][col] = head.value
//             head = head.next
//         bottom = bottom - 1

//         // 4. ke atas
//         for row from bottom down to top:
//             if head is null: break
//             matrix[row][left] = head.value
//             head = head.next
//         left = left + 1

//     return matrix

func SpiralMatrix4(m int, n int, head *ListNode8) [][]int {
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r[i][j] = -1
		}
	}
	top := 0
	bottom := m - 1
	left := 0
	right := n - 1

	for head != nil && top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			if head == nil {
				break
			}
			r[top][i] = head.Val
			head = head.Next
		}
		top++

		for i := top; i <= bottom; i++ {
			if head == nil {
				break
			}
			r[i][right] = head.Val
			head = head.Next
		}
		right--

		for i := right; i >= left; i-- {
			if head == nil {
				break
			}
			r[bottom][i] = head.Val
			head = head.Next

		}
		bottom--

		for i := bottom; i >= top; i-- {
			if head == nil {
				break
			}
			r[i][left] = head.Val
			head = head.Next
		}
		left++
	}

	return r
}
func buatinNodeGatauKeberapa(a []int) *ListNode8 {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode8{Val: a[0]}
	curr := head
	for i := 1; i < len(a); i++ {
		curr.Next = &ListNode8{Val: a[i]}
		curr = curr.Next
	}
	return head
}

// func main() {
// 	a := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
// 	h := buatinNodeGatauKeberapa(a)
// 	fmt.Println(SpiralMatrix4(3, 5, h))

// }
