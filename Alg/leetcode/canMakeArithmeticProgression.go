func TestAlg1010(t *testing.T) {
	A := []int{1, 5, 3}
	fmt.Println(canMakeArithmeticProgression(A))
}

//1502. 判断能否形成等差数列
func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr) - 1
	//数组先排序 快排
	qucikSortl(arr, 0, n)

	fixVal := arr[1] - arr[0]
	for i := 2; i <= n; i++ {
		if arr[i]-arr[i-1] != fixVal {
			return false
		}
	}

	return true
}

//排序
func qucikSortl(arr []int, start, end int) {
	if start >= end {
		return
	}
	p := partitionl(arr, start, end)
	qucikSortl(arr, start, p-1)
	qucikSortl(arr, p+1, end)
}

//分区点
func partitionl(arr []int, start, end int) int {
	i := start
	p := arr[end]

	for j := start; j < end; j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}
