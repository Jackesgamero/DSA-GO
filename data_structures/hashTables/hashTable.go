package main

import "fmt"

const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// bucket node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is
// stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a key, create a node with the key
// and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search will take in a key and return true if the bucket
// if the bucket has that key
func (b *bucket) search(k string) bool {
	for i := b.head; i != nil; i = i.next {
		if i.key == k {
			return true
		}
	}
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string) {
	if b.head == nil {
		return
	}

	for b.head != nil && b.head.key == k {
		b.head = b.head.next
	}

	prev := b.head
	for prev != nil && prev.next != nil {
		if prev.next.key == k {
			prev.next = prev.next.next
		} else {
			prev = prev.next
		}
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range ArraySize {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
}
