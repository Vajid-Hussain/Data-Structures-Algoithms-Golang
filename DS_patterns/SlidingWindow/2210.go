// func getDescentPeriods(prices []int) int64 {
//     var(
//         result int64
//         count int
//     )

//     for i:=0; i<len(prices)-1; i++{
//         result++
//         if prices[i] - prices[i+1] ==1{
//             count++
//             result+=int64(count)
//         }else{
//             count=0
//         }
//     }
//     return result+1
// }