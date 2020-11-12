func TestAlg1010(t *testing.T) {
	A := []int{4, 4, 7, 6, 7}
	fmt.Println(minSubsequence(A))
}

//1403. 非递增顺序的最小子序列
/*1.先排序，把大的数组往前排
2.计算这组数组的累计大小
3.用数组值的总大小-索引对应值的累计>索引对应值的累计！*/
func minSubsequence(nums []int) []int {
	var res []int
	sort.Ints(nums)
	sum := 0

	//累计和
	for _, n := range nums {
		sum += n
	}

	tInt := 0
	for i := 0; i < len(nums); i++ {
		tInt += nums[i]
		//如果超过时 i应该取上一步值 上一步正好是大于剩余数据之和且是最小序列
		if tInt > sum-tInt {

			//由于Num排序是从小到大 故此这里需要倒序排序
			return func(n []int) []int {
				var r []int
				for j := len(n) - 1; j >= 0; j-- {
					r = append(r, n[j])
				}
				return r
			}(nums[i-1:])
		}
	}

	return res
}