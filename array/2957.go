// func removeAlmostEqualCharacters(word string) int {
//     bytes:=[]byte(word)
//     var count int
//     for i:=0; i<len(bytes)-1; i++{
//         if bytes[i]== bytes[i+1] || (bytes[i]-1) == bytes[i+1] || (bytes[i]+1) == bytes[i+1]{
//             bytes[i+1]+=29
//             count++
//         }
//     }
//     return count
// }