package main

import (
	"fmt"
	//"math"
)

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{Val: val, Next: s.Top}
	s.Top = newNode
}

func (s *Stack)Min()int{
	if s.Top == nil {
		return -1
	}

	minVal :=  s.Top.Val
	current := s.Top.Next

	for current != nil {
		if current.Val < minVal{
			minVal = current.Val
			current = current.Next
		}
	}
	return minVal
}

func (s *Stack) Display() {
	if s.Top == nil {
		fmt.Println("Stack is empty")
	}else{
		current := s.Top
		for current != nil{
			fmt.Println(current.Val)
			current = current.Next
		}
	}
}

func (s *Stack)Pop()int{
	if s.Top == nil {
		return -1
	}
	value := s.Top.Val
	s.Top = s.Top.Next
	return value
}

func (s *Stack)Middle()int{
	if s.Top == nil || s.Top.Next == nil {
		return s.Top.Val
	}

	slow := s.Top
	fast := s.Top.Next

	for fast != nil && fast.Next != nil {
		slow = s.Top.Next
		fast = s.Top.Next.Next
	}
	return slow.Val
}

func (s *Stack) deleteMiddle(){
	if s.Top == nil || s.Top.Next == nil{
		return
	}

		slow := s.Top
		fast := s.Top.Next
		prev := slow 

	for fast != nil && fast.Next !=nil {
		prev = slow 
		slow = s.Top.Next
		fast = s.Top.Next.Next	
	}
		prev.Next = slow.Next

		if fast != nil{
			slow.Next = nil
		}
}



func main(){
	stk := &Stack{}
	stk.Push(1)
	stk.Push(7)
	stk.Push(4)
	stk.Push(8)
	stk.Push(11)

	// prime := stk.findPrime()
	// fmt.Println("Prime number is ",prime)
	middle := stk.Middle()
	fmt.Println("Middle amount of",middle)
	stk.Display()
	min :=stk.Min()
	fmt.Println("Min value of stack",min)
	stk.deleteMiddle()
	stk.Display()

	
}