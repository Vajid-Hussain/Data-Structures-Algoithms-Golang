// func compress(chars []byte) int {
// 	var (
//         count =1
// 		start  int
// 	)
//     for i:=1; i<=len(chars) ;i++{
//         if i<len(chars) && chars[i]==chars[i-1]{
//             count++
//         }else{
//             chars[start] = chars[i-1]
//             start++
//             if count>1{
//                 for _,ch:=range strconv.Itoa(count) {
//                     chars[start] =byte(ch)
//                     start++
//                 }
//             } 
//             count=1
//         }
//     }
//     return start
// }