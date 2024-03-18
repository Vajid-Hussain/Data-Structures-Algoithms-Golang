// func dailyTemperatures(temperatures []int) []int {
// 	var stack []int
//     var result = make([]int, len(temperatures))
// 	for i,_:=range temperatures{
//         for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]]{
//             index:= stack[len(stack)-1]
//             stack= stack[:len(stack)-1]
//             result[index]= i-index
//         }
//         stack= append(stack, i)
//     }
// 	return result
// }
