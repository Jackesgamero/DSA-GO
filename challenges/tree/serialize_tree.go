package tree

import (
	"strconv"
	"strings"
)

const (
	delimiter = ","
	nilNode   = "nil"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Serialize(root *BinaryTreeNode) string {
	if root == nil {
		return ""
	}

	result := ""
	queue := []*BinaryTreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = result + nilNode + delimiter
		} else {
			result = result + strconv.Itoa(node.Val) + delimiter
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for strings.HasSuffix(result, nilNode+delimiter) {
		result = strings.TrimSuffix(result, nilNode+delimiter)
	}
	for strings.HasSuffix(result, delimiter) {
		result = strings.TrimSuffix(result, delimiter)
	}

	return result
}

func Deserialize(s string) *BinaryTreeNode {
	if s == "" {
		return nil
	}

	nodes := []*BinaryTreeNode{}

	for _, node := range strings.Split(s, delimiter) {
		if node == nilNode {
			nodes = append(nodes, nil)
		} else {
			Value, _ := strconv.Atoi(node)
			nodes = append(nodes, &BinaryTreeNode{Val: Value, Left: nil, Right: nil})
		}
	}

	i := 1
	for _, node := range nodes {
		if node == nil {
			continue
		}

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.Left = nodes[i]
		}
		i++

		if i >= len(nodes) {
			continue
		}
		if nodes[i] != nil {
			node.Right = nodes[i]
		}
		i++
	}
	return nodes[0]
}
