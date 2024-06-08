package main

import "fmt"

type ListNode struct {
	val  byte
	next *ListNode
}

// type List struct {
// 	head *ListNode
// 	tail *ListNode
// }

func stringToList(s string) *ListNode {
	if len(s) == 0 {
		return nil
	}

	head := &ListNode{val: s[0]}
	current := head
	for i := 1; i < len(s); i++ {
		current.next = &ListNode{val: s[i]}
		current = current.next
	}
	return head
}

func ReversList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}



func listToString(head *ListNode) string {
	result := ""
	for head != nil {
		result += string(head.val)
		head = head.next
	}
	return result
}

func main() {
	s := "hello"
	head := stringToList(s)
	fmt.Println("Orginal list", listToString(head))
	reverseHead := ReversList(head)
	fmt.Println("Reverse list",listToString(reverseHead))
}