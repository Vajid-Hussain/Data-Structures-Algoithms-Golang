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
// 	fmt.Println(count(matrix))
// }

// func count(matrix [][]int) int {
// 	island := 0
// 	visited := make([][]bool, len(matrix))

// 	for i := range matrix {
// 		visited[i] = make([]bool, len(matrix[i]))
// 	}

// 	for i := range matrix {
// 		for j := range visited {
// 			if matrix[i][j] == 1 && !visited[i][j] {
// 				dfs(matrix, &visited, i, j)
// 				island++
// 			}
// 		}
// 	}
// 	return island
// }

// func dfs(matrix [][]int, visited *[][]bool, i, j int) {
// 	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || matrix[i][j] == 0 || (*visited)[i][j] {
// 		return
// 	}
// 	(*visited)[i][j] = true

// 	dfs(matrix, visited, i+1, j)
// 	dfs(matrix, visited, i-1, j)
// 	dfs(matrix, visited, i, j+1)
// 	dfs(matrix, visited, i, j-1)
// }
