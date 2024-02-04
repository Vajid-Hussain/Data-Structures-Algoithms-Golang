// package main
// func search(nums []int, target int) int {
//     if len(nums)==1 {
//         if nums[0]==target{
//             return 0
//         }else{
//             return -1
//         }
//     }

//     var min, max int
//     var point int
//     result :=false

//     for i:=0; i<len(nums)-1; i++{
//         if nums[i]>nums[i+1]{
//             point=i
//             result =true
//             break
//         }
//     }

//     if target>=nums[point] && result==true || target <=nums[point] && target >= nums[0]{
//         min=0
//         max=point
//     }else{
//         min=point+1
//         max=len(nums)-1
//     }

//     if result==false {
//         min=0
//         max=len(nums)-1
//     }

//     for min<=max{
//         mid:=(min+max)/2
//             fmt.Println(min,max)

//         if nums[mid]==target{
//             return mid
//         }
//         if target<nums[mid]{
//             max=mid-1
//         }else{
//             min=mid+1
//         }
//     }

//     return -1
// }