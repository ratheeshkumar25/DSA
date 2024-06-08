package main

import "fmt"

//Buckenode declared
type bucketNode struct {
	key int 
	next *bucketNode
}

type bucket struct{
	 head *bucketNode
}

type hashtable struct{
	arr[]*bucket
	size int 
}

//init will be create bucket to store each slot 
func Init(size int)*hashtable{
	result := &hashtable{
		arr: make([]*bucket,size), size: size,
	}

	for i := range result.arr{
		result.arr[i] = &bucket{}
	}
	return result 
}
//hashing key 
func (h *hashtable)hashing(key int)int {
	return key % h.size
}
//Insert to hashtable
func (h *hashtable)Insert(key int){
	index := h.hashing(key)
	h.arr[index].insert(key)
}
//Search to hashtable 
func (h *hashtable)Search(key int){
	index := h.hashing(key)
	h.arr[index].search(key)
}
//Delete
func (h *hashtable)Delete(key int){
	index := h.hashing(key)
	h.arr[index].delete(key)
}

//Display
func (h *hashtable)Display(){
	for i,b := range h.arr{
		fmt.Printf("Bucket %d : " ,i)
		b.display()
	}
}

//Bucket 

//Insertion to bucket ]
func(b *bucket) insert(k int){
	if !b.search(k){
		newNode := &bucketNode{key:k }
		newNode.next = b.head
		b.head = newNode
	}else {
		fmt.Println(k,"already exist")
	}
}


//Search 
func(b *bucket) search(k int)bool{
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//Delete 

func (b *bucket)delete(k int){
	if b.head == nil {
		return 
	}
	if b.head.key == k {
		return 
	}
	prevNode := b.head 
	for prevNode.next != nil {
		if prevNode.next.key == k{
			prevNode.next = prevNode.next.next 
			return
		} 
		prevNode = prevNode.next
	}	
}

//Display 

func (b *bucket)display(){
	currentNode := b.head
	for currentNode != nil {
		fmt.Printf("%d->",currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main(){
	hsize := 7
	hash := Init(hsize)

	list := []int{
		15,
		20,
		30,
		40,
		50,
		60,
		70,
	}

	for _,v := range list{
		hash.Insert(v)
	}
	hash.Display()
	fmt.Println("Search",hash)
}


