// func checkInclusion(s1 string, s2 string) bool {
//     var(
//         mp= make(map[byte]int)
//         j int
//     )
//     for i:=0; i<len(s1); i++{
//         mp[s1[i]]++
//     }

//     for i:=len(s1)-1; i<len(s2); i++{
//        result:= checking(mp, s2, j, i)
//        if result{
//         return true
//        }
//        j++
//     }
//     return false
// }

// func checking(mp map[byte]int, s string, start, end int)bool{
//     // fmt.Println(start, end)
//     var mp2 =make(map[byte]int)
//     for i:=start; i<=end; i++{
//         mp2[s[i]]++
//     }

//     if check(mp, mp2){
//         return true
//     }
//     return false
// }

// func check(mp, mp2 map[byte]int)bool{
//     // fmt.Println("----------",mp, mp2)
//     if len(mp)!=len(mp2){
//         return false
//     }

//     for key, val:=range mp{
//         if  val2:= mp2[key]; val != val2{
//             return false
//         }
//     }
//     return true
// }