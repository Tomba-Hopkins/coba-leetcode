package binarytree

// Question :

// Given the root of a binary tree, return the number of nodes where the value of the node is equal to the average of the values in its subtree.
// Note:
// The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
// A subtree of root is a tree consisting of root and all of its descendants.
// Example 1:
// Input: root = [4,8,5,0,1,null,6]
// Output: 5
// Explanation:
// For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
// For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
// For the node with value 0: The average of its subtree is 0 / 1 = 0.
// For the node with value 1: The average of its subtree is 1 / 1 = 1.
// For the node with value 6: The average of its subtree is 6 / 1 = 6.
// Example 2:
// Input: root = [1]
// Output: 1
// Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.
// Constraints:
// The number of nodes in the tree is in the range [1, 1000].
// 0 <= Node.val <= 1000
// Answer :
type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

// inti soal :
// - sum node bersama bawah2 nya misal ada 5 dan 6 sebagai anaknya
// - 5 + 6 = 11 -> next bagi dengan length nya [5, 6] -> 2
// - 11 / 2 = floor nya 5 -> node awal 5-> if penghitungan barusan == node awal -> count++

// pseudocode from gpt
// global count = 0
// function dfs(node):
//     if node == null:
//         return (sum=0, count=0)
//     leftSum, leftCount = dfs(node.left)
//     rightSum, rightCount = dfs(node.right)
//     totalSum = leftSum + rightSum + node.val
//     totalCount = leftCount + rightCount + 1
//     avg = totalSum / totalCount (dibulatkan ke bawah)
//     if avg == node.val:
//         count++
//     return (totalSum, totalCount)
// call dfs(root)
// return count

func CountNodesEqualToAverageOfSubtree(root *TreeNode3) int {
	globalCount := 0
	var dfs func(node *TreeNode3) (int, int)
	dfs = func(node *TreeNode3) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)
		totalSum := leftSum + rightSum + node.Val
		totalCount := leftCount + rightCount + 1
		average := totalSum / totalCount
		if average == node.Val {
			globalCount++
		}
		return totalSum, totalCount
	}
	dfs(root)
	return globalCount
}

// func main() {
// 	root := &TreeNode3{Val: 4}
// 	root.Left = &TreeNode3{Val: 8}
// 	root.Right = &TreeNode3{Val: 5}
// 	root.Left.Left = &TreeNode3{Val: 0}
// 	root.Left.Right = &TreeNode3{Val: 1}
// 	root.Right.Right = &TreeNode3{Val: 6}
// 	fmt.Println(CountNodesEqualToAverageOfSubtree(root))
// }
