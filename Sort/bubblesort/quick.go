package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1

	
}

func part(arr []int,low,high int)int{
	pivot := arr[high]
	i := low-1

	for j:=low; j < high; j++{
		if arr[j]<pivot{
			i++
			arr[i],arr[j] = arr[j],arr[j]
		}
	}
	arr[i+1],arr[high] = arr[high],arr[i+1]
	return i+1
}

func quick(arr[]int, low,high int){
	if low < high{
		piv := part(arr,low,high)
		quick(arr,low,piv-1)
		quick(arr,piv+1,high)
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array is", arr)
}