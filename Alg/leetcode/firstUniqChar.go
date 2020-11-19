func TestAlg1010(t *testing.T) {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}

//387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for i := 0; i < len(s); i++ {
		m[int32(s[i])]++ //记录次数
	}

	for k, v := range s {
		if m[v] == 1 {
			return k
		}
	}

	return -1

	//解法2
	//在数组中记录每个字母的最后一次出现的所在索引。然后再通过一次循环，比较各个字母第一次出现的索引是否为最后一次的索引
	//第一次循环容易想到 第二部太骚操作了
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}

	for i, k := range s {
		//寻找第一次出现的位置是否与最后一次出现位置一致
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}

	return -1
}