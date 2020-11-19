func TestAlg1010(t *testing.T) {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s, k))
}

//剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]

	//第二种 不利用切片
	var str string
	var str2 string
	for i := 0; i < len(s); i++ {
		if i < n {
			str += string(s[i])
		} else {
			str2 += string(s[i])
		}
	}

	return str2 + str
}






