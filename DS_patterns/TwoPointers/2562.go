func findTheArrayConcVal(nums []int) int64 {
	var i, j, result int
	j = len(nums) - 1
	for i < j {
		first := strconv.Itoa(nums[i])
		second := strconv.Itoa(nums[j])
		mixed:=first+second
        secondInt, _ := strconv.Atoi(mixed)
		result += secondInt 

		i++
		j--
	}
	if len(nums)%2 != 0 {
		result += nums[len(nums)/2]
	}
	return int64(result)
}
