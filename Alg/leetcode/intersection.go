//349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
    hash := make(map[int]bool)
    res := make([]int, 0)

    for _, v := range nums1 {
        hash[v] = true
    }

    for _, v2 := range nums2 {
        if hash[v2] == true {
            res = append(res, v2)
            hash[v2] = false
        }
    }

    return res
}