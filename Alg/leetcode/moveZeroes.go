func TestAlg1010(t *testing.T) {
	s := []int{0, 1, 0, 3, 12} //1,3,12,0,0
	moveZeroes(s)
	fmt.Println(s)
}

//283. 移动零
func moveZeroes(nums []int) {
	j := 0 //双指针 j为0指针
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 { //i一直++ 遇到不为0的值则兑换j和i的值
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}