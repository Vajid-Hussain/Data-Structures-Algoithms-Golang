// package main

// import "fmt"

// func main() {
// 	matrix := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 0, 0, 1},
// 		{0, 0, 0, 1, 1},
// 		{0, 0, 0, 0, 0},
// 		{1, 0, 1, 0, 1},
// 	}

// 	fmt.Println(countIsland(matrix))
// }

// func countIsland(matrix [][]int) int {
// 	var island int
// 	visited := make([][]bool, len(matrix))
// 	for i := range matrix {
// 		visited[i] = make([]bool, len(matrix[i]))
// 	}

// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			if !visited[i][j] && matrix[i][j]==1{
// 				serch(matrix, &visited, i, j)
// 				island++
// 			}
// 		}
// 	}
// 	return island
// }

// func serch(matrix [][]int, visited *[][]bool, i, j int) {
// 	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || (*visited)[i][j] || matrix[i][j] == 0 {
// 		return
// 	}
// 	(*visited)[i][j] = true
// 	serch(matrix, visited, i+1, j)
// 	serch(matrix, visited, i-1, j)
// 	serch(matrix, visited, i, j+1)
// 	serch(matrix, visited, i, j-1)

// }
