//start
//删除切片指定位置元素且保持原来元素的位置不变
func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 9 8]
}

//删除切片元素 保持原顺序
func remove(slice []int, i int) []int {
	fmt.Println(slice[i:], slice[i+1:], "====")
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//删除切片元素 不保持原顺序
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] //位置元素被最后一位替换
    return slice[:len(slice)-1]
}

//end