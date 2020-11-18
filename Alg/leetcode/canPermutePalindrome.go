func TestAlg1010(t *testing.T) {
	s := "abc"
	fmt.Println(canPermutePalindrome(s))
}


//面试题 01.04. 回文排列
func canPermutePalindrome(s string) bool {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		//统计每个字符出现次数
		m[string(s[i])]++
	}

	oddCnt := 0
	for k, _ := range m {
		//统计出现奇数字符的个数
		if m[k]%2 == 1 {
			oddCnt++
		}
	}

	//若出现多于一个情况下 则证明没有回文可能
	if oddCnt > 1 {
		return false
	}

	return true
}