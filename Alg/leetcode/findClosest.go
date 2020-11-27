func TestAlg1102(t *testing.T) {
	dictionary := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"
	fmt.Println(findClosest(dictionary, word1, word2))
}

//面试题 17.11. 单词距离
//两个指针 找位置
func findClosest(words []string, word1 string, word2 string) int {
	w1, w2, min := 0, 0, 0
	for k, _ := range words {
		//寻找位置
		if words[k] == word1 {
			w1 = k
			continue
		}

		if words[k] == word2 {
			w2 = k
			continue
		}

		//计算间距 1 min=0直接赋值， 2判断最新位置和min大小
		if w1 > 0 && w2 > 0 && (min == 0 || abs(w1-w2) < min) {
			min = abs(w1 - w2)
			continue
		}
	}

	return min
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}