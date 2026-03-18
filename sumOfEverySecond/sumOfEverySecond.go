package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (list *LinkedList) Append(value int) {
	if list.Head == nil {
		list.Head = &ListNode{Value: value}
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Value: value}
}

func SumOfEverySecond(list *LinkedList) int {
	current := list.Head
	total := 0
	for i := 1; current != nil; i++ {
		if i%2 == 0 {
			total += current.Value
		}
		current = current.Next
	}
	return total
}

func main() {
	list := LinkedList{}
	list.Append(5)
	sum := SumOfEverySecond(&list)
	fmt.Println("Sum of Every Second Node Value:", sum)
	list.Append(10)
	list.Append(15)
	sum = SumOfEverySecond(&list)
	fmt.Println("Sum of Every Second Node Value:", sum)
}
