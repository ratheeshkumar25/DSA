package main

import "fmt"
// Ascending order 
func selectionSort(num []int) {
	n := len(num)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if num[j] < num[minIndex] {
				minIndex = j
			}
		}
		num[i], num[minIndex] = num[minIndex], num[i]
	}
}
//	Descending Order
func selectionSortDes(num []int){
	n := len(num)
	for i :=0 ; i <n-1 ; i++{
		minIndex := i
		for j := i+1; j<n;j++{
			if num[j]> num[minIndex]{
				minIndex = j
			}
		}
		num[i],num[minIndex] = num[minIndex],num[i]
	}
}

// Sort from the particular elements 
func selctionPar(arr[]int,k int){
	n := len(arr)

	if k > n {
		n = k 
	}
   // Compaare the given the value until 
	for i :=0 ; i< k-1; i++{
		min := i
		for j :=i+1; j <k ; j++{
			if arr[j]< arr[min]{
				min = j
			}
		}
		if i != arr[min]{
			arr[i],arr[min] = arr[min],arr[i]
		}
	}
}

//Swap function 
func swap(arr[]int, first int, second int){
	temp := arr[first]
	arr[first] = arr[second]
	arr[second]= temp
}
//find the max index 
func getMaxindex(arr[]int, start int,end int)int{
	max := start

	for i:=start; i<=end;i++{
		if arr[max]<arr[i]{
			max = i
		}	
	}
	return max
}

func selctionSorting(arr []int){
	n:= len(arr)	
	for i:=0; i<n;i++{
		start :=0 
		last := n-1-i	
		maxIndex := getMaxindex(arr,start,last)
		swap(arr,maxIndex,last)
	}
}

//Unsorted array 
func InsertSort(arr []int){

	for i:=1; i<len(arr); i++{
		key := arr[i]
		j := i-1
		for j>=0 && arr[j]>key{
			arr[j+1]= arr[j]
			j--
		}
		arr[j+1] = key 
	}
}

//Insert in particular point of sorted array 

func InsertSorting(arr []int , targ int)[]int{
	arr = append(arr,0)
	for i:= len(arr)-2; i>= 0 ; i--{
		if arr[i]> targ{
			arr[i+1] = arr[i]
		}else {
			arr[i+1] = targ
			return arr
		}
	}
	arr[0] = targ
	return arr
}

func main() {
	arr := []int{8, 4, 7, 6, 5, 2, 1}
	selectionSort(arr)
	fmt.Println("After selection Sort", arr)
	selectionSortDes(arr)
	fmt.Println("After selection Sort", arr)
	selctionPar(arr,4)
	fmt.Println("After selection Sort", arr)
	num := []int{6,24,5,3,2,1,-32,0}
	selctionSorting(num)
	fmt.Println("After selection Sort", num)
	arr1 := []int{10,5,7,2,11,4,3,1}
	InsertSort(arr1)
	fmt.Println("Sorted array is",arr1)
	arr2 := []int{1,2,3,4,5,6,7}
	tar := 8
	arr2 = InsertSorting(arr2,tar)
	fmt.Println("Insertion of element in particular area",arr2)
}