// func numSubarraysWithSum(nums []int, goal int) int {
//     sum:=0
//     diffCount:=make(map[int]int)
//     diffCount[0]=1
//     answ:=0
//     for i:=0; i<len(nums); i++{
//         sum+=nums[i]
//         answ+=diffCount[sum-goal]
//         diffCount[sum]++
//     }
//     return answ
// }

// // broot force check
//     // var totalGoal, sequence int

//     // for i:=0; i<len(nums); i++{
//     //     j:=i
//     //     for totalGoal<=goal && j<len(nums){
//     //         if nums[j]==1{
//     //             totalGoal++
//     //         }
//     //         if totalGoal == goal{
//     //             sequence++
//     //         }
//     //         j++
//     //     }
//     //     totalGoal=0
//     // }
//     // return sequence