func TestAlg1010(t *testing.T) {
	s := "Hello"
	fmt.Println(toLowerCase(s))
}


//709. 转换成小写字母
func toLowerCase(str string) string {
	r := ""
	for _, v := range str {
		if v >= 65 && v <= 90 {
			r += string(v + 32)
			continue
		}

		r += string(v)
	}

	return r

	//解法2 位运算 骚气
	/*大写变小写、小写变大写：字符 ^= 32;
	大写变小写、小写变小写：字符 |= 32;
	大写变大写、小写变大写：字符 &= 33;*/
	r := ""
	for _, v := range str {
		r += string(v | 32)
	}

	return r
}