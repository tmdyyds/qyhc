//1. 两数之和
func twoSum(nums []int, target int) []int {
	r := []int{}
	m := make(map[int]int)

	for k, v := range nums {
		if kk, bool := m[target-v]; bool {
			r = append(r, kk)
			r = append(r, k)

			break
		}

		m[v] = k
	}

	return r
}