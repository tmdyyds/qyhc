//190. 颠倒二进制位
func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for i := 0; i < 32; i++ {
		ans = (ans << 1) + (num & 1) // num&1获取num的最低位, ans << 1 左移一位,然后加num的最低位
		num = num >> 1               //num右移
	}

	return ans


	//
	    var res uint32 = 0
    for i:=0; i<32; i++ {
        res <<= 1
        res |= num&1
        num >>= 1
    }
    return res

}
