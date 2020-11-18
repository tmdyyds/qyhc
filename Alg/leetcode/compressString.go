func TestAlg1010(t *testing.T) {
	s := "abc" //f1a4b1c5a3
	fmt.Println(compressString(s))
}

//面试题 01.06. 字符串压缩
func compressString(S string) string {
	if len(S) == 0 {
		return S
	}

	var b strings.Builder

	tmpStr := S[0]
	currentCnt := 1
	str := S + "*" //这里结尾加一个哨兵 不容易想到
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] { //这里相当于和str[0]比较
			currentCnt++
			continue
		}

		b.WriteByte(tmpStr)
		b.WriteString(strconv.Itoa(currentCnt))
		tmpStr = str[i]
		currentCnt = 1
	}

	//题意加如此判断
	if b.Len() >= len(S) {
		return S
	}

	return b.String()
}