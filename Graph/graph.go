package main

import "fmt"

type GraphNode struct {
	Vertex map[int][]int
}

func (g *GraphNode) Insert(vertex, edge int, isBi bool) {
	g.Vertex[vertex] = append(g.Vertex[vertex], edge)

	if isBi {
		g.Vertex[edge] = append(g.Vertex[edge], vertex)
	}
}

func (g *GraphNode) Display() {
	for vertex, edge := range g.Vertex {
		fmt.Println(vertex,edge)
	}
}

func (g *GraphNode)RemoveEdge(val, val1 int, isBi bool){
	arr := g.Vertex[val]
	for i,v := range arr{
		if v == val1 {
			g.Vertex[val] = append(arr[:i], arr[i+1:]...)
			break 
		}
	}
	if len(g.Vertex[val]) == 0{
		delete(g.Vertex,val)
	}

	if isBi{
		arr := g.Vertex[val1]
		for i,v := range arr{
			if v == val{
				g.Vertex[val] = append(arr[:i],arr[i+1:]...)
				break
			}
		}
		if len(g.Vertex[val1]) ==  0{
			delete(g.Vertex,val1)
		}
	}
}

func (g *GraphNode)RemoveVertex(val int){
	delete(g.Vertex,val)
	for key,arr := range g.Vertex{
		for i,e := range arr{
			if e == val{
				g.Vertex[key] = append(arr[:i],arr[i+1:]...)
			}
		}
	}
}

func (g *GraphNode)BFS(val int){
	visited := make(map[int]bool)
	visited[val] = true

	arr := []int{val}

	for len(arr)>0{
		value := arr[0]
		fmt.Println(value)
		arr = arr[1:]
		for _,x := range g.Vertex[val]{
			if !visited[x]{
				visited[x] = true
				arr = append(arr, x)
			}
		}
	}
}

func (g *GraphNode)DFS(val int){
	visited := make(map[int]bool)
	visited[val] = true 
	stack := []int{val}
	g.DFShelper(val,visited,stack)
}

func (g *GraphNode)DFShelper(val int,visited map[int]bool,stack[]int){
	if len(stack) == 0 {
		return	
	}
	value := stack[len(stack)-1]
	fmt.Println(value)
	stack = stack[:len(stack)-1]

	for _,v := range g.Vertex[val]{
		if !visited[v]{
			stack = append(stack, v)
			visited[v]= true
			g.DFShelper(v,visited,stack)
		}
	}
}

func main(){
	g := GraphNode{Vertex: make(map[int][]int)}
	g.Insert(1, 2, false)
	g.Insert(2, 3, false)
	g.Insert(3, 4, true)
	g.Insert(2, 5, false)
	g.Insert(1, 5, false)
	g.Insert(6, 7, false)
	g.Insert(6, 2, false)
	g.Insert(3, 2, false)
	g.Insert(4, 5, true)

	g.RemoveEdge(4, 5, true)
	g.RemoveVertex(3)
	g.Display()
	g.DFS(2)
	g.BFS(2)
}