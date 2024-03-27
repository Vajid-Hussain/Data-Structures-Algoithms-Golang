// func maxVowels(s string, k int) int {
// 	var result, current int
// 	var mp = map[string]struct{}{
// 		"a": struct{}{},
// 		"e": struct{}{},
// 		"i": struct{}{},
// 		"o": struct{}{},
// 		"u": struct{}{},
// 	}
// 	for i := 0; i < len(s); i++ {
// 		if i < k-1 {
// 			if _, ok := mp[string(s[i])]; ok {
// 				current++
// 			}
//             continue
// 		}
        
// 		if _, ok := mp[string(s[i])]; ok {
// 			current++
// 		}

//         if current > result{
//             result =current
//         }

// 		if _, ok := mp[string(s[i-(k-1)])]; ok {
// 			current--
// 		}
// 	}

// 	return result
// }
