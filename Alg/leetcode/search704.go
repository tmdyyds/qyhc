//704. 二分查找
func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return -1
}
