func TestAlg1010(t *testing.T) {
	s := "abcd"
	fmt.Println(repeatedSubstringPattern(s))
}

//459. 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	//如果s中包含重复的子字符串，那么说明s中至少包含两个子字符串，
	//s+s至少包含4个字串，前后各去掉一位，查找s是否在新构建的字符串中
	s1 := s + s
	s2 := s1[1 : len(s1)-1]

	return strings.Contains(s2, s)
}