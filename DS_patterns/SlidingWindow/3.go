// func lengthOfLongestSubstring(s string) int {
// 	var(
// 	   mp=make(map[byte]struct{})
// 	   start int
// 	   end int
// 	   result int
// 	)   
   
// 	for end<len(s){
// 	   _, exist := mp[s[end]]
// 	   if !exist{
// 		   // fmt.Println(end)
// 		   mp[s[end]]=struct{}{}
// 		   end++
// 	   }
   
// 	   if exist || end==len(s) {
// 		   // fmt.Println(mp, start,end)
// 		   if result < end-start{
// 			   result=end-start
// 		   }
   
// 		   if end==len(s){
// 			   return result
// 		   }
// 		   clear(mp, s[end], s, &start)
// 	   }
// 	}
// 	   return result
//    }
   
//    func clear(mp map[byte]struct{}, target byte, s string, start *int) {
// 	   for s[*start] != target{
// 		   delete(mp, s[*start])
// 		   *start++
// 	   }
// 	   delete(mp, s[*start])
// 	   *start++
//    } 