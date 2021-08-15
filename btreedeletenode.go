package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if root.Data > node.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if root.Data < node.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			store := root.Right
			root = nil
			return store
		} else if root.Right == nil {
			store := root.Left
			root = nil
			return store
		}
		store := BTreeMin(root.Right)
		root.Data = store.Data
		root.Right = BTreeDeleteNode(root.Right, store)
	}
	return root
}
