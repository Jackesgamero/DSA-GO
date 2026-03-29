package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solution(logs string) []string {
	// TODO: Your solution
	events := strings.Split(logs, ", ")

	borrowTime := make(map[string]int)
	totalTime := make(map[string]int)

	for _, event := range events {
		data := strings.Split(event, " ")
		id := data[0]
		action := data[1]
		time := parseTime(data[2])

		if action == "borrow" {
			borrowTime[id] = time
		} else {
			start := borrowTime[id]
			totalTime[id] += time - start
			delete(borrowTime, id)
		}
	}

	maxTime := 0
	for _, t := range totalTime {
		if t > maxTime {
			maxTime = t
		}
	}

	var result []string
	for id, t := range totalTime {
		if t == maxTime {
			result = append(result, id)
		}
	}

	sort.Strings(result)

	for i, id := range result {
		h := maxTime / 60
		m := maxTime % 60
		result[i] = fmt.Sprintf("%s %02d:%02d", id, h, m)
	}

	return result
}

func parseTime(t string) int {
	parts := strings.Split(t, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h*60 + m
}

func main() {
	logs := "1 borrow 09:00, 2 borrow 10:00, 1 return 12:00, 3 borrow 13:00, 2 return 15:00, 3 return 16:00"
	result := Solution(logs)
	for _, entry := range result {
		fmt.Println("Book " + strings.Split(entry, " ")[0] + " duration: " + strings.Split(entry, " ")[1])
	}
	// Expected output: "Book 2 duration: 05:00"
}
