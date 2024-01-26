// package main

// import "fmt"

// type graph struct {
// 	vertices []*vertex
// }

// type vertex struct {
// 	data     int
// 	adgecent []*vertex
// 	visited  bool
// }

// func (g *graph) insert(data int) {
// 	for _, val := range g.vertices {
// 		if val.data == data {
// 			return
// 		}
// 	}
// 	g.vertices = append(g.vertices, &vertex{data: data})
// }

// func (g *graph) removeVertex(data int) {
// 	for i, val := range g.vertices {
// 		if val.data == data {
// 			g.vertices = append(g.vertices[:i], g.vertices[1+i:]...)
// 		}
// 	}
// 	g.removeEdgeFromOtheVertex(data)
// }

// func (g *graph) removeEdgeFromOtheVertex(data int) {
// 	for _, val := range g.vertices {
// 		for i, value := range val.adgecent {
// 			if value.data == data {
// 				val.adgecent = append(val.adgecent[:i], val.adgecent[1+i:]...)
// 			}
// 		}
// 	}
// }

// func (g *graph) addEdge(from, to int) {
// 	fromVertex := g.getVertex(from)
// 	toVertex := g.getVertex(to)
// 	if fromVertex == nil || toVertex == nil {
// 		fmt.Println("can't make connection due to absentce of vertex")
// 		return
// 	}
// 	fromVertex.adgecent = append(fromVertex.adgecent, toVertex)
// 	// toVertex.adgecent = append(toVertex.adgecent, fromVertex)
// }

// func (g *graph) removeEdge(from, to int) {
// 	fromVertex := g.getVertex(from)
// 	toVertex := g.getVertex(to)
// 	if fromVertex == nil || toVertex == nil {
// 		fmt.Println("no vertices")
// 		return
// 	}
// 	g.removeEdgeLoop(from, toVertex)
// 	g.removeEdgeLoop(to, fromVertex)
// }

// func (g *graph) removeEdgeLoop(data int, vertex *vertex) {
// 	for i, val := range vertex.adgecent {
// 		if val.data == data {
// 			vertex.adgecent = append(vertex.adgecent[:i], vertex.adgecent[i+1:]...)
// 			return
// 		}
// 	}
// }

// func (g *graph) print() {
// 	for _, val := range g.vertices {
// 		fmt.Print(val.data, "  :")
// 		for _, value := range val.adgecent {
// 			fmt.Print(value.data, " ")
// 		}
// 		fmt.Println("")
// 	}
// }

// func (g *graph) getVertex(data int) *vertex {
// 	for _, val := range g.vertices {
// 		if val.data == data {
// 			return val
// 		}
// 	}
// 	return nil
// }

// type queue struct {
// 	arr []int
// }

// func (g *graph) BFS(vertex int) {
// 	queue := queue{}
// 	visited := make(map[int]bool)
// 	queue.arr = append(queue.arr, vertex)
// 	visited[vertex] = true
// 	for len(queue.arr) != 0 {
// 		// fmt.Println("*",queue.arr)
// 		fmt.Println(queue.arr[0])
// 		vertex = queue.arr[0]
// 		queue.arr = queue.arr[1:]
// 		for _, val := range g.getVertex(vertex).adgecent {
// 			// fmt.Println("&")
// 			if !visited[val.data] {
// 				visited[val.data] = true
// 				queue.arr = append(queue.arr, val.data)
// 			}
// 		}
// 	}

// }

// type stack struct {
// 	arr []int
// }

// func (g *graph) DFS(vertex int) {
// 	stack := stack{}
// 	visited := make(map[int]bool)
// 	stack.arr = append(stack.arr, vertex)
// 	visited[vertex] = true
// 	for len(stack.arr) != 0 {
// 		vertex = stack.arr[len(stack.arr)-1]
// 		fmt.Println(vertex)
// 		stack.arr = stack.arr[:len(stack.arr)-1]
// 		for _, val := range g.getVertex(vertex).adgecent {
// 			if !visited[val.data] {
// 				visited[val.data] = true
// 				stack.arr = append(stack.arr, val.data)
// 			}
// 		}
// 	}
// }

// // func (g *graph)isCyclic()bool{
// // 	for _,val:=range g.vertices{
// // 		if !val.visited && g.iscyclicUtil(val, nil){
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func (g *graph) iscyclicUtil(v *vertex, parent *vertex) bool{
// // 	v.visited=true
// // 	for _,adj:=range v.adgecent{
// // 		if !adj.visited{
// // 			if g.iscyclicUtil(adj, v){
// // 				return true
// // 			}else if adj != parent{
// // 				return true
// // 			}
// // 		}
// // 	}
// // 	return false
// // }

// func (g *graph) isCyclic()bool{
// 	for _,val:=range g.vertices{
// 		if !val.visited && g.iscyclicUtil(val, nil){
// 			return true
// 		}
// 	}
// 	return false
// }

// func (g *graph)iscyclicUtil(v *vertex, parent *vertex)bool{
// 	v.visited=true

// 	for _,val:=range v.adgecent{
// 		if !val.visited{
// 			if g.iscyclicUtil(val, v){
// 				return true
// 			}
// 		}else if parent!= val{
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	gp := graph{}
// 	gp.insert(1)
// 	gp.insert(2)
// 	gp.insert(3)
// 	// gp.insert(1)
// 	// gp.insert(8)
// 	// gp.insert(6)
// 	gp.addEdge(1, 2)
// 	gp.addEdge(2, 3)
// 	// gp.addEdge(3, 1)
// 	// gp.addEdge(1, 2)
// 	// // gp.removeEdge(1, 2)
// 	// // gp.removeVertex(6)
// 	// gp.print()
// 	fmt.Println("bfs")
// 	gp.BFS(1)
// 	fmt.Println("is cyclic:",gp.isCyclic())
// 	// fmt.Println("dfs")
// 	// gp.DFS(1)
// }
