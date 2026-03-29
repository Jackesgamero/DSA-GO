package main

import (
	"fmt"
	"sort"
	"strings"
)

func OrganizeInbox(inboxString string) []string {
	// TODO: Your solution
	mails := strings.Split(inboxString, "; ")
	emails := make(map[string]int)

	for _, mail := range mails {
		content := strings.Split(mail, ", ")
		emails[content[0]]++
	}

	type pair struct {
		email string
		count int
	}

	pairs := []pair{}
	for k, v := range emails {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].email < pairs[j].email
		}
		return pairs[i].count > pairs[j].count
	})

	result := []string{}
	for _, p := range pairs {
		result = append(result, fmt.Sprintf("%s %d", p.email, p.count))
	}
	return result
}

func main() {
	inboxString := "sender1@example.com, Subject1, 09:00; sender2@example.com, Subject2, 10:00; sender1@example.com, Subject3, 12:00"

	// Expected output: ["sender1@example.com 2", "sender2@example.com 1"]
	result := OrganizeInbox(inboxString)
	for _, entry := range result {
		fmt.Println(entry)
	}
}
