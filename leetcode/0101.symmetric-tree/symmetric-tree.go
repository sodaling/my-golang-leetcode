package problem0101

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(nodeA, nodeB *TreeNode) bool {
	if nodeA == nil && nodeB == nil {
		return true
	}
	if nodeA == nil || nodeB == nil {
		return false
	}
	if nodeA.Val == nodeB.Val {
		return isMirror(nodeA.Left, nodeB.Right) && isMirror(nodeA.Right, nodeB.Left)
	}
	return false

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
