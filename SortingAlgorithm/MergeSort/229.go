// package main
// func mergeSort(arr []int)[]int{
//     if len(arr)<2{
//         return arr
//     }
//     first:=mergeSort(arr[:len(arr)/2])
//     second:=mergeSort(arr[len(arr)/2:])
//     result:=merge(first,second)
//     return result
// }

// func merge(arr1, arr2 []int)[]int{
//     final:=[]int{}
//     i,j:=0,0
//     for i<len(arr1) && j< len(arr2){
//         if arr1[i]>arr2[j]{
//             final=append(final, arr1[i])
//             i++
//         }else{
//             final=append(final, arr2[j])
//             j++
//         }
//     }

//     for ;i<len(arr1); i++{
//         final=append(final, arr1[i])
//     }
//     for ;j<len(arr2); j++{
//         final=append(final, arr2[j])
//     }
//     return final
// }

// func majorityElement(nums []int) []int {
//     result:=mergeSort(nums)
//     fmt.Println(result)

//     var prev,i int
//     arr:=[]int{}
//     var count int
//     for i=0; i<len(nums); i++{
//         if count==0{
//             count++
//             fmt.Println(result[i],i)
//             prev=result[i]
//             continue
//         }

//         if result[i]==prev{
//             count++
//         }else{
//             if count>len(nums)/3{
//                 arr=append(arr, result[i-1])
//                 count=0
//                 i--
//             }else{
//                 count=0
//                 i--
//             }
//         }
//     }

//     if count>len(nums)/3{
//         arr=append(arr,result[i-1])
//     }
//     return arr
// }