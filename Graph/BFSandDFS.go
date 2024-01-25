package main

import "fmt"

type graph struct {
	vertices []*vertex
}

type vertex struct {
	data     int
	adjacent []*vertex
}

func (g *graph) insert(data int) {
	for _, val := range g.vertices {
		if val.data == data {
			return
		}
	}
	g.vertices = append(g.vertices, &vertex{data: data})
}

func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertices(from)
	toVertex := g.getVertices(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("vertes are null can't make a edges")
		return
	}
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}

func (g *graph) print() {
	for _, val := range g.vertices {
		fmt.Print("vertices: ", val.data, "    :")
		for _, value := range val.adjacent {
			fmt.Print(" ", value.data)
		}
		fmt.Println("")
	}
}

type queue struct {
	arr []int
}

func (g *graph) BFS(vertex int) {
	queue := queue{}
	visited := make(map[int]bool)
	queue.arr = append(queue.arr, vertex)
	visited[vertex] = true
	for len(queue.arr) != 0 {
		vertex = queue.arr[0]
		fmt.Println("**",vertex)
		queue.arr = queue.arr[1:]
		for _, val := range g.getVertices(vertex).adjacent {
			if !visited[val.data] {
				visited[val.data] = true
				queue.arr = append(queue.arr, val.data)
			}
		}
	}
}

type stacks struct{
	arr []int
}

func (g *graph) DFS(vertex int){
	stack:=stacks{}
	visited:=make(map[int]bool)
	stack.arr=append(stack.arr, vertex)
	visited[vertex]=true
	for len(stack.arr)!=0{
		vertex=stack.arr[len(stack.arr)-1]
		stack.arr=stack.arr[:len(stack.arr)-1]
		fmt.Println(vertex)
		for _,val:=range g.getVertices(vertex).adjacent{
			if !visited[val.data]{
				visited[val.data]=true
				stack.arr=append(stack.arr, val.data)
			}
		}
	}
}

func (g *graph) getVertices(data int) *vertex {
	for _, val := range g.vertices {
		if val.data == data {
			return val
		}
	}
	return nil
}

func main() {
	gp := graph{}
	gp.insert(5)
	gp.insert(4)
	gp.insert(9)
	gp.insert(1)
	gp.addEdge(1, 9)
	gp.addEdge(9, 4)
	gp.addEdge(4, 5)
	gp.addEdge(1, 5)
	gp.print()
	gp.BFS(1)
	fmt.Println("**")
	gp.DFS(5)
}
