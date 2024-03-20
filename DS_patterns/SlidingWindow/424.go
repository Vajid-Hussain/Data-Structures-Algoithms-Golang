// func characterReplacement(s string, k int) int {
//     var(
//         frequency int
//         mp =make(map[byte]int)
//         j int
//         totalLength int
//         max int
//     )

//     for i:=0; i<len(s); i++{
//         mp[s[i]]++
//         totalLength, max=maxlen(mp)
//         if totalLength-max >k{
//             // fmt.Println(mp)
//             for j<i {
//                 delete(mp, s[j])
//                 // fmt.Print("---",j)
//                 totalLength, max = maxlen(mp)
//                 if totalLength-max <= k{
//                     j++
//                     break
//                 }
//                 j++
//             }
//         }
//         fmt.Println(mp,j,i)

//         if frequency < totalLength{
//             // fmt.Println( j, i,totalLength)
//             frequency= totalLength
//         }

//     }
//     return frequency
// }

// func maxlen(mp map[byte]int)(totalLength int, max int){
//     // fmt.Println(mp)
//     for _,val:=range mp{
//         totalLength+=val
//         if max<val{
//             max=val
//         }
//     }
//     return
// }

// //full correct 
// func characterReplacement(s string, k int)int{
//     count:= make(map[byte]int)
//     var res, maxf, left int

//     for right:=0; right<len(s); right++{
//         count[s[right]]+=1

//         maxf = max(maxf, count[s[right]])

//         if (right-left+1) - maxf>k{
//             count[s[left]]--
//             left++
//         }

//         res= max(res, right-left+1)
//     }
//     return res
// }

// func max(a, b int)int{
//     if a>b{
//         return a
//     }
//     return b
// }