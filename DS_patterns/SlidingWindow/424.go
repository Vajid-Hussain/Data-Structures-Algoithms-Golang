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
