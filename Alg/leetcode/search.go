func TestAlg1201(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

//33. 搜索旋转排序数组
func search(nums []int, target int) int {
	n := len(nums)
	low := 0
	high := n - 1

	for low <= high { //条件注意
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}

		//主要思路: mid=6 把数组分隔成有序和无序数组[4, 5, 6] 和 [7, 0, 1, 2]
		//[low, mid] 有序数组, [mid+1, high] 无序数组
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			//[mid+1, high]
			if nums[mid] < target && target <= nums[n-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}