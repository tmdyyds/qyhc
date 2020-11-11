func TestAlg1010(t *testing.T) {
    A := []int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}
    fmt.Println(average(A))
}


//1491. 去掉最低工资和最高工资后的工资平均值
func average(salary []int) float64 {
    maxIndex := func(salary []int) int {
        r := 0
        t := salary[0]

        for i := 1; i < len(salary); i++ {
            if salary[i] > t {
                t = salary[i]
                r = i
            }
        }

        return r
    }(salary)

    maxEnd := int(maxIndex + 1)
    salary = append(salary[:maxIndex], salary[maxEnd:]...)

    minIndex := func(salary []int) int {
        r := 0
        t := salary[0]

        for i := 1; i < len(salary); i++ {
            if salary[i] < t {
                t = salary[i]
                r = i
            }
        }

        return r
    }(salary)

    minEnd := int(minIndex + 1)
    salary = append(salary[:minIndex], salary[minEnd:]...)

    n := len(salary)
    nums := func(salary []int) int {
        nb := 0
        for _, v := range salary {
            nb += v
        }

        return nb
    }(salary)

    tmp := float64(nums) / float64(n)
    res, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", tmp), 64)

    return res
}