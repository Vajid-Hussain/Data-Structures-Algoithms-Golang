// func maxArea(height []int) int {
//     var waterLeavel int
//     var diffrence int
//     var i=0
//     var j= len(height)-1
//     for i<j{
//         currentLeavel:=0
//         diffrence=j-i
//         if height[i]>height[j] {
//             currentLeavel=height[j] * diffrence
//             j--
//         }else{
//             currentLeavel=height[i] * diffrence
//             i++
//         }
//         if waterLeavel< currentLeavel{
//             waterLeavel=currentLeavel
//         }
//     }
//     return waterLeavel
// }