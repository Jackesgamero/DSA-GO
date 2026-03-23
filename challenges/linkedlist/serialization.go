package linkedlist

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

const separator = `->`

func Serialize(node *Node) string {
	var rep []string

	for node != nil {
		rep = append(rep, strconv.Itoa(node.Val))
		node = node.Next
	}

	return strings.Join(rep, separator)
}

func Deserialize(stringRepresentation string) *Node {
	if stringRepresentation == "" {
		return nil
	}

	parts := strings.Split(stringRepresentation, separator)
	node := &Node{Val: atoi(parts[0])}

	curr := node
	for i := 1; i < len(parts); i++ {
		curr.Next = &Node{Val: atoi(parts[i])}
		curr = curr.Next
	}
	return node
}

func atoi(number string) int {
	i, err := strconv.Atoi(number)
	if err != nil {
		return -1
	}
	return i
}
