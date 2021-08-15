package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	levels := BTreeLevelCount(root)
	for i := 1; i < levels+1; i++ {
		applyTolevel(root, f, i)
	}
}

func applyTolevel(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		applyTolevel(root.Left, f, level-1)
		applyTolevel(root.Right, f, level-1)
	}
}
