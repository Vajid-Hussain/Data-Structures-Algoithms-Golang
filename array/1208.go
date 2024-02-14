// func minimumSum(nums []int) int {
// 	var result = math.MaxInt
//     arr:=&nums
// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i+1; j < len(nums)-1; j++ {
// 			for k := j+1; k < len(nums); k++ {
// 				if (*arr)[i] < (*arr)[j] && (*arr)[j] > (*arr)[k] {
// 					if result > ((*arr)[i] + (*arr)[k] + (*arr)[j]) {
// 						result = (*arr)[i] + (*arr)[k] + (*arr)[j]
// 					}
// 				}
// 			}
// 		}
// 	}
// 	if result == math.MaxInt {
// 		return -1
// 	}
// 	return result
// }