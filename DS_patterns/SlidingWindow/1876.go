// func countGoodSubstrings(s string) int {
//     var cound int
//     var first byte
//     var second byte
//     for i:=2;i<len(s); i++{
//         first=s[i-2]
//         second=s[i-1]
//         if first != s[i] && second != s[i] && first!= second{
//             cound++
//         }
//     }
//     return cound
// }