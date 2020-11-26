func TestAlg1010(t *testing.T) {
	s := "asba"
	fmt.Println(validPalindrome(s))
}

//680. 验证回文字符串 Ⅱ
func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	//验证回文串
	isValidPalindrome := func(slow, hight int, s string) bool {
		for slow < hight {
			if s[slow] == s[hight] {
				slow++
				hight--
				continue
			}

			return false
		}

		return true
	}

	//思路 删除i或者j的位置 判断剩余是否是回文 若有一个是回文 则直接return
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			l := isValidPalindrome(i+1, j, s) //删除i位置的字符 判断剩余是否回文
			r := isValidPalindrome(i, j-1, s) //同理

			return l || r
		}
	}

	return true
}
