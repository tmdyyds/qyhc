func TestAlg1010(t *testing.T) {
	A := []int{1, 1, 2, 2, 2, 3}
	fmt.Println(frequencySort(A))
}

//1636. 按照频率将数组升序排序
func frequencySort(nums []int) []int {
	hash := make(map[int]int)
	//hash出值频率：次数
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	//排序
	sort.Slice(nums, func(i, j int) bool {
		return hash[nums[i]] < hash[nums[j]] || hash[nums[i]] == hash[nums[j]] && nums[i] > nums[j]
	})

	return nums
}