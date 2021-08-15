package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		if root.Left.Data > root.Data {
			return false
		}
		if root.Right.Data <= root.Data {
			return false
		}
	}
	if !BTreeIsBinary(root.Left) {
		return false
	} else if !BTreeIsBinary(root.Right) {
		return false
	}
	return true
}
