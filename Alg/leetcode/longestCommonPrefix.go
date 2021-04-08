//14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n == 0 {
        return ""
    }

    prefix := strs[0]
    for i := 1; i < n; i++ {
        prefix = commonPrefix(prefix, strs[i])

        if len(prefix) == 0 {
            break
        }
    }

    return prefix
}

func commonPrefix(prefix, strs string) string  {
    len := min(len(prefix), len(strs))
    index := 0

    for index < len && prefix[index] == strs[index] {
        index++
    }

    return prefix[:index]
}

func min(x, y int) int  {
    if x < y {
       return x
    }

    return y
}