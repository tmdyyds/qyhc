func TestAlg1010(t *testing.T) {
	s1 := "a good   example"
	fmt.Println(reverseWords(s1))
}

//剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	//解法1
	r := ""

	//分隔字符
	st := strings.Split(strings.TrimSpace(s), " ")

	//倒序循环
	for i := len(st) - 1; i >= 0; i-- {
		if len(st[i]) == 0 { //注意这里 用len原因在于字符串中多个空格问题 例如:ab   sd sadf
			continue
		}

		r += " " + st[i]
	}

	return strings.TrimSpace(r)

	//解法2
	var res []string
	st := strings.TrimSpace(s)
	i := len(st) - 1
	j := i
	for i >= 0 {
		for i >= 0 && st[i] != 32 { //搜索首个空格位置
			i-- //i--从后往前走
		}

		//找到第一个为空的值时 把前面经过的字符放到Res中
		//注意 i+1已包括空格
		res = append(res, string(st[i+1:j+1]))

		for i >= 0 && st[i] == 32 {
			i-- //若为空格 则继续往前搜索
		}

		//前面不是空格时 i = j,继续上面流程
		j = i
	}

	return strings.Join(res, " ")
}