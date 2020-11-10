func TestAlg1010(t *testing.T) {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	fmt.Println(restoreString(s, indices))
}

//1528. 重新排列字符串
func restoreString(s string, indices []int) string {
	n := len(s)
	t := make([]string, n)
	for i := 0; i < n; i++ {
		t[indices[i]] = string(s[i])
	}

	return strings.Join(t, "")
}