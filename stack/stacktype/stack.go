// package main

// import "fmt"

// type Node struct {
// 	data interface{}
// 	next *Node
// }

// type stack struct {
// 	top *Node
// }

// // Push to stack
// func (s *stack) Push(data interface{}) {
// 	newNode := &Node{data: data}
// 	if s.top == nil {
// 		s.top = newNode
// 	} else {
// 		newNode.next = s.top
// 		s.top = newNode
// 	}

// }

// //Pop operation
// func (s *stack) Pop(){
// 	if s.top == nil {
// 		fmt.Println("Empty stack")
		
// 	}else{
// 		s.top = s.top.next
// 	}	
// }

// //Peek
// func (s*stack)Peek()any{
// 	if s.top == nil{
// 		return -1
// 	}
// 	return s.top.data
// }

// //Display stack 
// func (s*stack)Display(){
// 	if s.top == nil{
// 		fmt.Println("Emplty stack")
// 	}else{
// 		current := s.top
// 		for current != nil{
// 			fmt.Println(current.data)
// 			current = current.next
// 		}
// 	}
// }

// func main(){
// 	stk := &stack{}
// 	stk.Push(20)
// 	stk.Push(30)
// 	stk.Push(40)
// 	stk.Push(50)
// 	stk.Display()

// 	stk.Pop()
// 	stk.Display()
// }