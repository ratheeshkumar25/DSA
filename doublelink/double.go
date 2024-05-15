package main 

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doublylinkedlist struct {
	head   *node
	tail   *node
	length int
}

//Add the list in double linked list 
func (l *doublylinkedlist) Add(data int) {
	newNode := &node{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	}else{
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.length++
}

func (l *doublylinkedlist) Display() {
	current := l.head

	// fmt.Printf("Linked list")
	for current != nil {
		fmt.Printf("%d ->",current.data)
		current = current.next
	}
	fmt.Println()
}

//Middle element of the doublelinked list
func Middle(head *node)*node{
	if head == nil{
		return nil
	}
	slow := head
	fast := slow.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func main(){
	list := &doublylinkedlist{}
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)
	list.Add(60)
	list.Add(70)
	list.Add(80)
	list.Display()
	m := Middle(list.head)
	fmt.Println("midlle of double link list",m.data)
	
}