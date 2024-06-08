package main

import "fmt"

func mergeSort(arr []int) []int {
	//Base condition
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// func merge(left , right []int)[]int{
// 	result := make([]int,len(left)+len(right))
// 	l, r,s := 0,0,0

// 	for l < len(left) && r < len(right){
// 		if left[l] < right[r]{
// 			result[s] = left[l]
// 			l++
// 		}else{
// 			result[s] = right[r]
// 			r++
// 		}
// 		s++
// 	}

// 	for l < len(left){
// 		result[s] = left[l]
// 		l++
// 		s++
// 	}
// 	for r < len(right){
// 		result[s] = right[r]
// 		r++
// 		s++
// 	}

// 	return result
// }

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	arr1 := []int{12, 11, 13, 5, 6, 7}

	arr2 := append(arr,arr1...)
	fmt.Println("Original array",arr2)

	sorted := mergeSort(arr2)
	fmt.Println("Sorted array",sorted)
}