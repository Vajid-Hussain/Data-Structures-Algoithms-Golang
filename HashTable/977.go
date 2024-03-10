// func findJudge(n int, trust [][]int) int {
//     if len(trust)==0 && n==1{
//         return n
//     }
//     var mp = make(map[int]int)
//     var judge int
//     for _,val :=range trust{
//         mp[val[1]]++
//     }
//     for key,count:=range mp{
//         fmt.Println(key,count)
//         if count == n-1{
//             judge= key
//             break
//         }
//     }

//     if judge==0{
//         return -1
//     }
    
//     for _,val :=range trust{
//         if val[0] == judge{
//             return -1
//         }
//     }
    
//     return judge
// }
