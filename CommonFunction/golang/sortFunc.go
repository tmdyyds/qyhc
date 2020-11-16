//全部关于go中排序相关的函数

//给map按着key排序,并输出/打印map排序后数据
// m := map[int]int{
// 		4: 40,
// 		1: 10,
// 		3: 30,
// 	}
// 	sortMap(m)
func sortMapByKey(m map[int]int) {
	slice := []int{}
	for k, _ := range m {
		slice = append(slice, k)
	}

	sort.Ints(slice)
	for _, k := range slice {
		fmt.Println(m[k])
	}
}