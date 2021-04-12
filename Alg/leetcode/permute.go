
// result = []
// func backtrack(路径，选择列表) {
// 	if 满足结束条件 {
// 		result.add(路径)
// 	}
// 	return

// 	for 选择 in 选择列表 {
// 		做选择
// 		backtrack(路径，选择列表)
// 		撤销选择
// 	}
// }

//46. 全排列
var res [][]int

func permute(nums []int) [][]int {
	var trace []int
	used := make([]bool, len(nums))
	res = [][]int{}
	backtrack(nums, trace, used)
	return res
}

func backtrack(nums []int, trace []int, used []bool) {
	//结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(trace) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, trace)

		//把本次结果追加到最终结果上
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		//检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true

			//做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			trace = append(trace, nums[i])
			backtrack(nums, trace, used)

			//撤销操作
			trace = trace[:len(trace)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}
