// 1207. 独一无二的出现次数

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		//arr的值为k,出现次数为v 重新赋值map
		m[v]++
	}

	mk := make(map[int]bool)
	for _, vv := range m {

		//vv为出现次数 已存在则有重复次数 返回false
		if _, ok := mk[vv]; ok {
			return false
		}

		//不存在 则赋值
		mk[vv] = true
	}

	return true
}