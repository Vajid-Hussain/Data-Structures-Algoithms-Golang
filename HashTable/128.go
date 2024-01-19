// func longestConsecutive(nums []int) int {
//     var mp=make(map[int]struct{})
//     for _,val:=range nums{
//         mp[val]=struct{}{}
//     }

//     var longest int
//     var length int
//     for val:=range mp{
//         if _,ok:=mp[val-1]; !ok{
//             length++
//             for _,ok:=mp[val+length]; ok; _,ok=mp[val+length]{
//                 length++
//             }
//             longest=max(longest,length)
//             length=0
//         }  
//     }
//     return longest
// }

// func max(first, second int)int{
//     if first>second{
//         return first
//     }
//     return second
// }