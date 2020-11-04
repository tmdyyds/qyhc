//冒泡
test := []int{4, 5, 6, 3, 2, 1}
bubbleSort(test)
fmt.Println(test)

//插入排序
test := []int{4, 5, 6, 3, 2, 1}
InsertionSort(test)
fmt.Println(test)

//选择排序
test : = []int{4, 5, 6, 3, 2, 1}
SelectionSort(test)
fmt.Println(test)

//选择排序 时间复杂度O(N^2) 空间复杂度O(1) 原地排序
func SelectionSort(r []int) {
	n := len(r)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minVal := i //最小值下标(最下值)
		for j := i + 1; j < n; j++ {
			if r[j] < r[minVal] {
				minVal = j
			}
		}

		//交换
		r[i], r[minVal] = r[minVal], r[i]
	}
}


//插入排序 时间复杂度O(N^2) 空间复杂度O(1) 原地排序
func InsertionSort(r []int) {
	n := len(r)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		val := r[i]
		j := i - 1 //为了下面插入位置

		//这里for循环仅为了在已排好序的数据中寻找位置和移动数据
		for ; j >= 0; j-- {
			if r[j] > val {
				r[j+1] = r[j] //数据移动,每次把当前j位置移往后移动一位
			} else {
				break
			}
		}

		//插入数据
		r[j+1] = val
	}
}

//原地排序特指空间复杂度O(1)的排序算法
//冒泡排序 时间复杂度O(N^2) 空间复杂度O(1)
func BubbleSort(r []int) {
	n := len(r)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := false //提前结束冒泡标识

		//n-i减去已排好序
		//n-1跟下一个元素对比(j+1) 不用直接到最后
		for j := 0; j < n-i-1; j++ {
			//从小到大排序
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]

				flag = true
			}
		}

		if !flag {
			break
		}
	}
}