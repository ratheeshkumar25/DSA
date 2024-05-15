package main

import (
	"fmt"
	
)

//Factorial using recursive function
func factorial (n int)int{
	if n<= 1{
		return 1
	}
	return n*factorial(n-1)
}

//Recursion add integer number until 5 times 
func Add(n int)int{
	if n == 5 {
		//fmt.Println(n)
		return n
	}
	return Add(n+1)
}

// Fibanci 
func Fib(n int)int{
	//Base condition 

	if n <2{
		return n
	}
	return Fib(n-1)+Fib(n-2)// Recursive relationship 
}

//Reverse the array with recursive function 
func ReverseArry(arr[]int,start int , end int)[]int{
	//Base conditions
	if start >= end{
		return arr
	}
	arr[start],arr[end] = arr[end],arr[start]
	//Recursive function 
	return ReverseArry(arr,start+1,end)
}

//Binary Search Using Recurive function
func BinaryRecursion(nums[]int,target,low,high int)int{
	mid := (low+high)/2
	if low <=high{
		if nums[mid] == target{
			return mid
		}else if target < nums[mid]{
			BinaryRecursion(nums,target,low,mid-1)
		}else{
			BinaryRecursion(nums,target,low,mid+1)
		}
	}
	return mid
}
//Sum of array
func RecursiveSum(arr[]int)int{
	if len(arr) == 0{
	  return 0
	}
	return arr[0] + RecursiveSum(arr[1:])
  }

func IsPlalindrome(s string)bool{
	//Base Condition
	if len(s)<=1{
		return true
	}

	//Body 
	if s[0] != s[len(s)-1]{
		return false
	}
	//Recursive call 
	return IsPlalindrome(s[1:len(s)-1])
}

func ReverseString(s string)string{
	//Base conditions
	if len(s) <= 1{
		return s 
	}
	return ReverseString(s[1:] + string(s[0]))
}


func main(){

	m := 5
	arr := []int{25,30,40,50,60,70,80,90,100}
	val := 60
	low, high := 0, len(arr)-1
	fmt.Println("Items found",BinaryRecursion(arr,val,low,high))
	
	fmt.Println(factorial(m))
	fmt.Println("Total of interger after adding ",Add(1))
	fmt.Println("Fibanaci number is ",Fib(4))
	fmt.Println("After reverse the array is",ReverseArry(arr,0,len(arr)-1))
	sum := RecursiveSum(arr)
	fmt.Println("Sum of array is",sum)
	h := "hello"
	fmt.Println(IsPlalindrome(h))

	str := "string"
	fmt.Println("Reverse string",ReverseString(str))
}