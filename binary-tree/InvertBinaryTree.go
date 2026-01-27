package binarytree

// Question :

// Given the root of a binary tree, invert the tree, and return its root.
// Example 1:
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Example 2:
// Input: root = [2,1,3]
// Output: [2,3,1]
// Example 3:
// Input: root = []
// Output: []
// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

// Answer :

import "fmt"

type TreeNode6 struct {
	Val   int
	Left  *TreeNode6
	Right *TreeNode6
}

// pseudocode from gpt
// function invert(node):
//     if node == null:
//         return null
//     temp = node.left
//     node.left = node.right
//     node.right = temp
//     invert(node.left)
//     invert(node.right)
//     return node

func InvertBinaryTree(root *TreeNode6) *TreeNode6 {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	InvertBinaryTree(root.Left)
	InvertBinaryTree(root.Right)
	return root
}

func printSemlenget(root *TreeNode6) {
	if root == nil {
		return
	}
	printSemlenget(root.Left)
	fmt.Print(root.Val, " ")
	printSemlenget(root.Right)

}

// func main() {
// 	root := &TreeNode6{Val: 4}
// 	root.Left = &TreeNode6{Val: 2}
// 	root.Right = &TreeNode6{Val: 7}
// 	root.Left.Left = &TreeNode6{Val: 1}
// 	root.Left.Right = &TreeNode6{Val: 3}
// 	root.Right.Left = &TreeNode6{Val: 6}
// 	root.Right.Right = &TreeNode6{Val: 9}
// 	printSemlenget(InvertBinaryTree(root))
// }
