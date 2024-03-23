// func dividePlayers(skill []int) int64 {
// 	var (
// 		sum     int
// 		players int
//         i int
//         j = len(skill)-1
//         result int
// 	)

// 	for _, val := range skill {
// 		sum += val
// 		players++
// 	}

// 	if sum%(players/2) != 0 {
// 		return -1
// 	}

//     teamSkill:= sum / (players/2)
//     slices.Sort(skill)

//     for i<j{
//         if skill[i] + skill[j]==teamSkill{
//             result+= skill[i] *skill[j]
//         }else{
//             return -1
//         }
//         i++
//         j--
//     }
// 	return int64(result)
// }