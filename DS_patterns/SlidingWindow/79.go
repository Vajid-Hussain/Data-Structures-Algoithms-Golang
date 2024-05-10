// func minWindow(s string, t string) string {
// 	var (
// 		mp1        = make(map[byte]int)
// 		mp2        = make(map[byte]int)
// 		rStart     int
// 		rEnd       int = 100000
// 		start      int
// 		have, need int
//         result string
// 	)
// 	for i := 0; i < len(t); i++ {
// 		mp1[t[i]] = 0
// 		mp2[t[i]]++
// 	}
// 	need = len(mp2)

// 	for i := 0; i < len(s); i++ {
// 		mp1[s[i]]++
// 		if mp1[s[i]] == mp2[s[i]] {
// 			have++
// 		}

// 		if have == need {
// 			if rEnd-rStart > i-start {
// 				rEnd = i
// 				rStart = start
// 			}
// 		}

// 		for have == need {
// 			_, exist := mp1[s[start]]
// 			if exist {
// 				mp1[s[start]]--
// 				if mp1[s[start]] < mp2[s[start]] {
// 					have--
// 				}
// 			}
// 			start++

// 			if have == need {
// 				if rEnd-rStart > i-start {
// 					rEnd = i
// 					rStart = start
// 				}
// 			}
// 		}
// 	}

//     if rEnd==100000{
//         return ""
//     }

// 	for i:=rStart; i<= rEnd; i++{
//         result+= string(s[i])
//     }
// 	return result
// }