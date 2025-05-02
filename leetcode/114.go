package leetcode

// 114. Flatten Binary Tree to Linked List
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
////
// We processes the tree with reverse pre-order right -> left -> root,
// so that we cna build the list from the tail backward.

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func Sol_medium_114(root *TreeNode) {
	var prev *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		dfs(node.Left)

		node.Right = prev
		node.Left = nil
		prev = node
	}

	dfs(root)
}
