package binarytree

// Question :

// Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].
// Example 1:
// Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
// Output: 32
// Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
// Example 2:
// Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
// Output: 23
// Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
// Constraints:
// The number of nodes in the tree is in the range [1, 2 * 104].
// 1 <= Node.val <= 105
// 1 <= low <= high <= 105
// All Node.val are unique.

// Answer :
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pseudocode from gpt
// function rangeSum(node):
//     if node == null:
//         return 0
//     if node.value < low:
//         return rangeSum(node.right)
//     if node.value > high:
//         return rangeSum(node.left)
//     return node.value
//            + rangeSum(node.left)
//            + rangeSum(node.right)

func RangeSumOfBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return RangeSumOfBST(root.Right, low, high)
	}
	if root.Val > high {
		return RangeSumOfBST(root.Left, low, high)
	}
	return root.Val + RangeSumOfBST(root.Left, low, high) + RangeSumOfBST(root.Right, low, high)
}

// explain:
// - nilai awal 10, 10 > low dan 10 < high
// - jadi dia masuk return paling bawah
// - ditambah nilai baru yang > low dan < high
// - left ada 7, right ada 15 -> 7 + 15 + 10 = 32

// func main() {
// 	root := &TreeNode{Val: 10}
// 	root.Left = &TreeNode{Val: 5}
// 	root.Right = &TreeNode{Val: 15}
// 	root.Left.Left = &TreeNode{Val: 3}
// 	root.Left.Right = &TreeNode{Val: 7}
// 	root.Right.Right = &TreeNode{Val: 18}
// 	fmt.Println(RangeSumOfBST(root, 7, 15))
// }
