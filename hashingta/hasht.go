package main

import (
	"errors"
	"fmt"
)

type Node struct {
	key  int
	val  int
	Next *Node
}

type Hashtable struct {
	table []*Node
	size  int
}

func NewHashtabel(size int) *Hashtable {
	return &Hashtable{table: make([]*Node, size), size: size}
}

func (h *Hashtable) hashKey(key int) int {
	return key % h.size
}

//Put into the hash table

func (h *Hashtable) Put(k, v int) {
	index := h.hashKey(k)
	newNode := &Node{key: k, val: v}
	if h.table[index] != nil {
		newNode.Next = h.table[index]
		h.table[index] = newNode
	} else {
		h.table[index] = newNode
	}
}

// Get the hashtable
func (h *Hashtable) Get(k int) (int,error) {
	index := h.hashKey(k)

	if h.table[index] != nil {
		curr := h.table[index]
		for curr != nil {
			if curr.key == k {
				return curr.val, nil
			}
			curr = curr.Next
		}
	}
	return 0,errors.New("not found")
}

func (h *Hashtable)Print(){
	for i,node := range h.table{
		if node != nil {
			fmt.Printf("Bucket %d:",i)
			curr := node
			for curr != nil{
				fmt.Printf("(%d,%d)",curr.key,curr.val)
				curr = curr.Next
			}
		}
	}
}

func main(){
	hash := NewHashtabel(10)
	hash.Put(1, 10)
	hash.Put(2, 20)
	hash.Put(11, 30) 
	hash.Put(21, 40) 

	hash.Print()
	
}