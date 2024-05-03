package main

import (
	"fmt"

)

//single linked list
type Node struct {
	data int   // Data value
	next *Node // Reference
}

type LinkedList struct {
	head *Node
}

//AddtoHead data into the single linked list

func (list *LinkedList) AddtoHead(data int) {
	newNode := &Node{}
	newNode.data = data
	
	if list.head != nil{
		newNode.next = list.head
	}
	list.head = newNode

	// if list.head == nil {
	// 	list.head = newNode
	// } else {	
	// 	current := list.head
	// 	for current.next != nil {
	// 		current = current.next
	// 	}
	// 	current.next = newNode
	// }

}

// Display the linked list 

func (list *LinkedList)Display(){
	current := list.head
	
	if current == nil{
		fmt.Println("Linked list are empty")
		return
	}
	
	fmt.Println("Linked list:")
	for current != nil{
		fmt.Printf("%d", current.data)
		current = current.next
	}
	fmt.Println()
}
func main() {
	list := LinkedList{}

	list.AddtoHead(40)
	list.AddtoHead(50)
	list.AddtoHead(60)
	list.AddtoHead(70)

	list.Display()

	
}