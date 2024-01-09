// package main

// import "fmt"

// func sortSentence(s string) string {
// 	var word string
// 	var mp = make(map[int]string)
// 	i := 0
// 	var count int

// 	for i < len(s) {
// 		if s[i] <= 57 && s[i] >= 48 {
// 			count++
// 			mp[int(s[i]-'0')] = word
// 			word = ""
// 			i += 2
// 			continue
// 		}
// 		word += string(s[i])
// 		i++
// 	}

// 	fmt.Println("**", mp, word)

// 	i = 1
// 	for i <= count {
// 		word += mp[i]
// 		if i != count {
// 			word += " "
// 		}
// 		i++
// 	}

// 	fmt.Println("@@", word)
// 	return s
// }

// func main() {
// 	sortSentence("Myself2 Me1 I4 and3")
// }
