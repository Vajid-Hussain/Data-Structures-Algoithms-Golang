// func countCompleteSubarrays(nums []int) int {
//     var (
//         distinct= make(map[int]int)
//         current= make(map[int]int)
//         result int
//         begin int
//     )

//     for _,val:=range nums{
//         distinct[val]++
//     }

//     // fmt.Println(distinct)
//     // fmt.Println(current)

//     // for i:=0; i<len(nums); i++{
//     //     current[nums[i]]++
//     //     // fmt.Println(current)
//     //     if i-begin >= len(distinct)-1 && mapEqual(distinct, current) {
//     //         // fmt.Println("inner loop")
//     //         for mapEqual(distinct, current) {
//     //             // fmt.Println(distinct, current, "===")
//     //             result+= len(nums) - i
//     //             decrementMap(current, nums[begin])
//     //             begin++
//     //         }
//     //     }
//     // }

//     for i:=0; i<len(nums); i++{
//         current[nums[i]]++
//         for len(current) == len(distinct) {
//             current[nums[begin]]--
//             if current[nums[begin]]==0{
//                 delete(current, nums[begin])
//             }
//             result+= len(nums)-i
//             begin++
//         }
//     }

//     return result
// }

// // func mapEqual(mp1, mp2 map[int]int) bool{
// //     for key:=range mp1{
// //         if _,ok:= mp2[key]; !ok{
// //             return false
// //         }
// //     }
// //     return true
// // }

// // func decrementMap(mp map[int]int, value int) {
// //     // fmt.Println(mp, value)
// //      val,ok:= mp[value]
// //      if val>1 && ok{
// //         mp[value]--
// //      }else{
// //         delete(mp,value)
// //      }
// // }