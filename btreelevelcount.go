package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 1
	}
	var leftCount int
	if root.Left != nil {
		leftCount = BTreeLevelCount(root.Left)
	}
	var rightCount int
	if root.Right != nil {
		rightCount = BTreeLevelCount(root.Right)
	}
	if leftCount > rightCount {
		return leftCount + 1
	} else {
		return rightCount + 1
	}
}
