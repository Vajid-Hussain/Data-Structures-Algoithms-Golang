package problems

func FindDifference(nums1 []int, nums2 []int) [][]int {
	var map1 = make(map[int]int)
	var map2 = make(map[int]int)
	var arr1 []int
	var arr2 []int
	var result [][]int

	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]]++
	}

	for key := range map1 {

		_, ok := map2[key]
		if !ok {
			arr1 = append(arr1, key)
		}
	}

	for key := range map2 {

		_, ok := map1[key]
		if !ok {
			arr2 = append(arr2, key)
		}
	}

	result = append(result, arr1)
	result = append(result, arr2)
	return result
}
