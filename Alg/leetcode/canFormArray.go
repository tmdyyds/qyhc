func TestAlg1010(t *testing.T) {
	arr := []int{37, 69, 3, 74, 46}
	pieces := [][]int{{37, 69, 3, 74, 46}}
	fmt.Println(canFormArray(arr, pieces))
}

//1640. 能否连接形成数组
func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int)

	//pieces每个slice中的第一个值作为key
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i + 1
	}

	//这里有一个隐含条件 如果pieces某个元素是单个的话是不需要检查的
	for j := 0; j < len(pieces); j++ {
		//长度为1的slice 如果m中不存在 则返回false,不存在的值 会为0
		if len(pieces[j]) == 1 && m[pieces[j][0]] < 1 {
			return false
		}

		if len(pieces[j]) > 1 {
			for k := 0; k < len(pieces[j])-1; k++ {
				//slice中元素与arr中元素value得挨着
				if m[pieces[j][k+1]]-m[pieces[j][k]] != 1 {
					return false
				}
			}
		}
	}

	return true
}