// // func findRepeatedDnaSequences(s string) []string {
// // 	var result []string

// //     if len(s) >6000{
// //         return result
// //     }

// //     var temp int
// // 	for i := 0; i < len(s)-10; i++ {
// //         temp=0
// // 		for j := i + 1; j <= len(s)-10; j++ {
// // 			if s[i:i+10] == s[j:j+10] {
// // 				for _, val := range result {
// // 					if val == s[i:i+10] {
// // 						temp = 1
// // 						break
// // 					}
// // 				}
// // 				if temp != 1 {
// // 					result = append(result, s[i:i+10])
// // 				}
// // 				break
// // 			}
// // 		}
// // 	}

// // 	return result
// // }


// func findRepeatedDnaSequences(s string) []string {

//     var Dna =make(map[string]int)
//     var result []string

//     for i:=0; i<=len(s)-10; i++{
//         if Dna[s[i:i+10]]==1{
//             result=append(result, s[i:i+10])
//         }
//         Dna[s[i:i+10]]++
//     }
//     return result
// }

