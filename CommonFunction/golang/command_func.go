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

//start
// defer trace("bigSlowOperation")() //注意最后括号
// time.Sleep(10 * time.Second)
// 函数或功能的执行时间
func TimeDiff(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
//end

//start
//map的key,不为string为slice(这种方式可以处理任意一种不可比较的key,不仅仅是slice)
var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

//比较
s1 := []string{"a", "b"}
s2 := []string{"c", "d"}
s3 := []string{"e", "f"}
s4 := []string{"a", "b"}

Add(s1)
Add(s2)
Add(s3)

//判断
for key, v := range m {
	if key == k(s4) {
		fmt.Println(key, v, "===")
	}
}

//end


//给map按着key排序,并输出/打印map排序后数据
// m := map[int]int{
// 		4: 40,
// 		1: 10,
// 		3: 30,
// 	}
// 	sortMap(m)
func SortMapByKey(m map[int]int) {
	slice := []int{}
	for k, _ := range m {
		slice = append(slice, k)
	}

	sort.Ints(slice)
	for _, k := range slice {
		fmt.Println(m[k])
	}
}


//start
//删除切片指定位置元素且保持原来元素的位置不变
// s := []int{5, 6, 7, 8, 9}
// fmt.Println(remove(s, 2)) // "[5 6 9 8]

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

//start

// 1*1=1
// 1*2=2 2*2=4
// 1*3=3 2*3=6 3*3=9
// 1*4=4 2*4=8 3*4=12 4*4=16
// 1*5=5 2*5=10 3*5=15 4*5=20 5*5=25
// 1*6=6 2*6=12 3*6=18 4*6=24 5*6=30 6*6=36
// 1*7=7 2*7=14 3*7=21 4*7=28 5*7=35 6*7=42 7*7=49
// 1*8=8 2*8=16 3*8=24 4*8=32 5*8=40 6*8=48 7*8=56 8*8=64
// 1*9=9 2*9=18 3*9=27 4*9=36 5*9=45 6*9=54 7*9=63 8*9=72 9*9=81
//九九乘法表
func MpTable99() {
	level := 9
	for i := 1; i <= level; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, " ")
		}

		fmt.Println()
	}
}

//end