func TestAlg1010(t *testing.T) {
	array1 := []int{4, 1, 2, 1, 1, 2} //11
	array2 := []int{3, 6, 3, 3}       //15
	fmt.Println(findSwapValues(array1, array2))
}


//面试题 16.21. 交换和
//sum1-a+b = sum2-b+a，即sum1-sum2=2(a-b)，
//显然，A、B两个数组的和之差应该是偶数，否则肯定不存在符合条件的a、b。
func findSwapValues(array1 []int, array2 []int) []int {
	res := []int{}
	m := make(map[int]int)

	sum1 := 0
	for _, v := range array1 {
		sum1 += v
	}

	sum2 := 0
	for k, v := range array2 {
		m[v] = k
		sum2 += v

	}

	if (sum2+sum1)%2 == 1 {
		return res
	}

	diff := (sum1 - sum2) / 2
	for i := 0; i < len(array1); i++ {
		t := array1[i] - diff
		if _, ok := m[t]; ok {
			return []int{array1[i], t}
		}
	}

	return res
}