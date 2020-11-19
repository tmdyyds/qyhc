func TestAlg1010(t *testing.T) {
	haystack := "hesadfasdfasfll"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

//28. 实现 strStr()
func strStr(haystack string, needle string) int {
	hn := len(haystack)
	nn := len(needle)
	if nn == 0 {
		return 0
	}

	if nn > hn {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		//对比needle开始位置 在haystack中查找 再寻找过程中 注意条件溢出情况
		//在找到首字母与needle一样时 直接截取i:i+nn长度的字符判断与needle是否相同
		if haystack[i] == needle[0] && hn-i >= nn && haystack[i:i+nn] == needle {
			return i
		}
	}

	return -1
}