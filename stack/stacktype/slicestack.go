// package main

// import "fmt"

// type Stack struct {
// 	item []int
// }

// // Push to stack
// func (s *Stack) Push(data int) {
// 	s.item = append(s.item, data)
// }

// // Pop
// func (s *Stack) Pop() int {
// 	if len(s.item) == 0 {
// 		return -1
// 	}

// 	items := s.item[len(s.item)-1]
// 	s.item = s.item[:len(s.item)-1]
// 	return items
// }

// func main() {
// 	stack := &Stack{}
// 	stack.Push(50)
// 	stack.Push(60)
// 	stack.Push(70)
// 	stack.Push(80)
// 	stack.Pop()
// 	fmt.Println("Stack after pushing items",stack.item)
// }

package main

import "fmt"


type Node struct{
	key int
	next *Node
}

type stack struct{
	top *Node
}

func (s *stack)Push(k int){
	 newNode := &Node{key:k}

	 if s.top == nil{
		s.top = newNode
	 }else{
		newNode.next = s.top
		s.top = newNode
	 }
}

func (s *stack)Pop()int {
	if s.top == nil {
	return-1
	}
	val := s.top.key
	s.top = s.top.next
	return val 
}

func (s *stack)Display(){
	  if s.top == nil{
		fmt.Println("Empty stack")
	  }else{
		current := s.top
		fmt.Println("Stack is")
		for current != nil {
			fmt.Printf("->%d",current.key)
			current = current.next
		}
	  }
}

func (s *stack)Removemiddle()int{
		if s.top == nil || s.top.next == nil{
			return -1
		}

		var prev *Node
		slow := s.top
		fast := s.top 

		for fast != nil && fast.next != nil{
			prev = slow
			slow = slow.next
			fast = fast.next.next
		}

		if prev == nil {
			s.top = slow.next
		}else{
			prev.next = slow.next
		}
		return slow.key

	}

	func main(){
		stack := &stack{}

		stack.Push(20)
		stack.Push(40)
		stack.Push(30)
		stack.Push(50)
		stack.Push(60)

		stack.Display()

		middle := stack.Removemiddle()
		fmt.Println("Removed",middle)
		stack.Display()
	}