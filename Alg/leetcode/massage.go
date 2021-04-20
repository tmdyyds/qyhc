// 面试题 17.16. 按摩师
func massage(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[l-1]
}