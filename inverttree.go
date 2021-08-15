package piscine

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

func InvertTree(root *TNode) *TNode {
	if root == nil {
		return nil
	}
	left := InvertTree(root.Left)
	right := InvertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
