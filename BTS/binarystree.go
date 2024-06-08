package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySTree struct {
	Root *Node
}

// Insert into a new Node into BST
func (b *BinarySTree) Insert(val int) {

	newNode := &Node{Val: val}
	if b.Root == nil {
		b.Root = newNode
		return
	}

	current := b.Root

	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}

}

//// Check if the BST contains a given value

func (b *BinarySTree) Contains(val int) bool {

	current := b.Root

	for current != nil {
		if val == current.Val {
			return true
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

//Delete the node

func (b *BinarySTree) Delete(val int) {
	b.Root = deleteNode(b.Root, val)
}

func deleteNode(root *Node, val int) *Node {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = deleteNode(root.Left, val)
		return root
	} else if val > root.Val {
		root.Right = deleteNode(root.Right, val)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		successor := findMin(root.Right)
		root.Val = successor.Val
		root.Right = deleteNode(root.Right, val)
		return root
	}
}

func findMin(node *Node) *Node {
	current := node

	for current.Left != nil {
		current = current.Left
	}
	return current
}

//Traversal towards BST - Inorder Travesal , Preoder Travesal , Postorder Travesal

// Inorder Travesal - Left --> Root --> Right
func (b *BinarySTree) InorderTravesal() {
	inorderTravesal(b.Root)
}
func inorderTravesal(node *Node) {
	if node != nil {
		inorderTravesal(node.Left)
		fmt.Print(node.Val," ")
		inorderTravesal(node.Right)
	}
}

//Preorder Travesal Root -- > Left--> Right 
func (b *BinarySTree)PreoderTravesal(){
	preorderTravesal(b.Root)
}

func preorderTravesal(node *Node){
	if node != nil{
		fmt.Print(node.Val," ")
		preorderTravesal(node.Left)
		preorderTravesal(node.Right)
	}
}


//Postorder Travesal Left----> Right--->Root 

func (b *BinarySTree)PostorderTravesal(){
	postorderTravesal(b.Root)
}

func postorderTravesal(node *Node){
	if node != nil{
		postorderTravesal(node.Left)
		postorderTravesal(node.Right)
		fmt.Print(node.Val," ")
	}
}


func findHeight(node *Node)int{
	if node == nil{
		return -1
	}
	left := findHeight(node.Left)
	right := findHeight(node.Right)

	return max(left,right)+1
}

func isBalance(node *Node)bool{
	left := findHeight(node.Left)
	right := findHeight(node.Right)

	return max(left,right) - min(left,right) <= 1
}



func main() {
	bst := &BinarySTree{}

	bst.Insert(20)
	bst.Insert(30)
	bst.Insert(40)
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(60)
	bst.Insert(50)

	fmt.Println("InorderTravesal")
	bst.InorderTravesal()
	fmt.Println()

	fmt.Println("PreorderTravesal")
	bst.PreoderTravesal()
	fmt.Println()

	fmt.Println("PostOrderTravesal")
	bst.PostorderTravesal()
	fmt.Println()

	fmt.Println("Contains 60",bst.Contains(60))
	fmt.Println("Contains 7",bst.Contains(7))

	bst.Delete(10)
	fmt.Println("Inorder Travesal after deletion:")
	bst.InorderTravesal()
	fmt.Println()

	height :=findHeight(bst.Root)
	fmt.Println("Height of BST",height)

	balanced := isBalance(bst.Root)
	fmt.Println("Height of tree is balanced",balanced)


}


