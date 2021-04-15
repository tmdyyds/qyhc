//35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
