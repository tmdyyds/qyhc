func TestAlg1010(t *testing.T) {
	A := []int{1, 2, 1}
	fmt.Println(largestPerimeter(A))
}

//976. 三角形的最大周长
func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}