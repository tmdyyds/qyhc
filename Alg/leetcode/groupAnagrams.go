func TestAlg1010(t *testing.T) {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(s))
}


//面试题 10.02. 变位词组
//对strs 中的每一项进行排序，排序之后用做map的key，value是一个list
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	tmp := map[string][]string{}
	for _, item := range strs {
		s := strings.Split(item, "")
		sort.Strings(s)
		k := strings.Join(s, "") //key排序

		//排序后的字段如果已存在 则证明为变为词
		if _, ok := tmp[k]; ok {
			tmp[k] = append(tmp[k], item) //value为数组
		} else {
			tmp[k] = []string{item}
		}
	}

	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}