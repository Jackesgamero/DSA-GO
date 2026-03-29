package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func AnalyzeCompetition(logs string) [][]int {
	// TODO: Your solution
	type stats struct {
		score   int
		success int
		fails   int
	}
	students := make(map[string]stats)

	events := strings.Split(logs, ", ")
	for _, event := range events {
		data := strings.Split(event, " ")
		id := data[0]
		stat := students[id]
		if data[1] == "fail" {
			stat.fails++
		} else {
			score, _ := strconv.Atoi(data[2])
			stat.score += score
			stat.success++
		}
		students[id] = stat
	}

	results := [][]int{}
	for key, val := range students {
		if val.success == 0 {
			continue
		}
		id, _ := strconv.Atoi(key)
		results = append(results, []int{id, val.score, val.success, val.fails})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i][1] > results[j][1]
	})

	return results
}

func main() {
	logs := "1 solve 50, 2 solve 60, 1 fail, 3 solve 40, 2 fail, 3 solve 70"

	// Expected output: [[3, 110, 2, 0], [2, 60, 1, 1], [1, 50, 1, 1]]
	result := AnalyzeCompetition(logs)
	for _, entry := range result {
		fmt.Printf("[%d, %d, %d, %d]\n", entry[0], entry[1], entry[2], entry[3])
	}
}
