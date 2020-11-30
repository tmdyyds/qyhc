func TestAlg1102(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

//剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	//优于解法3
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] >= 2 {
			return v
		}
	}

	return 0

	//解法3
/*	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return 0*/
}