package binarytree

// Description :
// You are given two binary trees root1 and root2.
// Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.
// Return the merged tree.
// Note: The merging process must start from the root nodes of both trees.
// Example 1:
// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]
// Example 2:
// Input: root1 = [1], root2 = [1,2]
// Output: [2,2]
// Constraints:
// The number of nodes in both trees is in the range [0, 2000].
// -104 <= Node.val <= 104

// Answer :

import "fmt"

type TreeNode8 struct {
	Val   int
	Left  *TreeNode8
	Right *TreeNode8
}

// pseudocode from gpt
// function merge(node1, node2):
//     if node1 == null AND node2 == null:
//         return null
//     if node1 == null:
//         return copy of node2
//     if node2 == null:
//         return copy of node1
//     newNode = new TreeNode8()
//     newNode.val = node1.val + node2.val
//     newNode.left  = merge(node1.left,  node2.left)
//     newNode.right = merge(node1.right, node2.right)
//     return newNode

func MergeTwoBinaryTrees(root1 *TreeNode8, root2 *TreeNode8) *TreeNode8 {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	newRoot := &TreeNode8{}
	newRoot.Val = root1.Val + root2.Val
	newRoot.Left = MergeTwoBinaryTrees(root1.Left, root2.Left)
	newRoot.Right = MergeTwoBinaryTrees(root1.Right, root2.Right)
	return newRoot
}

func printLembarBaru(root *TreeNode8) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode8{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Print(node.Val, " ")
		queue = append(queue, node.Left, node.Right)
	}
}

// func main() {
// 	root1 := &TreeNode8{Val: 1}
// 	root1.Right = &TreeNode8{Val: 2}
// 	root1.Left = &TreeNode8{Val: 3}
// 	root1.Left.Left = &TreeNode8{Val: 5}
// 	root2 := &TreeNode8{Val: 2}
// 	root2.Left = &TreeNode8{Val: 1}
// 	root2.Right = &TreeNode8{Val: 3}
// 	root2.Left.Left = &TreeNode8{Val: 4}
// 	root2.Right.Right = &TreeNode8{Val: 4}
// 	printLembarBaru(MergeTwoBinaryTrees(root1, root2))
// }
