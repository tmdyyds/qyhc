func TestAlg1102(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

//剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {

	//最优
	for i := 0; i < len(nums); i++ {
		//索引值等于索引时 寻找下一位
		if i == nums[i] {
			continue
		}

		//索引 nums[i] 处的值也为 nums[i]，即找到一组相同值，返回 nums[i] 即可
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		//执行交换操作，目的是为了使索引与值一一对应，即索引 0 的值为 0，索引 1 的值为 1
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return -1

	//解法2
	//优于解法3
/*	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] >= 2 {
			return v
		}
	}

	return 0*/

	//解法3
/*	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return 0*/
}