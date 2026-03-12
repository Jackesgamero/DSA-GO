package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Printf("%d ", curr.data)
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.head == nil {
		return
	}

	for l.head != nil && l.head.data == value {
		l.head = l.head.next
		l.length--
	}

	prev := l.head
	for prev != nil && prev.next != nil {
		if prev.next.data == value {
			prev.next = prev.next.next
			l.length--
		} else {
			prev = prev.next
		}
	}
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 48}
	node5 := &node{data: 16}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.deleteWithValue(48)

	mylist.printListData()
	//fmt.Println(mylist)
}
