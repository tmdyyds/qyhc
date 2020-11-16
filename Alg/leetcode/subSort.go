func TestAlg1010(t *testing.T) {
	slice := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	subSort(slice)
}

//面试题 16.16. 部分排序
func subSort(array []int) []int {
	max := math.MinInt64 //最小值
	min := math.MaxInt64
	l := -1 //找不到时 默认-1
	r := -1
	//1,2,4,7,10,11,7,12,6,7,16,18,19
	//找到最右边的最大值
	for i := 0; i < len(array); i++ {
		if array[i] < max {
			r = i
		} else {
			max = array[i]
		}
	}

	//找到最左边最小值
	for j := len(array) - 1; j >= 0; j-- {
		if array[j] > min {
			l = j
		} else {
			min = array[j]
		}
	}

	return []int{l, r}
}