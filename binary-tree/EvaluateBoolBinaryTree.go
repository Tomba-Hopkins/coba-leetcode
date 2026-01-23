package binarytree

// Question :

// You are given the root of a full binary tree with the following properties:
// Leaf nodes have either the value 0 or 1, where 0 represents False and 1 represents True.
// Non-leaf nodes have either the value 2 or 3, where 2 represents the boolean OR and 3 represents the boolean AND.
// The evaluation of a node is as follows:
// If the node is a leaf node, the evaluation is the value of the node, i.e. True or False.
// Otherwise, evaluate the node's two children and apply the boolean operation of its value with the children's evaluations.
// Return the boolean result of evaluating the root node.
// A full binary tree is a binary tree where each node has either 0 or 2 children.
// A leaf node is a node that has zero children.
// Example 1:
// Input: root = [2,1,3,null,null,0,1]
// Output: true
// Explanation: The above diagram illustrates the evaluation process.
// The AND node evaluates to False AND True = False.
// The OR node evaluates to True OR False = True.
// The root node evaluates to True, so we return true.
// Example 2:
// Input: root = [0]
// Output: false
// Explanation: The root node is a leaf node and it evaluates to false, so we return false.
// Constraints:
// The number of nodes in the tree is in the range [1, 1000].
// 0 <= Node.val <= 3
// Every node has either 0 or 2 children.
// Leaf nodes have a value of 0 or 1.
// Non-leaf nodes have a value of 2 or 3.

// Answer :
type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

// pseudocode from gpt
// function evaluate(node):
//     if node is leaf:
//         return (node.val == 1)
//     leftResult  = evaluate(node.left)
//     rightResult = evaluate(node.right)
//     if node.val == 2:
//         return leftResult OR rightResult
//     if node.val == 3:
//         return leftResult AND rightResult

func EvaluateBoolBinaryTree(root *TreeNode4) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	leftResult := EvaluateBoolBinaryTree(root.Left)
	rightResult := EvaluateBoolBinaryTree(root.Right)
	if root.Val == 2 {
		return leftResult || rightResult
	}
	if root.Val == 3 {
		return leftResult && rightResult
	}
	return true
}

// func main() {
// 	root := &TreeNode4{Val: 2}
// 	root.Left = &TreeNode4{Val: 1}
// 	root.Right = &TreeNode4{Val: 3}
// 	root.Right.Left = &TreeNode4{Val: 0}
// 	root.Right.Right = &TreeNode4{Val: 1}
// 	fmt.Println(EvaluateBoolBinaryTree(root))
// }
