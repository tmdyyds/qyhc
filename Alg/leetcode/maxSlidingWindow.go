//239. 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 0 {
		return nil
	}
	var (
		windows []int
		res     []int
	)
	//遍历数组
	for i, v := range nums {
		//i-k = 当窗口满时的最小下标
		//i-k >= 窗口最左边的数，则代表窗口已经满了
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}

		for {
			//如果窗口最后一个小于v，则去掉。为了维护了每次窗口最左边值最大
			if len(windows) > 0 && nums[windows[len(windows)-1]] <= v {
				windows = windows[:len(windows)-1]
			} else {
				break
			}
		}

		windows = append(windows, i)

		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}

	return res
}





// func maxSlidingWindow(nums []int, k int) []int {
// 	res := []int{}
// 	tmp := []int{}

// 	for i := 0; i < len(nums); i++ {
// 		if i < k-1 {
// 			tmp = append(tmp, nums[i])
// 		} else {
// 			tmp = append(tmp, nums[i])
// 			res = append(res, getMax(tmp))
// 			tmp = tmp[1:]
// 		}
// 	}

// 	return res
// }

// func getMax(n []int) int {
// 	max := math.MinInt32
// 	for _, v := range n {
// 		if max < v {
// 			max = v
// 		}
// 	}

// 	return max
// }
