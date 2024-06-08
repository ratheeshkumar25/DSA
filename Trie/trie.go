	package main

	import "fmt"

	//Alphabet Size is the number of possible characters in trie 

	const Alphabetsize = 26


	//Node
	type  Node struct{
		children [Alphabetsize]*Node
		isEnd bool
	}

	//Trie

	type Trie struct{
		root *Node
	}

	//init trie create the new Trie 
	func InitTrie() *Trie{
		result := &Trie{root:&Node{}}
		return result 
	}

	//Insert  will take in word and add it to trie
	func (t *Trie)Insert(s string){
		n := len(s)
		currrentNode := t.root
		for i :=0; i <n; i++{
			charIndex := s[i]-'a'
			if currrentNode.children[charIndex] == nil{
				currrentNode.children[charIndex] = &Node{}
			}
			currrentNode = currrentNode.children[charIndex]
		}
		currrentNode.isEnd = true
	}

	//Search

	func (t *Trie)Search(w string) bool{
		n := len(w)
		currentNode := t.root

		for i:=0 ; i<n ; i++{
			charIndex := w[i] -'a'
			if currentNode.children[charIndex] == nil {
				return false
			}
			currentNode = currentNode.children[charIndex]
		}

		if currentNode.isEnd == true{
			return true
		}
		
		return false 
	}

	func main(){
		testTrie := InitTrie()

		toAdd := []string{
			"aragaon",
			"revathy",
			"ratheesh",
		}

		for _,v := range toAdd{
			testTrie.Insert(v)
		}
		fmt.Println("testTrie",toAdd)

		fmt.Println("Word available in trie",testTrie.Search("ratheeshg"))

	}