// func minSteps(s string, t string) int {
//     var mp = make(map[byte]int)
//     var mp2 = make(map[byte]int)
//     var result int

//     for i:=0; i<len(s); i++{
//         mp[s[i]]++
//         mp2[t[i]]++
//     }

//     for key,val:=range mp{
//         if value,ok:= mp2[key]; ok{
//             ans:=val-value
//             if ans>0{
//                 result+=ans
//             }
//         }else{
//             result+=val
//         }
//     }
//     return result
// }
