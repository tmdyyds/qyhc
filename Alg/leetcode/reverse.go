//7. 整数反转
func reverse(x int) int {
	min := int(math.Pow(-2, 31))
	max := int(math.Pow(2, 31) - 1)

	reverse := 0
	for x != 0 {
		reverse = x%10 + reverse*10
		x = x / 10
	}

	if reverse > max || reverse < min {
		return 0
	}

	return reverse
}
