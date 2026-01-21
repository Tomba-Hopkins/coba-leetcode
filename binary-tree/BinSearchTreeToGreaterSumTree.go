package binarytree

// Question :

// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
// As a reminder, a binary search tree is a tree that satisfies these constraints:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
// Example 1:
// Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
// Example 2:
// Input: root = [0,null,1]
// Output: [1,null,1]
// Constraints:
// The number of nodes in the tree is in the range [1, 100].
// 0 <= Node.val <= 100
// All the values in the tree are unique.

// Answer :
type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

// pseudocode from gpt
// sum = 0
// function dfs(node):
//     if node == null:
//         return
//     dfs(node.right)
//     sum = sum + node.val
//     node.val = sum
//     dfs(node.left)
// call dfs(root)
// return root

func BinSearchTreeToGreaterSumTree(root *TreeNode2) *TreeNode2 {
	sum := 0
	var dfs func(*TreeNode2) // depth first search
	dfs = func(tn *TreeNode2) {
		if tn == nil {
			return
		}
		dfs(tn.Right)
		sum += tn.Val
		tn.Val = sum
		dfs(tn.Left)
	}
	dfs(root)
	return root

}

func printNih(root *TreeNode2) {
	if root == nil {
		return
	}
	printNih(root.Left)
	print(root.Val, " ")
	printNih(root.Right)
}

// func main() {
// 	root := &TreeNode2{Val: 4}
// 	root.Left = &TreeNode2{Val: 1}
// 	root.Right = &TreeNode2{Val: 6}
// 	root.Left.Left = &TreeNode2{Val: 0}
// 	root.Left.Right = &TreeNode2{Val: 2}
// 	root.Left.Right.Right = &TreeNode2{Val: 3}
// 	root.Right.Left = &TreeNode2{Val: 5}
// 	root.Right.Right = &TreeNode2{Val: 7}
// 	root.Right.Right.Right = &TreeNode2{Val: 8}
// 	BinSearchTreeToGreaterSumTree(root)
// 	printNih(root)
// }
