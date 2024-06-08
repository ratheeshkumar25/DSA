// package main

// import "fmt"

// type Node struct {
// 	data int
// 	Next *Node
// }

// type circularlist struct {
// 	head *Node
// 	tail *Node
// }
// //Add data to the node
// func (c *circularlist) Add(data int) {
// 	newNode := &Node{data: data, Next: nil}

// 	if c.head == nil {
// 		c.head = newNode
// 		c.tail = newNode
// 		newNode.Next = c.head
// 	} else {
// 		c.tail.Next = newNode
// 		c.tail = newNode
// 		newNode.Next = c.head
// 	}
// }

// //Display the circular list
// func (c *circularlist) Travese() {
// 	if c.head == nil {
// 		return
// 	}
// 	current := c.head
// 	fmt.Println("linked list")
// 	for current != nil {
// 		fmt.Printf("%d ->",current.data)
// 		current = current.Next
// 		if current == c.head{
// 			break
// 		}
// 	}
// }

// //Check it linklist circular linklist
// func (c *circularlist)isCircular()bool{
// 	if c.head == nil {
// 		return false
// 	}
// 	slow := c.head
// 	fast := c.head

// 	for fast !=nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next

// 		if slow == fast{
// 			return true
// 		}
// 	}
// 	return false
// }

// func main(){
// 	list := &circularlist{}
// 	list.Add(20)
// 	list.Add(30)
// 	list.Add(40)
// 	list.Add(50)
// 	list.Travese()

// 	if list.isCircular(){
// 		fmt.Println("The linklist are circular list")
// 	}else {
// 		fmt.Println("The link list is not circular list")
// 	}
// }

	package main

	import "fmt"

	type Node struct {
		Val int
		Next *Node
	}

	type LinkedList struct{
		head *Node
	}

	func (l * LinkedList)Add(data int){
		newNode := &Node{Val:data}

		if l.head == nil{
			l.head = newNode
		}else{
			current := l.head
			for current != nil {
				current = current.Next
			}
			current.Next = newNode
		}
	}

	func Middle(head *Node)*Node{
		if head == nil || head.Next == nil {
			return head
		}
		slow,fast := head , head.Next

		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	func mergeTwoList(l1,l2 *Node) *Node{
		dummy := &Node{}
		tail := dummy 

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val{
				tail.Next = l1
				l1 = l1.Next
			}else{
				tail.Next = l2
				l2 = l2.Next
			}
			tail = tail.Next
		}

		if l1 != nil {
			tail.Next = l1
		}else{
			tail.Next = l2
		}

		return dummy.Next
	}

	func mergeSort(head *Node) *Node{
		if head == nil || head.Next == nil {
			return head
		}

		middel := Middle(head)
		nextMiddle := middel.Next
		middel.Next = nil 

		left := mergeSort(head)
		right := mergeSort(nextMiddle)

		sorted := mergeTwoList(left,right)
		return sorted 
	}

	func (l *LinkedList)Display(){
		current := l.head

		for current != nil {
			fmt.Printf("Linked list -%d >",current.Val)
			current = current.Next
		}
	}

	func main(){
		link := &LinkedList{}

		link.Add(20)
		link.Add(40)
		link.Add(30)
		link.Add(10)

		link.Display()
		
		link.head = mergeSort(link.head)

		fmt.Println("Sorted list:")
		link.Display()
	}


