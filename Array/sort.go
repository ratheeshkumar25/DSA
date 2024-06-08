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

func SecondLarget(arr[]int)int{

	if len(arr)<2{
		return-1
	}
	lar := arr[0]
	secondlar := -1

	for i := 0 ; i < len(arr); i++{
		if arr[i]>lar{
			secondlar = lar
			lar = arr[i]
		}else if arr[i]<lar &&arr[i]>secondlar{
			secondlar = arr[i]
		}
	}
	return secondlar
}

func RemoveDuplicate(arr []int)[]int{
	// num := []int{}
	// for i :=0 ; i<len(arr)-1;i++{
	// 	isDuplicate := false
	// 	for j:=i+1;j<len(arr);j++{
	// 		if arr[i]  == arr[j]{
	// 			isDuplicate = true
	// 			break
	// 		}
	// 	}
	// 	if !isDuplicate{
	// 		num = append(num,arr[i])
	// 	}
	// }
	// return num

	removeMap := make(map[int]bool)
	unique := []int{}
	
	for _,val := range arr{
		if !removeMap[val] {
			removeMap[val] = true
			unique = append(unique, val)
		}
	}
	return unique
}

func ReverseArray(arr []int){
	start :=0
	end := len(arr)-1

	for start < end {
		arr[start],arr[end] = arr[end],arr[start]
		start++
		end--
	}

}

func main(){
	arr := []int{24,15,15,22,15,10,33,44,15}
	fmt.Println("After removing the duplicate element",RemoveDuplicate(arr))
    sortArray(arr)
	fmt.Println("Result of Sorted Array",arr)
	
	fmt.Println("Second largest",SecondLarget(arr))
	ReverseArray(arr)
	fmt.Println("After reverse array",arr)
	
}