// package main

// import "fmt"

// func findTheLongestBalancedSubstring(s string) int {

// 	var statement1, statement2 = true, true
// 	var i int
// 	var count1, count2 int
// 	var result int

// 	for i < len(s) {

// 		for statement1 && i < len(s) {
// 			if count1 == 0 {
// 				count1 += 1
// 				i++
// 				continue
// 			}
// 			if s[i] == s[i-1] {
// 				count1 += 1
// 				i++
// 			} else {
// 				statement1 = false
// 				statement2 = true
// 			}
// 		}

// 		for statement2 && i < len(s) {
// 			fmt.Println("$$")
// 			if count2 == 0 {
// 				count2 += 1
// 				i++
// 				continue
// 			}
// 			if s[i] == s[i-1] {
// 				count2 += 1
// 				i++
// 			} else {
// 				statement2 = false
// 				statement1 = true
// 			}
// 		}
// 		fmt.Println("**", count1, count2)
// 		if (count1 + count2) > result {
// 			if count1 == count2 {
// 				result = count1 + count2
// 			} else if count1 > count2 {
// 				if (count2 * 2) > result {
// 					result = count2 * 2
// 				}
// 			} else if count2 > count1 {
// 				if count1*2 > result {
// 					result = count1 * 2
// 				}
// 			}
// 		}
// 		count1, count2 = 0, 0
// 	}
// 	return result
// }

// func main() {
// 	fmt.Println(findTheLongestBalancedSubstring("10"))
// }
