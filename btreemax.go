package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return BTreeMax(root.Right)
	}
	return root
}
