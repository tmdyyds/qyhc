## GO面试题相关
#### iota输出
    const (
        a    = iota
        b    = iota
        c, d = iota, iota
    )

    fmt.Println(a, b, c, d)
    //0 1 2 2

#### test