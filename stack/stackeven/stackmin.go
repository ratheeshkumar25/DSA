package main

import (
	"fmt"
	//"math"
)

type Node struct {
	Value int
	Next  *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{Value: val,Next: s.top}
	s.top = newNode
}

func (s *Stack)Min()int{
	// if s.top == nil {
	// 	return math.MaxInt64
	// }

	minval := s.top.Value
	current := s.top.Next
	for current != nil {
		if current.Value < minval{
			minval = current.Value
		}
		current = current.Next

	}
	return minval
}

func (s *Stack)Max()int{

	maxValue := s.top.Value
	current := s.top.Next

	for current != nil {
		if current.Value > maxValue {
			maxValue = current.Value
		}
		current = current.Next
	}
	return maxValue
}

func (s* Stack)Middle()int{
	if s.top == nil || s.top.Next == nil{
		return s.top.Value
	}

	slow,fast := s.top, s.top 

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Value
}
func (s Stack) Display() {
	current := s.top
		fmt.Println("stack")
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func main(){
	stk := &Stack{}
	stk.Push(10)
	stk.Push(30)
	stk.Push(35)
	stk.Push(40)
	stk.Push(50)
	stk.Push(20)

	stk.Display()

	min := stk.Min()

	fmt.Println("Min value of stack",min)

	max := stk.Max()

	fmt.Println("Max value of stack",max)

	middle := stk.Middle()
	fmt.Println("Middle value of stack",middle)
}