// func isPalindrome(s string) bool {
//     j:=len(s)-1
//     for i:=0; i<j; i++{
//         val:=s[i]
//         if (val>=65 && val<=90) || (val>= 97 && val<= 122) || (val>=48 && val <=57){
//             if val>=65 && val <=90{
//                 val+=32
//             }
//             for j>(len(s)/2)-1 {
//                 endVal:=s[j]
//                 if (endVal>=65 && endVal<=90) || (endVal>= 97 && endVal<= 122) || (endVal>=48 && endVal <=57){
//                     if endVal>=65 && endVal <=90{
//                         endVal+=32
//                     }
//                     if val!=endVal{
//                         return false
//                     }
//                     j--
//                     break
//                 }
//                 j--
//             }
//         }
//     }
//     return true
// }