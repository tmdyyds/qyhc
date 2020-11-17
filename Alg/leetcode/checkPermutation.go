func TestAlg1010(t *testing.T) {
	s1 := "abc"
	s2 := "bca"
	fmt.Println(CheckPermutation(s1, s2))
}


//面试题 01.02. 判定是否互为字符重排
func checkPermutation(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	//长度必须一致
	if n1 != n2 {
		return false
	}

	m := make(map[string]int)

	//s1中所有字符出现次数累计
	for _, v := range s1 {
		m[string(v)]++
	}

	for _, vv := range s2 {
		if _, ok := m[string(vv)]; ok {
			m[string(vv)]--
		} else {
			//若s2中字符有不同于s1中 则直接返回false
			return false
		}
	}

	//最后需要判断一下 m中是否全部为0 否则证明s1中数据必定有和s2中不一致
	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true


	//优化
	n1 := len(s1)
	n2 := len(s2)
	//长度必须一致
	if n1 != n2 {
		return false
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for i := 0; i < n1; i++ {
		m1[string(s1[i])]++
		m2[string(s2[i])]++
	}

	for _, v := range s1 {
		if m1[string(v)] != m2[string(v)] {
			return false
		}
	}

	return true
}