// 53. 最大子序和
func maxSubArray(nums []int) int {
	r := nums[0]
	maxSum := 0

	for i := 0; i < len(nums); i++ {
		if maxSum > 0 {
			maxSum += nums[i]
		} else {
			maxSum = nums[i]
		}

		r = compareInt(r, maxSum)
	}

	return r
}

func compareInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
