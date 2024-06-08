// package main

// import "fmt"

// func nextGreterElement(nums []int) []int {
// 	result := make([]int, len(nums))
// 	stack := []int{}

// 	for i := len(nums) - 1; i > 0; i-- {
// 		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
// 			stack = stack[:len(stack)-1]
// 		}
// 		if len(stack) > 0 {
// 			result[i] = nums[stack[len(stack)-1]]
// 		} else {
// 			result[i] = -1
// 		}
// 		stack = append(stack, i)
// 	}
// 	return result
// }

// func main() {
// 	arr := []int{2, 1, 2, 4, 3}
// 	result := nextGreterElement(arr)
// 	fmt.Println("Next Greater element is", result)
// }

package main

import "fmt"

type Node struct{
	Val int 
	Next *Node
}

type Stack struct{
	Top *Node
}

func (s *Stack)Push(val int){
	newNode := &Node{Val: val,Next : s.Top}
	s.Top = newNode
}

func (s *Stack)Pop()int{
	if s.Top == nil {
		return -1
	}
	value := s.Top.Val
	s.Top = s.Top.Next
	return value
}

func (s *Stack)isEmpty()bool{
	return s.Top == nil 
}

type Queue struct{
	Top  *Node
	Rear *Node
}

func (q *Queue)Enqueue(val int){
	newNode := &Node{Val: val}
	if q.Rear != nil{
		q.Rear.Next = newNode
	}
	q.Rear = newNode
	
	if q.Top == nil {
		q.Top = newNode
	}
}

func (q *Queue)Dequeue()int{
	if q.Top == nil{
		return -1
	}

	value := q.Top.Val
	q.Top = q.Top.Next

	if q.Top == nil{
		q.Rear = nil
	}
	return value
}

func (q *Queue)Display(){
	current := q.Top
	for current != nil{
		fmt.Printf(">>%d",current.Val)
		current = current.Next
	}
	fmt.Println()
}

func covertoQueue(stack Stack)*Queue{
	queue := &Queue{}
	tempStack := &Stack{}

	for !stack.isEmpty(){
		tempStack.Push(stack.Pop())
	}

	for !tempStack.isEmpty(){
		queue.Enqueue(tempStack.Pop())
	}
	return queue
}

func main(){
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	queue := covertoQueue(*stack)
	queue.Display()
	queue.Dequeue()
	queue.Display()
}