package main

import "fmt"

func main() {
	// TODO: Create a map to store unique user IDs
	set := make(map[string]struct{})
	// TODO: Add two different user IDs to your map and try adding one of them again
	set["User1"] = struct{}{}
	set["User2"] = struct{}{}
	set["User1"] = struct{}{}

	// TODO: Display the number of unique users
	fmt.Print(len(set))
}
