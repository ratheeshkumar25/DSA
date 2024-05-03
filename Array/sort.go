package main

import "fmt"

//Sort the array with loop ascending order / descendinf order 
func sortArray(arr[]int){
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if arr[i]>arr[j]{                 //for descending order if arr[i]<arr[j] ---this condition need to given
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
}

//Time Complexcity = O(n^2)--- Quardratic Time Complexcity
//Space Complexcity = O(1) --- Constant Time Complexcity

func main1(){
	arr := []int{24,15,15,22,10,33,44,55}
    sortArray(arr)
	fmt.Println("Result of Sorted Array",arr)
}