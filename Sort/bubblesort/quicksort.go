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