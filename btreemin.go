package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return BTreeMin(root.Left)
	}
	return root
}
