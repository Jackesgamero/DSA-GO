package tree

func ReverseBinaryTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	root.Left = ReverseBinaryTree(root.Left)
	root.Right = ReverseBinaryTree(root.Right)

	root.Left, root.Right = root.Right, root.Left
	return root
}
