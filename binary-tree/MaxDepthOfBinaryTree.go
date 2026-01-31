package binarytree

// Description :

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Example 2:
// Input: root = [1,null,2]
// Output: 2
// Constraints:
// The number of nodes in the tree is in the range [0, 104].
// -100 <= Node.val <= 100

// Answer :

// Pseudocode from gpt
// function MaxDepthOfBinaryTree(node):
//     if node == null:
//         return 0
//     leftDepth  = MaxDepthOfBinaryTree(node.left)
//     rightDepth = MaxDepthOfBinaryTree(node.right)
//     return max(leftDepth, rightDepth) + 1

type TreeNode10 struct {
	Val   int
	Left  *TreeNode10
	Right *TreeNode10
}

func MaxDepthOfBinaryTree(root *TreeNode10) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepthOfBinaryTree(root.Left)
	rightDepth := MaxDepthOfBinaryTree(root.Right)
	return max(leftDepth, rightDepth) + 1
}

// func main() {
// 	root := &TreeNode10{Val: 3}
// 	root.Left = &TreeNode10{Val: 9}
// 	root.Right = &TreeNode10{Val: 20}
// 	root.Right.Left = &TreeNode10{Val: 15}
// 	root.Right.Right = &TreeNode10{Val: 7}
// 	fmt.Println(MaxDepthOfBinaryTree(root))
// }
