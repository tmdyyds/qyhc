func TestAlg1010(t *testing.T) {
	A := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(A))
}

//922. 按奇偶排序数组 II

//解法1
func sortArrayByParityII(A []int) []int {
	n := len(A)
	res := make([]int, n)
	x := 0 //偶数key
	y := 1 //奇数key
	for _, v := range A {
		if v%2 == 0 {
			res[x] = v
			x += 2
		} else {
			res[y] = v
			y += 2
		}
	}

	return res
}

//解法2
func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		//仅检查奇数 偶数直接循环
		if A[i]%2 != 0 {
			//如果有偶数位置为奇数 则一定有一个奇数位置为偶数
			//若为偶数 则直接互换i 和j的位置
			for A[j]%2 == 1 {
				j += 2
			}

			A[i], A[j] = A[j], A[i]
			j += 2
		}
	}

	return A
}