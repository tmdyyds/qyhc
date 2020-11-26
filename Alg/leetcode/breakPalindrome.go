func TestAlg1102(t *testing.T) {
	s := "abctcba"
	fmt.Println(breakPalindrome(s))
}

//1328. 破坏回文串
// 1.字符串长度小于等于1，返回空串。——单个字母字符串是回文串。
// 2.由左到右找到第一非a字符，将该字符替换为a。
// 下面两种情况，需要将最后一个a替换为b。
// 2.1 找到字符串中间位置全部都是a。
// 2.2 奇数长度字符串，只有最中间一个为非a字母。——最中间字母替换，还是回文串。
func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n <= 1 {
		return ""
	}

	p := -1
	//找到不是a的位置
	for i := 0; i < (n-1)/2+1; i++ {
		if palindrome[i] != 'a' {
			p = i
			break
		}
	}

	r := ""

	//p=-1 证明全是a, p = n/2 只有最中间一个为非a字母
	if p == -1 || p == n/2 {
		r = palindrome[0:n-1] + "b"
	} else {
		//palindrome[0:p]左开右闭 故此处p被抛弃
		r = palindrome[0:p] + "a" + palindrome[p+1:]
	}

	return r
}