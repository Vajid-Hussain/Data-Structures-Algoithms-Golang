// func maxFrequencyElements(nums []int) int {
//     var mp = make(map[int]int)
//     for _,val:=range nums{
//         mp[val]++
//     }

//     var largest int
//     for _,occurance:=range mp{
//         if occurance > largest{
//             largest=occurance
//         }
//     }

//     var result int

//     for _,occurance:=range mp{
//         if occurance==largest{
//             result+=largest
//         }
//     }

//     return result
// }