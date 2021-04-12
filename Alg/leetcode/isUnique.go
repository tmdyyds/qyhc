//面试题 01.01. 判定字符是否唯一
func isUnique(astr string) bool {
	m := make(map[int32]bool)
	for _, v := range astr {
		if m[v] {
			return false
		}

		m[v] = true
	}
	return true
}