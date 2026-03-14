package main

import "fmt"

// AlphabetSize is the number of posible characters in the trie
const AlphabetSize = 26

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {
	wordLegth := len(w)
	currentNode := t.root
	for i := 0; i < wordLegth; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wordLegth := len(w)
	currentNode := t.root
	for i := 0; i < wordLegth; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("aragorn")
	fmt.Println(myTrie.Search("aragorn"))
}
