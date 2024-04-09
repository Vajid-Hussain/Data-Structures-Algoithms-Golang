// func minimumLengthEncoding(words []string) int {
//     var(
//         mp = make(map[string]struct{})
//         sum int
//     )

//     for _,val:=range words{
//         mp[val]=struct{}{}
//     }

//     // fmt.Println(mp)

//     for word:=range mp{
//         for i:=1; i<len(word); i++{
//             delete(mp,word[i:])
//         }
//     }

//     // fmt.Println(mp)

//     for word:=range mp{
//         sum+=len(word)
//     }
//     sum+=len(mp)
//     return sum
// }