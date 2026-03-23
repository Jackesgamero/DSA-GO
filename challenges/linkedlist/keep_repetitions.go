package linkedlist

func KeepRepetitions(head *Node) *Node {
	new := &Node{Val: 0, Next: nil}
	num := make(map[int]bool)
	curr := new

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val && !num[head.Val] {
			curr.Next = &Node{Val: head.Val}
			curr = curr.Next
			num[head.Val] = true
		}

		head = head.Next
	}
	return new.Next
}
