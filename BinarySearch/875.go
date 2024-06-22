// func minEatingSpeed(piles []int, h int) int {
//     var (
//         min int = 1
//         max int 
//         mid int
//         result int = math.MaxInt
//         hour int
//     )   

//     for i:=0; i<len(piles); i++{
//         if max < piles[i]{
//             max = piles[i]
//         }
//     }

//     for min <= max{
//         hour = 0
//         mid = (min + max) /2

//         for i:=0; i<len(piles); i++{
//             hour+= piles[i] / mid
//             if piles[i] % mid !=0{
//                 hour++
//             }
//         }
        
//         if hour <= h {
//             max = mid-1
//         }else{
//             min = mid+ 1
//         }

//         if result > mid && hour <= h{
//             result = mid
//         }
//     }

//     return result 
// }