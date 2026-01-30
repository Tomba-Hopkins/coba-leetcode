package binarytree

// Description :

// Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only one right child.
// Example 1:
// Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// Example 2:
// Input: root = [5,1,7]
// Output: [1,null,5,null,7]
// Constraints:
// The number of nodes in the given tree will be in the range [1, 100].
// 0 <= Node.val <= 1000

// Answer :

import "fmt"

type TreeNode9 struct {
	Val   int
	Left  *TreeNode9
	Right *TreeNode9
}

// pseudocode from gpt
// resultRoot = null
// prev = null
// function inorder(node):
//     if node == null:
//         return
//     inorder(node.left)
//     if resultRoot == null:
//         resultRoot = node
//     else:
//         prev.right = node
//     node.left = null
//     prev = node
//     inorder(node.right)
// call inorder(root)
// return resultRoot

func IncreasingOrderSearchTree(root *TreeNode9) *TreeNode9 {
	var resultRoot *TreeNode9 = nil
	var prev *TreeNode9 = nil
	var inorder func(*TreeNode9)
	inorder = func(node *TreeNode9) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if resultRoot == nil {
			resultRoot = node
		} else {
			prev.Right = node
		}
		node.Left = nil
		prev = node
		inorder(node.Right)
	}
	inorder(root)
	return resultRoot
}

func printSamsak(root *TreeNode9) {
	if root == nil {
		return
	}
	printSamsak(root.Left)
	fmt.Print(root.Val, " ")
	printSamsak(root.Right)

}

// func main() {
// 	root := &TreeNode9{Val: 5}
// 	root.Left = &TreeNode9{Val: 3}
// 	root.Right = &TreeNode9{Val: 6}
// 	root.Left.Left = &TreeNode9{Val: 2}
// 	root.Left.Right = &TreeNode9{Val: 4}
// 	root.Left.Left.Left = &TreeNode9{Val: 1}
// 	root.Right.Right = &TreeNode9{Val: 8}
// 	root.Right.Right.Left = &TreeNode9{Val: 7}
// 	root.Right.Right.Right = &TreeNode9{Val: 9}
// 	printSamsak(IncreasingOrderSearchTree(root))
// }
