// // func deleteGreatestValue(grid [][]int) int {
// //     var result int
// //     var value int
// //     length:=len(grid)
// //     column:=len(grid[0])
// //     for key:=range grid{
// //         grid[key]=sort(grid[key])
// //     }
// //     fmt.Println(grid, length)

// //     for i:=0 ; i<column; i++{
// //         value=0
// //         for j:=0; j< length; j++{
// //             // fmt.Println(j)
// //             top:=top(&grid[j])
// //             if value < top{
// //                 value=top
// //             }
// //             // fmt.Println(top,"top")
// //         }
// //         fmt.Println(value,"**")
// //         result+=value
// //     }
// //     return result
// // }

// // func top(arr *[]int)int{
// //     // fmt.Println(arr)
// //     top:=(*arr)[0]
// //     (*arr)[0], (*arr)[len((*arr))-1]= (*arr)[len((*arr))-1], (*arr)[0]
// //     *arr=(*arr)[:len((*arr))-1]
// //     return top
// // }

// // func sort(arr []int)[]int{
// //     size:=len(arr)-1
// //     for i:=len(arr)/2; i>=0; i--{
// //         heapify(&arr, i, size)
// //     }

// //     // for i:=0; i<len(arr)-1; i++{
// //     //     arr[0], arr[size]= arr[size], arr[0]
// //     //     size--
// //     //     heapify(&arr, 0, size)
// //     // }
// //     return arr
// // }

// // func heapify(arr *[]int, index , size int){
// //     largest:=index
// //     left:=index*2 +1
// //     right:=index *2 +2

// //     if left<=size && (*arr)[left] > (*arr)[largest]{
// //         largest=left
// //     }
// //     if right<=size && (*arr)[right] > (*arr)[largest]{
// //         largest=right
// //     }
// //     if largest!= index{
// //         (*arr)[index], (*arr)[largest]= (*arr)[largest], (*arr)[index]
// //         heapify(arr, largest, size)
// //     }
// // }

// func deleteGreatestValue(grid [][]int) int {
//     res := 0
   
//     for len(grid) > 0 {
//         currMax := math.MinInt
//         flag := true
//         for idx,row := range grid {
//             if len(row) > 0 {
//             h := make(PQ, len(row))
//             copy(h,row)
//             heap.Init(&h)
//             ele := heap.Pop(&h).(int)
//             // fmt.Println("ele is ", ele)
//             if ele > currMax {
//                 currMax = ele
//             }
//             grid[idx] = h
//             if len(grid[idx]) == 0 {
//                 flag = false
//             }
//             }
//         }
//         // fmt.Println("curr Max is ", currMax)
//         res += currMax
//         if !flag {
//             break
//         }
//     }
//     return res
// }

// type PQ []int

// func (pq PQ) Len() int {
//     return len(pq)
// }

// func (pq PQ) Swap(i,j int) {
//     pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq PQ) Less(i,j int) bool {
//     return pq[i] < pq[j]
// }

// func (pq *PQ) Push(elem interface{}) {
//     *pq = append(*pq, elem.(int))
// }

// func (pq *PQ) Pop() interface{} {
//     old := *pq
// 	n := len(old)
// 	x := old[n-1]
// 	*pq = old[0 : n-1]
// 	return x
// }