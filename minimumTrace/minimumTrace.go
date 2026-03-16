package main

import (
	"fmt"
)

type MaxStack struct {
	stack     []int
	maxValues []int
}

func (m *MaxStack) Push(x int) {
	// add element to the stack and track the max value
	if len(m.maxValues) == 0 || m.maxValues[len(m.maxValues)-1] <= x {
		m.maxValues = append(m.maxValues, x)
	}
	m.stack = append(m.stack, x)
}

func (m *MaxStack) Pop() {
	// remove an element from the stack and account for when removing the max element
	n := len(m.stack)
	if n > 0 {
		if m.stack[n-1] == m.maxValues[len(m.maxValues)-1] {
			m.maxValues = m.maxValues[:len(m.maxValues)-1]
		}
		m.stack = m.stack[:n-1]
	}
}

func (m *MaxStack) Top() int {
	if len(m.stack) == 0 {
		return -1
	}
	return m.stack[len(m.stack)-1]
}

func (m *MaxStack) GetMax() int {
	// TODO: retrieve the maximum value in the stack
	if len(m.stack) > 0 {
		return m.maxValues[len(m.maxValues)-1]
	}
	return -1
}

func main() {
	maxStack := &MaxStack{}
	maxStack.Push(-2)
	maxStack.Push(0)
	maxStack.Push(-3)
	fmt.Println(maxStack.GetMax()) // Expected Output: 0
	maxStack.Pop()
	fmt.Println(maxStack.Top())    // Expected Output: 0
	fmt.Println(maxStack.GetMax()) // Expected Output: 0

	maxStack.Push(4)
	fmt.Println(maxStack.GetMax()) // Expected Output: 4
	maxStack.Push(2)
	fmt.Println(maxStack.GetMax()) // Expected Output: 4
}
