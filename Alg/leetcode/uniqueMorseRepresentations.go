func TestAlg1010(t *testing.T) {
	s := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(s))
}


//804. 唯一摩尔斯密码词
//leetcode题意不太好理解,最终它要的返回值是不重复摩尔斯密码词的数量
func uniqueMorseRepresentations(words []string) int {
	letter26 := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool)
	for _, v := range words {
		s := ""
		for _, vv := range v {
			s += letter26[vv-'a']
		}

		if _, ok := m[s]; !ok {
			m[s] = true
		}
	}

	return len(m)
}