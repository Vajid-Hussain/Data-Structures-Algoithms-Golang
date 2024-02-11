// func search(nums []int, target int) bool {
//     var i int
//     for i=0; i<len(nums)-1; i++{
//         if nums[i+1] < nums[i]{
//         break
//         }
//     }

//     var min,max int

//     if target <= nums[i] && target>= nums[0]{
//         min=0
//         max=i
//     }else{
//         min=i+1
//         max=len(nums)-1
//     }

//     for min<= max{
//         mid:=(min+max)/2
//         fmt.Println(mid)
//         if nums[mid]==target{
//             return true
//         }
//         if target> nums[mid]{
//             min=mid+1
//         }else{
//             max=mid-1
//         }
//     }
//     return false
// }