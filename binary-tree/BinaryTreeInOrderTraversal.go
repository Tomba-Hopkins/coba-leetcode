package binarytree

// Question :

// Given the root of a binary tree, return the inorder traversal of its nodes' values.
// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,3,2]
// Explanation:
// Example 2:
// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
// Output: [4,2,6,5,7,1,3,9,8]
// Explanation:
// Example 3:
// Input: root = []
// Output: []
// Example 4:
// Input: root = [1]
// Output: [1]
// Constraints:
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
// Follow up: Recursive solution is trivial, could you do it iteratively?

// Answer :

type TreeNode7 struct {
	Val   int
	Left  *TreeNode7
	Right *TreeNode7
}

// pseudocode from gpt
// function inorder(node):
//     if node is null:
//         return
//     inorder(node.left)
//     dfs
//     inorder(node.right)

// dfs
// buat result kosong
// dfs(node):
//     kalau null → return
//     dfs(left)
//     result append node.val
//     dfs(right)
// panggil dfs(root)
// return result

func BinaryTreeInOrderTraversal(root *TreeNode7) (r []int) {
	if root == nil {
		return
	}
	BinaryTreeInOrderTraversal(root.Left)
	var dfs func(node *TreeNode7)
	dfs = func(node *TreeNode7) {
		if node == nil {
			return
		}
		dfs(node.Left)
		r = append(r, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	BinaryTreeInOrderTraversal(root.Right)
	return
}

// func main() {
// 	root := &TreeNode7{Val: 1}
// 	root.Right = &TreeNode7{Val: 2}
// 	root.Right.Left = &TreeNode7{Val: 3}
// 	fmt.Println(BinaryTreeInOrderTraversal(root))
// }
