

func Heapify(arr []int, size int,i int){
	largest := i
	left :=  2*i+1
	right := 2*i+2

	if left < size && arr[left] > arr[largest]{
		largest = left
	}
	 if right < size && arr[right] > arr[largest]{
		largest = right
	 }

	 if largest != i {
		arr[i],arr[largest] = arr[largest],arr[i]
	 }
	 Heapify(arr,size,largest)
}

func (b *BST)Insert(val int){
	newNode := &Node{Val :val}

	if b.root == nil {
		b.root = newNode
		return
	}

	current := b.root 

	for {
		if val < current.Val{
			if current.Left != nil{
				current = current.Left
				return
			}
			current = current.Left
		}else if val > current.Val{
			if current.Right != nil{
				current = current.Right
				
			}
			current 
		}
	}
}


