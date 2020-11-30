func TestAlg1102(t *testing.T) {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}

//1295. 统计位数为偶数的数字
func findNumbers(nums []int) int {
	//效率优于下面方式
	nc findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		cur := 0 //位数
		for v > 0 {
			//除10
			v /= 10
			cur++
		}

		//是否是偶数
		if cur%2 == 0 {
			res++
		}
	}
	return res

/*	n := 0
	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			n++
		}
	}

	return n*/
}