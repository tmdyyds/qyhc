// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}

	length := 0
	tmp := 0
	for i := 1; i < l; i++ {
		for j := tmp; j < i; j++ {
			if s[i] == s[j] {
				tmp = j + 1
			}
		}

		if length < (i + 1 - tmp) {
			length = i + 1 - tmp
		}
	}

	return length
}
