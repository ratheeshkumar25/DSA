package main

import "fmt"

//Create a heap
func CreateHeap(arr []int,n int){
	for i := (n/2)-1 ; i >= 0 ; i--{
		Heapify(arr,n,i)
	}
}

func left(i int)int{
	return 2*i+1
}

func right(i int)int{
	return 2*i+2
}

func Heapify(arr []int,size int ,i int){
	largest := i
	l := left(i)
	r := right(i)

	//if left child is largest than root 
	if l < size && arr[l] > arr[largest]{
		largest = l
	}

	if r < size && arr[r] > arr[largest]{
		largest = r
	}
	
	if largest != i {
		arr[i],arr[largest] = arr[largest],arr[i]
		Heapify(arr,size,largest)
	}
}

func HeapSort(arr []int){
	n := len(arr)
	CreateHeap(arr,n)
	for i := len(arr)-1; i >=0 ; i--{
		arr[0],arr[i] = arr[i],arr[0]
		Heapify(arr,i,0)
	}
}


func main(){
	arr := []int{45,25,55,65,75,80}
	HeapSort(arr)
	fmt.Println("After the heap sort array will be",arr)
}

package main

import "fmt"


func CreatHeap(arr[]int){
	for i := (len(arr)/2)-1 ; i >=0 ; i--{
		Heapify(arr,len(arr),i)
	}
}

func left(i int)int{
	return 2*i+1
}

func right(i int)int{
	return 2*i+2
}

func Heapify(arr[]int, size int,i int){
	largest := i
	l := left(i)
	r := right(i)

	if l < size && arr[l] > arr[largest]{
		largest = l
	}

	if r < size && arr[r]> arr[largest]{
		largest = r
	}

	if largest != i {
		arr[i],arr[largest] = arr[largest],arr[i]
		Heapify(arr,size,largest)
	}
}

func HeapSort(arr []int){
	CreatHeap(arr)
	for i := len(arr)-1; i >= 0; i--{
		arr[0],arr[i] = arr[i], arr[0]
		Heapify(arr,i,0)
	}
}

func findMax(arr []int)int{
	return arr[0]
}

func main(){
	arr := []int{45,55,65,25,20,15,10}
	fmt.Println("Max element in Max",findMax(arr))
	HeapSort(arr)
	fmt.Println("After the heap sort array will be",arr)

}