package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList(value int) *LinkedList {
	return &LinkedList{Head: &ListNode{Value: value}}
}

func (ll *LinkedList) Append(value int) {
	newNode := &ListNode{Value: value}
	current := ll.Head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) RemoveDuplicatesFromStart() {
	if ll.Head == nil {
		return
	}

	count := make(map[int]int)
	current := ll.Head

	// store the frequency of values in the linked list
	for current != nil {
		count[current.Value]++
		current = current.Next
	}

	// remove duplicates that appear earlier in the linked list
	var prev *ListNode
	current = ll.Head

	for current != nil {
		if count[current.Value] > 1 {
			count[current.Value]--
			if prev == nil {
				ll.Head = current.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}
		current = current.Next
	}

}

func (ll *LinkedList) PrintList() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func main() {
	linkedList := NewLinkedList(1)
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(1)

	linkedList.RemoveDuplicatesFromStart()
	linkedList.PrintList()
}
