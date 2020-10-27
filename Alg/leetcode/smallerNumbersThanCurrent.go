//1365. 有多少小于当前数字的数字

/*
    暴力破解 应该是最简单的 直接每个数字去比较数组中所有值
    时间复杂度：O(N^2)，其中 N 为数组的长度。
    空间复杂度：O(1)
*/
func smallerNumbersThanCurrent(nums []int) []int {
    var r = make([]int, len(nums))
    cnt := len(nums)
    if cnt < 1 {
        return r
    }

    for i := 0; i < cnt; i++ {
        n := 0
        for j := 0; j < len(nums); j++ {
            if nums[i] > nums[j] {
                n++
                r[i] = n
            }
        }
    }

    return r
}

/*
    计数排序
    时间复杂度：O(N+k)，k为开辟值域大小。
    空间复杂度：O(k )
*/
func smallerNumbersThanCurrent(nums []int) []int {
    cnt := [101]int{}
    n := len(nums)

    //遍历原数组，统计每个数字出现次数
    for _, v := range nums {
        cnt[v]++
    }
    //fmt.Println(cnt)

    //遍历统计数组，求出小于每个数字的数字数
    for i := 0; i < 100; i++ {
        cnt[i+1] += cnt[i] //这里的key 是nums的value
    }

    r := make([]int, n)
    for k, v := range nums {
        if v > 0 {
            r[k] = cnt[v-1]
        }
    }

    return r
}