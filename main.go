package main

import "fmt"

func Palindrome(s string) bool {
	if len(s) == 0 || len(s)==1 {
		return true
	}
	if s[0] == s[len(s)-1]{
		return Palindrome(s[1:len(s)-1])
	}
	return false	
}

func main(){
	s := "malayalam"
	r := "rat"
	fmt.Println("Number is palindrome",Palindrome(s))
	fmt.Println("Number is palindrome",Palindrome(r))
}