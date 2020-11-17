func TestAlg1010(t *testing.T) {
	s1 := "               "
	k := 5
	fmt.Println(replaceSpaces(s1, k))
}

//面试题 01.03. URL化
//个人开始想法把字符按着空格分隔,但是这样增加复杂度 理论应该是可以处理
func replaceSpaces(S string, length int) string {
	//简单方便
	return strings.Replace(S[:length], " ", "%20", -1)

	//解法2
	// 初始化一个strings.Builder对象
	b := strings.Builder{}

	// 直接扩容到目标大小
	b.Grow(length)

	// 取真实长度，然后遍历
	for s, i := []rune(S[:length]), 0; i < len(s); i++ {
		// 遇到空格就写"%20"
		if s[i] == ' ' {
			b.WriteString("%20")
			continue
		}
		// 直接写rune
		b.WriteRune(s[i])
	}

	// 转换成String返回
	return b.String()
}