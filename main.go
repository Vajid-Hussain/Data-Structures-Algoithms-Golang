package main

import (
	problems "Leetcode/Problems"
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3}
	arr := []int{2, 4, 6}

	result := problems.FindDifference(nums, arr)

	fmt.Println(result)
}
