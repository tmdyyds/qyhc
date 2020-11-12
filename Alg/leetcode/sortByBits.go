func TestAlg1010(t *testing.T) {
	A := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(A))
}

//1356. 根据数字二进制下 1 的数目排序
func sortByBits(arr []int) []int {

	//排序
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		cx, cy := onesCount(x), onesCount(y)

		//关键点在这里 两种情况 1 二进制1的个数 2 个数相同 则比较值大小
		return cx < cy || cx == cy && x < y
	})

	return arr
}

//这种效率低于下面方式
func onesCount(x int) (c int) {
	for ; x > 0; x /= 2 {
		c += x % 2
	}

	return
}

//位运算 x = x & (x - 1)
func onesCount(x int) (c int) {
	c = 0
	for x > 0 {
		c++
		x = x & (x - 1)
	}

	return
}
