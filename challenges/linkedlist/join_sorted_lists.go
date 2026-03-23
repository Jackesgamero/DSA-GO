package linkedlist

func JoinTwoSortedLinkedLists(l1, l2 *Node) *Node {
	head := &Node{Val: 0, Next: nil}
	curr := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = &Node{Val: l1.Val}
			l1 = l1.Next
		} else {
			curr.Next = &Node{Val: l2.Val}
			l2 = l2.Next
		}
		curr = curr.Next
	}

	for l1 != nil {
		curr.Next = &Node{Val: l1.Val}
		l1 = l1.Next
		curr = curr.Next
	}

	for l2 != nil {
		curr.Next = &Node{Val: l2.Val}
		l2 = l2.Next
		curr = curr.Next
	}

	return head.Next
}
