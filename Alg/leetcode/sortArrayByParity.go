//905. 按奇偶排序数组
func sortArrayByParity(A []int) []int {
	ret := []int{}
	for _, n := range A {
		if n%2 == 0 {
			ret = append([]int{n}, ret...)
			continue
		}

		ret = append(ret, n)
	}

	return ret



/*	n := len(A)
	j := n - 1
	i := 0
	for i <= j {
		if A[i]%2 == 0 {
			i++
			continue
		}

		A[i], A[j] = A[j], A[i]
		j--
	}

	return A*/


/*	n := len(A)
	j := n - 1
	i := 0
	for i <= j {
		if A[i]%2 == 0 {
			i++
			continue
		}

		if A[j]%2 == 0 {
			A[i], A[j] = A[j], A[i]
		}

		j--
	}

	return A*/
}
