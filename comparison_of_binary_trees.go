package main

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var res bool = false

	if p == nil && q == nil {
		res = true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			res = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

	return res
}
