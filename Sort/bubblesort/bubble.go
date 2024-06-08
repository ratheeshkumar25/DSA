package main

import (
	"fmt"
	"sort"
)

// First Approach
func bubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func bubble(nums []int){
	n :=len(nums)

	for i :=0 ; i< n; i++{
		for j :=0 ; j<n-i-1; j++{
			if nums[j]> nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
}

// Second Approach
func bubbleSort1(nums []int) {
	n := len(nums)
	swap := true

	for i := 0; i < n-1 && swap; i++ {
		swap = false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
	}
}

// Sorted in Decending order
func bubbleDecending(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func distributeCandies(candyType []int) int {
	uniqueCan := make(map[int]bool)
 
	for _,candy := range candyType {
	 uniqueCan[candy] = true
	}
 
	if len(uniqueCan) > len(candyType)/2{
		 return len(candyType)/2
	}
	 return len(uniqueCan)
 }

func main() {
	arr := []int{5, 2, 1, 3, 6, 7, 4}
	bubbleSort(arr)
	fmt.Println("SORTED ARRAY AFTER THE BUBBLE SORT", arr)
	num := []int{1, 2, 1, 4, 3, 2, 7, 5, 8, 9}
	bubbleSort1(num)
	fmt.Println("Second Apporach to sort array with bubble sort", num)
	nums := []int{3, 1, 4, 1, 5}
	sort.Ints(nums)
	fmt.Println("Sorted array in ascending order", nums)
	bubbleDecending(arr)
	fmt.Println("Desecending sorted array", arr)
	
	fmt.Println("Distrubted candy",distributeCandies(arr))
}
