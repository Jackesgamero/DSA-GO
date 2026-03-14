package main

import "fmt"

// Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key > k {
		if n.Left != nil {
			n.Left.Insert(k)
		} else {
			n.Left = &Node{Key: k}
		}
	} else {
		if n.Right != nil {
			n.Right.Insert(k)
		} else {
			n.Right = &Node{Key: k}
		}
	}
}

// Search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key > k {
		return n.Left.Search(k)
	} else if n.Key < k {
		return n.Right.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(20)

	fmt.Println(tree.Search(300))
}
