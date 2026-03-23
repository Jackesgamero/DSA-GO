package linkedlist

func ReverseLinkedList(head *Node) *Node {
	var last *Node

	for head != nil {
		next := head.Next
		head.Next = last
		last = head
		head = next
	}

	return last
}
