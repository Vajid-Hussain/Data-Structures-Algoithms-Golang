package main

import "fmt"

type Graph struct {
	vertices map[int]map[int]bool
}

func (g *Graph) addVertices(vertex int) {
	if _, ok := g.vertices[vertex]; !ok {
		g.vertices[vertex] = make(map[int]bool)
	}
}

func (g *Graph) addEdge(vertex1, vertex2 int) {
	g.addVertices(vertex1)
	g.addVertices(vertex2)

	g.vertices[vertex1][vertex2] = true
	g.vertices[vertex2][vertex1] = true
}

func (g *Graph) print() {
	for key, val := range g.vertices {
		for neiber := range val {
			fmt.Println("graph: ", key, neiber)
		}
	}
}

func main(){
	graph:=Graph{vertices: make(map[int]map[int]bool)}
	graph.addEdge(1,2)
	graph.addEdge(2,3)
	graph.print()
	fmt.Println(graph.vertices)
}