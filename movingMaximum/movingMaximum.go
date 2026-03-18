package main

import (
	"fmt"
)

type MovingMaximum struct {
	size   int
	window []int
	max    int
}

func Constructor(size int) MovingMaximum {
	return MovingMaximum{size: size, window: []int{}, max: -1}
}

func (this *MovingMaximum) Next(val int) int {
	// add the new value to the monitored window and ensure the size limit is expected
	maxDeleted := false

	if this.size == len(this.window) {
		if this.max == this.window[0] {
			maxDeleted = true
		}
		this.window = this.window[1:]
	}

	this.window = append(this.window, val)

	if maxDeleted {
		this.max = this.window[0]
		for _, v := range this.window {
			if v > this.max {
				this.max = v
			}
		}
	} else if this.max < val {
		this.max = val
	}
	// return the maximum value within the window
	return this.max
}

func main() {
	m := Constructor(3)
	fmt.Println(m.Next(1))  // returns 1
	fmt.Println(m.Next(5))  // returns 5
	fmt.Println(m.Next(10)) // returns 10
	fmt.Println(m.Next(5))  // returns 10
	fmt.Println(m.Next(5))  // returns 10
	fmt.Println(m.Next(3))  // returns 5
}
