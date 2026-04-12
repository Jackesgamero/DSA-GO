package tree

import (
	"fmt"
	"strconv"
)

type stringBinaryTreeNode struct {
	val   string
	left  *stringBinaryTreeNode
	right *stringBinaryTreeNode
}

func EvaluateBinaryExpressionTree(node *stringBinaryTreeNode) (float64, error) {
	if node == nil || node.val == "" {
		return 0, nil
	}

	left, err := EvaluateBinaryExpressionTree(node.left)
	if err != nil {
		return -1, fmt.Errorf("could not evaluate value on left node")
	}

	right, err := EvaluateBinaryExpressionTree(node.right)
	if err != nil {
		return -1, fmt.Errorf("could not evaluate value on right node")
	}

	switch node.val {
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	}

	val, err := strconv.ParseFloat(node.val, 64)
	if err != nil {
		return -1, fmt.Errorf("could not parse value")
	}

	return val, nil
}
