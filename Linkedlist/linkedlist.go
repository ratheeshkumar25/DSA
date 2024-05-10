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

func (link *LinkedList) AddtoHead(data int){
	newNode := &Node{data:data}

	if link.head == nil {
		link.head = newNode
	}else{
		current := link.head
		for current.next != nil{
			current = current.next
		}
		current.next = newNode
	}

}
	
// Display the linked list 

func (list *LinkedList)Display(){
	current := list.head

	fmt.Println("Linked List")
	for current != nil{
		fmt.Println(current.data,"")
		current = current.next
	}
	fmt.Println()
	
	// if current == nil{
	// 	fmt.Println("Linked list are empty")
	// 	return
	// }
	
	// fmt.Println("Linked list:")
	// for current != nil{
	// 	fmt.Printf("%d", current.data)
	// 	current = current.next
	// }
	// fmt.Println()
}
//link list find the middle of the node
func middle(head *Node)*Node{
	 if head == nil{
		return nil
	 }
	 slow := head
	 fast := slow.next

	 for fast !=nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	 }
	 return slow
}

// Reverse the linked list 
func (l1 *LinkedList)Reverese(){
	var prev *Node
	current := l1.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
		
	}
	l1.head = prev
}

//Linked list to array 

func linkedListArray(head *Node)[]int{
	//find the length of array 
	len := 0
	
	for curr := head; curr != nil ; curr = curr.next {
		len++
	}

	arr := make([]int,len)

	i := 0
	for curr := head; curr != nil ; curr = curr.next{
		arr[i] = curr.data
		i++
	}
	return arr
}

//Delete the middel of node 
func (l *LinkedList)DeleteMiddle(){
	if l.head == nil || l.head.next == nil{
		return
	}
	slow := l.head
	fast := l.head 
	prevslow := slow 

	for fast != nil && fast.next != nil {
		prevslow = slow 
		slow = slow.next
		fast = fast.next.next
	}
	prevslow.next = slow.next

	if fast != nil{
		slow.next = nil
	}

}

func main() {
	list := LinkedList{}

	list.AddtoHead(40)
	list.AddtoHead(50)
	list.AddtoHead(60)
	list.AddtoHead(70)
	list.AddtoHead(80)

	middle := middle(list.head)
	
	fmt.Println("Middle of the list",middle.data)
	list.Display()

	list.Reverese()
	fmt.Println("Revised Link List")
	list.Display()

	list.DeleteMiddle()
	fmt.Println("Deleted the middle of list")
	list.Display()

	index := linkedListArray(list.head)
	fmt.Println("After reversing link list to array",index)
	

	
}