// package main

// // func mergeSort(arr []int)[]int{
// //     if len(arr)<2{
// //         return arr
// //     }
// //     first:=mergeSort(arr[len(arr)/2:])
// //     second:=mergeSort(arr[:len(arr)/2])
// //     return merge(first,second)
// // }

// // func merge(arr1 []int, arr2 []int)[]int{
// //     var final []int
// //     i,j:=0,0
// //     for j<len(arr2) && i<len(arr1){
// //         if arr1[i]<arr2[j]{
// //             final=append(final,arr1[i])
// //             i++
// //         }else{
// //             final=append(final, arr2[j])
// //             j++
// //         }
// //     }
// //     for ;i<len(arr1); i++{
// //         final=append(final, arr1[i])
// //     }
// //     for ;j<len(arr2); j++{
// //         final=append(final,arr2[j])
// //     }
// //     return final
// // }

// func sortedSquares(nums []int) []int {
//     for key,val:=range nums{
//         nums[key]=val*val
//     }
//     fmt.Println(nums)

//     for i:=0; i<len(nums)-1; i++{
//         min:=i
//         for j:=i+1; j<len(nums); j++{
//             if nums[min]>nums[j]{
//                 min=j
//             }
//         }
//         nums[min], nums[i]= nums[i], nums[min]
//     }
//     return nums
// }