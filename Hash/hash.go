// package main

// import "fmt"

// const arrsize = 7

// //hashtable will hold the array'

// type hashTable struct {
// 	arr [arrsize]*bucket
// }

// // Bucket node is a linked with each slot of array
// type bucket struct {
// 	head *bucketNode
// }

// // Bucket node is linked list that hold key value
// type bucketNode struct {
// 	key  string
// 	next *bucketNode
// }

// //Insertion
// func (h *hashTable)Insert (key string){
// 	index := hashing(key)
// 	h.arr[index].insert(key)
// }
// //Search
// func (h *hashTable)Search (key string) bool {
// 	index := hashing(key)
// 	return h.arr[index].search(key)
// }
// //Delete
// func (h *hashTable)Delete (key string){
// 	index := hashing(key)
// 	h.arr[index].delete(key)
	
// }

// //Bucket insertion 
// //Insert 
// func (b *bucket) insert(k string){

// 	if !b.search(k){
// 	newNode := &bucketNode{key:k}
// 	newNode.next = b.head
// 	b.head = newNode
// 	}else{
// 		fmt.Println(k,"already exist")
// 	}
	
// }
// //Search
// func (b *bucket) search(k string)bool{
// 	currentNode := b.head

// 	for currentNode != nil {
// 		if currentNode.key == k {
// 			return true
// 		}
// 		currentNode = currentNode.next
// 	}
// 	return false 
// }
// //Delete 
// func (b *bucket)delete(k string){
// 	if b.head.key == k{
// 		b.head = b.head.next
// 		return
// 	}

// 	prevNode := b.head
// 	for prevNode.next != nil {
// 		if prevNode.next.key == k{
// 			prevNode.next = prevNode.next.next
// 		}
// 	 prevNode = prevNode.next
// 	}
// }


// //Hash function 

// func hashing(key string)int{
// 	sum := 0
// 	for _,v := range key{
// 		sum += int(v)
// 	}
// 	return sum % arrsize
// }

// //init will create a bucket in each slot of hashtable 
// func Init()*hashTable{
// 	result := &hashTable{}
	
// 		for i := range result.arr{
// 			result.arr[i] = &bucket{}
// 		}
// 		return result
// }

// func findlucky(arr []int)int{
// 	freq := make(map[int]int)

// 	for _,num := range arr{
// 		freq[num]++
// 	}
	
// 	maxLucky := -1

// 	for num,count := range freq{
// 		if num == count && num > maxLucky{
// 			maxLucky = num 
// 		}
// 	}
// 	return maxLucky
// }

// func main() {
// 	hash := Init()
// 	// fmt.Println("Hashtable",hash)
// 	// fmt.Println(hashing("RANDY"))
// 	// buk := &bucket{}
// 	// buk.insert("Randy")
// 	// buk.delete("Randy")
// 	// fmt.Println(buk.search("Randy"))
// 	// fmt.Println(buk.search("Eric"))



// 	list := []string{
// 		"Vishnulal",
// 		"Nikhil",
// 		"Celestine",
// 	}

// 	for _,	v := range list{
// 		hash.Insert(v)
// 	}

// 	hash.Delete("Nikhil")
// 	fmt.Println("Nikhil",hash.Search("Nikhil"))

// 	arr := []int{1,2,3,4,5}
// 	fmt.Println("Luckiest number",findlucky(arr))
// }


package main

import "fmt"

type Hashtable struct{
	table map[int]int
}

func Newhashtable()*Hashtable{
	return &Hashtable{table:make(map[int]int)}
}

func(h *Hashtable)Add(key,val int){
	h.table[key] = val
}

func (h *Hashtable)RemoveDuplicate(){
	encount := map[int]bool{}
	for key,val := range h.table{
		if encount[val]{
			delete(h.table,key)
		}else{
			encount[val]= true
		}
	}
}

func main(){
	hsh := Newhashtable()
	hsh.Add(1,10)
	hsh.Add(2,20)
	hsh.Add(3,10)
	hsh.Add(4,30)
	hsh.Add(5,40)

	fmt.Println("Before removing dupliocate")
	for key,val := range hsh.table{
		fmt.Println("Key",key,"val",val)
	}

	hsh.RemoveDuplicate()
	fmt.Println("After Removing the element")
	for key,val := range hsh.table{
		fmt.Println("Key",key, "Val",val)
	}
}
