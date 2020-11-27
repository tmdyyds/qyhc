func TestAlg1102(t *testing.T) {
	ts := 0.1
	fmt.Println(printBin(ts))
}

//面试题 05.02. 二进制数转字符串
func printBin(num float64) string {
	r := "0."
	i := 0
	for ; i < 32; i++ {
		num = num * 2

		//为1则证明没有小数 退出for
		if num == 1 {
			r += "1"
			break
		}

		//结果加1 继续for
		if num > 1 {
			r += "1"
			num = num - 1
			continue
		}

		//小于0 直接加0
		r += "0"
	}

	//是否超过32位
	if i >= 32 {
		return "ERROR"
	}

	return r
}