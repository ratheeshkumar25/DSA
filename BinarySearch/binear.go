package main

import "fmt"

//Binary search
func binarySearch(nums []int , val int)bool{
	low, high := 0, len(nums)-1
	for low <= high{
		mid := (low+high)/2
		if  val < nums[mid] {
			low = mid+1
		}else if val > nums[mid]{
			high = mid-1
		}else{
			return true
		}
	}
	return false
}

//Binary Recursion 

//First Occurance 
func firstOccurrence(nums []int , target int )int{
	low,high := 0,len(nums)-1
	firstOccur := -1

	for low <= high{
		mid := (low+high)/2
		if nums[mid] == target{
			firstOccur = mid
			high = mid-1
		}else if nums[mid]<target{
			low = mid+1
		}else{
			high = mid-1
		}
	}
	return firstOccur
}
//Last Occurance 
func lastOccurrence(nums []int , val int )int{
	low,high := 0,len(nums)-1
	lastOccur := -1
	for low <= high{
		mid := (low+high)/2
		if nums[mid]== val{
			lastOccur = mid
			low = mid+1
		} else if val < nums[mid]{
			high = mid-1
			
		}else{
			low = mid+1
		}
	}
	return lastOccur
}

func main(){
	arr := []int{1,2,9,10,10,10,11,45,73}
	tar := 10
	index := firstOccurrence(arr,tar)
	ind := lastOccurrence(arr,tar)
	fmt.Println(binarySearch(arr,44))
	fmt.Println(binarySearch(arr,10))
	fmt.Println("First Occurance",index)
	fmt.Println("Last Occurance",ind)
}