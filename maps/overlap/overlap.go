package main

import "fmt"

func IsAttendeeOverlap(event1, event2 []int) bool {
	attendeeMap := make(map[int]struct{})

	// TODO: check attendee list in both events to determine if an overlap is present using maps in go
	for _, person := range event1 {
		attendeeMap[person] = struct{}{}
	}

	for _, person := range event2 {
		if _, result := attendeeMap[person]; result {
			return true
		}
	}

	return false
}

func main() {
	event1 := []int{1, 2, 3, 4, 5}
	event2 := []int{6, 7, 8, 9, 10}

	if IsAttendeeOverlap(event1, event2) {
		fmt.Println("Yes, there is an attendee overlap between both concerts.")
	} else {
		fmt.Println("No, there is no attendee overlap between both concerts.")
	}
}
