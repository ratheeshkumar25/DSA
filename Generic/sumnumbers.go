// package main

// import "fmt"

// //Time complexcity O(n)linear Time Complexcity, Space Complexcity O(1) constant Complexcity

// func sumNumbers[T int | float64](nums ...T) T {
// 	var sum T
// 	for _, num := range nums {//Loop iterate through each element of num in the slice of numsT
// 		sum = num + sum
// 	}
// 	return sum
// }

// func main() {
// 	num := []int{1, 2, 3, 5, 7}
// 	result := sumNumbers(num...)
// 	fmt.Println("Total Sum",result)


// 	//Two Dimensional Slices
// 	// var twoDArray [8][8]int
// 	// twoDArray [3][6]= 18
// 	// twoDArray[7][4]=3
// 	// fmt.Println(twoDArray)
	
// 	var rows int 
// 	var cols int 
// 	rows = 7
// 	cols = 9
// 	var twoslices = make([][]int,rows)
// 	for i:= range twoslices{
// 		twoslices[i] = make([]int,cols)
// 	}
// 	fmt.Println(twoslices)
// }
package main 

import"fmt"

func quickSorting(arr []int)[]int{
  //Base condition
  if len(arr)<= 1{
    return arr
  }
  pivot := arr[len(arr)/2]

  left,right := []int{},[]int{}


  for _,v := range arr {
    if v < pivot{
      left = append(left,v)
    }else if v > pivot{
      right = append(right,v)
    }
  }
  //Recursive function call
  start :=quickSorting(left)
  end :=quickSorting(right)
  return append(append(start,pivot),end...)
}

func main(){
  arr := []int{40,30,20,10,50}
  fmt.Println("Before sorting the array is ",arr)
  fmt.Println("After sorting the array is ",quickSorting(arr))
}

