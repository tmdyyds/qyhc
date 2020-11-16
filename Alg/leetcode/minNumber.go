func TestAlg1010(t *testing.T) {
	slice := []int{3, 30, 34, 5, 9}
	minNumber(slice)
}

//剑指 Offer 45. 把数组排成最小的数
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})

	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}

	return res.String()
}

/*a b
3 30  a+b=330 b+a=303*/
//接下来对字符串进行排序,自定义对比规则 a+b > b+a 则 a > b(排序规则)
func compareNumber(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	//如果a+b < b+a 则希望b 在前面
	if str1 < str2 {
		return true
	}
	return false
}