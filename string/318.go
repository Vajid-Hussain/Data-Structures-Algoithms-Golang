// func maxProduct(words []string) int {
// 	var (
// 		collection = make(map[int]map[rune]struct{})
// 		result     int
// 		temp       int
// 	)

// 	for index, word := range words {
// 		mp := make(map[rune]struct{})
// 		for _, text := range word {
// 			mp[text] = struct{}{}
// 		}
// 		collection[index] = mp
// 	}
// 	// fmt.Println(collection)

// 	for i := 0; i < len(words); i++ {
// 		for j := 0; j < len(words); j++ {
// 			temp = 0
// 			if i != j {
// 				for key, _ := range collection[i] {
// 					_, exist := collection[j][key]
// 					if exist {
// 						temp = 1
// 						break
// 					}
// 				}
// 				if temp == 0 {
//                     // fmt.Println(collection[i], collection[j], i, j)
// 					// result += len(words[i]) * len(words[j])
//                     if result<len(words[i]) * len(words[j]){
//                         result= len(words[i]) * len(words[j])
//                     }
// 				}
// 			}
// 		}
// 	}
// 	return result
// }