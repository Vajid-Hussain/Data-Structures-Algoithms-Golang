// func sumIndicesWithKSetBits(nums []int, k int) int {
//     var result int
//     var binary int
//     var count int
//     for i:=0; i<len(nums); i++{
//         binary=Binary(i)
//         count=OnesCount(binary)
//         if count==k{
//             result+=nums[i]
//         }
//     }
//     return result
// }

// func Binary(data int)int{
//     value:= strconv.FormatInt(int64(data), 2)
//     intvalue,_:= strconv.Atoi(value)
//     return intvalue
// }

// func OnesCount(data int) int{
//     var count int
//     var rem int
//     for data >0{
//         rem=data%10
//         if rem==1{
//             count++
//         }
//         data=data/10
//     }
//     return count
// }