// func lengthOfLongestSubstring(s string) int {
// 	var (
// 		count  int
// 		result int
// 		mp     = make(map[byte]bool)
// 		j      int
// 	)

// 	for i := 0; i < len(s); i++ {
//         // if i==6{
//         //     fmt.Println(mp, i,j ,count)
//         // }
// 		count++
//         // fmt.Println(mp, i,j, count)
// 		if mp[s[i]] {
// 			// fmt.Println(mp, count)
// 			if result < count-1 {
// 				// fmt.Println(i, j)
// 				result = count-1
// 			}

// 			for ; j < i; j++ {
// 				count--
// 				if s[i] == s[j] {
// 					break
// 				}
// 				delete(mp, s[j])
// 			}
// 			fmt.Println()
// 			continue
// 		}

// 		mp[s[i]] = true
// 	}

// 	if count > result {
// 		return count
// 	}
// 	return result
// }
