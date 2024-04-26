// func findWinners(matches [][]int) [][]int {
//     var mp =map[int]int{}
//     var mp2= map[int]int{}
//     var(
//          result [][]int
//          arr1 []int
//          arr2 []int
//     )


//     for _,val:=range matches{
//         mp[val[0]]++
//         mp2[val[1]]++
//     }

//     for key,_:=range mp{
//         if _,ok:=mp2[key]; !ok{
//             arr1=append(arr1, key)
//         }
//     }

//     for key,val:=range mp2{
//         if val==1{
//             arr2=append(arr2, key)
//         }
//     }

//     slices.Sort(arr1)
//     slices.Sort(arr2)
//     // fmt.Println(arr1, arr2)
//     result=append(result, arr1)
//     result=append(result, arr2)
//     return result
// }