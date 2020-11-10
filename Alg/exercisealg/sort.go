//桶排序



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

//start
//归并排序 时间复杂度O(nlogn) 空间复杂度O(n)
func MergeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	mergeSortRec(arr, 0, n-1)
}

func mergeSortRec(arr []int, start, end int) {
	//结束条件
	if start >= end {
		return
	}
	mid := (start + end) / 2 //中间位置

	//分治
	mergeSortRec(arr, start, mid)
	mergeSortRec(arr, mid+1, end)

	//真正排序合并处理
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start   //左边开始位置
	j := mid + 1 //右边开始位置
	k := 0

	for ; i <= mid && j <= end; k++ {
		if arr[i] > arr[j] {
			tmpArr[k] = arr[j]
			j++
		} else {
			tmpArr[k] = arr[i]
			i++
		}
	}

	//start,mid若有剩余
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}

	//mid+1,end 若有剩余
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	//tmpArr复制到arr
	copy(arr[start:end+1], tmpArr)
}

//end


//start
//快速排序 时间复杂度O(nlogn) 空间复杂度O(n)
func Quicksort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	quickSortRec(arr, 0, n-1)
}

func quickSortRec(arr []int, start, end int) {
	if start >= end {
		return
	}

	m := partition(arr, start, end)
	quickSortRec(arr, start, m-1)
	quickSortRec(arr, m+1, end)
}

//分区点位置&左边小于分区点右边大于分区点
func partition(arr []int, start, end int) int {
	i := start
	pivot := arr[end] //随机选择 这里选择最后一个

	for j := start; j < end; j++ {
		//小于pivot的在左边
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}

			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}

//end