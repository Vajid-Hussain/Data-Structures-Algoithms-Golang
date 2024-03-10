// func matchPlayersAndTrainers(players []int, trainers []int) int {
// 	var result int
// 	var pla, tra = len(players), len(trainers)
// 	slices.Sort(players)
// 	slices.Sort(trainers)
// 	i, j := 0, 0
// 	for i < pla && j < tra {
// 		if players[i] <= trainers[j] {
// 			result++
// 			i++
// 			j++
// 		} else {
// 			j++
// 		}
// 	}
// 	return result
// }
