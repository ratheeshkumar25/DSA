package main

import (
	"fmt"
	
)

//time complexcity(o(n2)), space complexcity(o(1))
// func twoSum(nums []int, target int)[]int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if target == nums[i]+nums[j] {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func twoSumHashTable(nums[]int,target int)[]int{
// 	seen := make(map[int]int)
// 	for i,num := range nums{
// 		complement := target-num
// 		if index,ok := seen[complement];ok{
// 			return []int{index,i}
// 		}
// 		seen[num]= i
// 	}
// 	return []int{}
// }

//Move to last the target
func movetoEnd(nums []int, target int)  {
	var i = 0
	var j = len(nums) - 1
    // Move all non target elements to the front
	for i < j {
		if nums[i] != target && nums[j] == target {
			i++
			j--
			continue
		}
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}
	// move all target element to the end 
	j = len(nums)-1
	for i < j {
		if nums[i] == target{
			nums[i],nums[j] = nums[j],nums[i]
			j--
		}else{
			i++
		}
	}

}

func plusOne(digits []int) []int {
   	// Start from the end of the array
	for i := len(digits) - 1; i >= 0; i-- {
		// Increment the current digit
		digits[i]++
		// If the digit becomes 10, set it to 0 and continue to the next digit
		if digits[i] < 10 {
			return digits
		}
		// Otherwise, carry the one to the next digit
		digits[i] = 0
	}
	// If we reached here, it means all digits were 9, so add one more digit to the front
	return append([]int{1}, digits...)
	
    
}

//Reverse the array

func main() {

	arr := []int{1, 2, 3, 3, 5, 3, 5, 6, 3, 9}
	tar := 5
	movetoEnd(arr, tar)
	fmt.Println("After moving to the last ", arr)
	arr1 := []int {4,4,4,5}
	plusOne(arr1)
	fmt.Println("Plusone",arr1)





}
