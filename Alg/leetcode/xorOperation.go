func TestAlg1102(t *testing.T) {
	n := 5
	start := 0
	fmt.Println(xorOperation(n, start))
}

//1486. 数组异或操作
func xorOperation(n int, start int) int {
	num := 0
	for i := 0; i < n; i++ {
		tn := start + 2*i
		num ^= tn
	}

	return num
}