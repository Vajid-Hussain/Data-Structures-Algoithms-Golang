// package main

// import "fmt"

// type graph struct {
// 	verices []*vertex
// }

// type vertex struct {
// 	data     int
// 	adjacent []*vertex
// }

// func (g *graph) addvertex(data int) {
// 	if exist := g.vertexExist(data); exist == nil {
// 		g.verices = append(g.verices, &vertex{data: data})
// 	}
// }

// func (g *graph) addEdge(from, to int) {
// 	fromVertex := g.vertexExist(from)
// 	toVertex := g.vertexExist(to)
// 	if fromVertex == nil || toVertex == nil {
// 		fmt.Println("can't make edge missing vertexes")
// 		return
// 	}
// 	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
// 	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
// }

// func (g *graph) vertexExist(data int) *vertex {
// 	for _, val := range g.verices {
// 		if val.data == data {
// 			return val
// 		}
// 	}
// 	return nil
// }

// func (g *graph) print() {
// 	for _, val := range g.verices {
// 		fmt.Print("vertex is:", val.data)
// 		for _, val := range val.adjacent {
// 			fmt.Println(" ", val.data)
// 		}
// 	}
// }

// func (g *graph) deleteaVertex(data int){
// 	for i,val:=range g.verices{
// 		if val.data==data{
// 			g.verices=append(g.verices[:i], g.verices[1+i:]...)
// 		}
// 	}
// }

// func (g *graph)deleteAdjesent(data int){
// 	for i,val:=range g.verices{
// 		if val.data==data{
// 			g.verices=append(g.verices[:i], g.verices[1+i:]...)
// 			g.
// 		}
// 	}
// }

// func main() {
// 	graph := graph{verices: []*vertex{}}
// 	graph.addvertex(3)
// 	graph.addvertex(5)
// 	graph.addvertex(7)
// 	graph.addEdge(3, 5)
// 	graph.addEdge(3, 4)
// 	graph.deleteaVertex(3)
// 	graph.print()
// }
