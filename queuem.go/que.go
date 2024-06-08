// package main

// import "fmt"

// type Node struct {
// 	Val  int
// 	Next *Node
// }

// type Queue struct {
// 	Front *Node
// 	Rear  *Node
// }

// //Enqueue in Queue
// func (q *Queue)Enqueue(val int){
// 	newNode := &Node{Val: val}
// 	if q.Front == nil{
// 		q.Rear = newNode
// 		q.Front = newNode
// 		return
// 	}
// 	q.Rear.Next = newNode
// 	q.Rear = newNode
// }

// //Dequeue the in the

// func (q *Queue)Dequeue()*Node{
// 	node := q.Front
// 	q.Front = q.Front.Next
// 	return node
// }

// //Display Queue

// func (q *Queue)Display(){
// 	current := q.Front
// 	for current != nil {
// 		fmt.Printf(">>%d",current.Val)
// 		current = current.Next
// 	}

// }

// func main(){
// 	que := &Queue{}
// 	que.Enqueue(20)
// 	que.Enqueue(30)
// 	que.Enqueue(40)
// 	que.Enqueue(60)
// 	que.Enqueue(50)
// 	que.Enqueue(70)

// 	que.Dequeue()

// 	que.Display()
// }

package main

import "fmt" 


type Node struct{
	Val rune
	Next *Node
}
type Stack struct{
	Top *Node
}

func (s *Stack)Push(val rune){
	newNode := &Node{Val :val,Next : s.Top}
	s.Top = newNode
}

func (s *Stack)Pop()rune{
	if s.IsEmpty(){
		return 0
	}
	value := s.Top.Val
	s.Top = s.Top.Next
	return value
}

func (s*Stack)IsEmpty()bool{
	return s.Top == nil
}

func ReverseString(s string)string{
	stack := &Stack{}

	//Push the all charcter to stack 
	for _,char := range s{
		stack.Push(char)
	}

	//Pop the all charcter from the stack 

	reversed := make([]rune,0,len(s))

	for !stack.IsEmpty(){
		reversed = append(reversed, stack.Pop())
	}
	return string(reversed)
}


func main(){
	str := "Hello World"
	reversed := ReverseString(str)
	fmt.Println("Orginal string",str)
	fmt.Println("Reveised string",reversed)
}
