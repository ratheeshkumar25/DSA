 package main

import "fmt"

//Linear Search 
func LinearSearch(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}

func LinearSearc(nums []int,target int)int{ // Time complexcity O(n), O(1) is space complexcity
	for i,item := range nums{
		if item == target{
			return i+1
		}
	}
	return 0
}

//Multiple occurrence 

func multiplEOccurrence(nums []int, target int)[]int{
	var occurrence []int 
	for i,num := range nums{
		if num == target{
			occurrence = append(occurrence, i)
		}
	}
	return occurrence
}

func main() {
	arr := []int{20, 30, 40, 30,50,30, 60, 70, 80, 90}
	val := 30
	target := LinearSearc(arr,30)
	result := multiplEOccurrence(arr,val)
	fmt.Println("The source of array",target)
	fmt.Println("Search item found",LinearSearch(arr,30))
	//mutliple occurrence 
	fmt.Println("Multiple occurrence",result,target)
	//if item not found this session need to execute 
	// if len(result)> 0{
	// 	fmt.Println("Mutliple occurence found",result,val)
	// }else{
	// 	fmt.Println("Multiple occurence not found",val)
	// }

}
