package binarytree

// Question :

// You are given the root of a binary search tree (BST) and an integer val.
// Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.
// Example 1:
// Input: root = [4,2,7,1,3], val = 2
// Output: [2,1,3]
// Example 2:
// Input: root = [4,2,7,1,3], val = 5
// Output: []
// Constraints:
// The number of nodes in the tree is in the range [1, 5000].
// 1 <= Node.val <= 107
// root is a binary search tree.
// 1 <= val <= 107

// Answer :
type TreeNode5 struct {
	Val   int
	Left  *TreeNode5
	Right *TreeNode5
}

// pseudocode from gpt
// function SearchInAnBinarySearchTree(node, val):
//     if node == null:
//         return null
//     if node.value == val:
//         return node
//     if val < node.value:
//         return SearchInAnBinarySearchTree(node.left, val)
//     else:
//         return SearchInAnBinarySearchTree(node.right, val)

func SearchInAnBinarySearchTree(root *TreeNode5, val int) *TreeNode5 {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return SearchInAnBinarySearchTree(root.Left, val)
	} else {
		return SearchInAnBinarySearchTree(root.Right, val)
	}

}

func printNoh(root *TreeNode5) {
	if root == nil {
		return
	}

	printNoh(root.Left)
	print(" ", root.Val, " ")
	printNoh(root.Right)

}

// func main() {
// 	root := &TreeNode5{Val: 4}
// 	root.Left = &TreeNode5{Val: 2}
// 	root.Right = &TreeNode5{Val: 7}
// 	root.Left.Left = &TreeNode5{Val: 1}
// 	root.Left.Right = &TreeNode5{Val: 3}
// 	printNoh(SearchInAnBinarySearchTree(root, 2))
// }
