func TestAlg1010(t *testing.T) {
	test := []int{3, 6, 2, 1, 5, 4}
	//test := []int{3, 2, 3, 1, 2, 4, 5, 5, 6} k := 4
	k := 2
	fmt.Println(findKthLargest(test, k))
}

//215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano()) //按时间随机种子
	n := len(nums)

	//第k大索引为n-k
	return quickSortRecAlg(nums, 0, n-1, n-k)
}

func quickSortRecAlg(nums []int, start, end, index int) int {
	m := randomPartition(nums, start, end)

	if m == index {
		return nums[m]
	} else if m < index {
		return quickSortRecAlg(nums, m+1, end, index)
	} else {
		return quickSortRecAlg(nums, start, m-1, index)
	}
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l //生成start-end之间一个随机数 或者直接用最后一个(若倒序数字时间复杂度为O(n^2))
	a[i], a[r] = a[r], a[i]

	return partitionAlg(a, l, r)
}

//分区点
func partitionAlg(nums []int, start, end int) int {
	i := start
	pivot := nums[end] //随机选择 这里选择最后一个

	for j := start; j < end; j++ {
		//小于pivot的在左边
		if nums[j] <= pivot {
			if !(i == j) {
				// 交换位置
				nums[i], nums[j] = nums[j], nums[i]
			}

			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}