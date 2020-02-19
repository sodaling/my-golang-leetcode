package problem0094

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var visitArray []*TreeNode

	curNode := root
	for curNode != nil || len(visitArray) != 0 {
		for curNode != nil {
			visitArray = append(visitArray, curNode)
			curNode = curNode.Left
		}
		pop := visitArray[len(visitArray)-1]
		visitArray = visitArray[:len(visitArray)-1]

		ret = append(ret, pop.Val)

		if pop.Right != nil {
			curNode = pop.Right
		}
	}
	return ret
}
