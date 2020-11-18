func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	//拼接两个s2，则如果是由s1旋转而成，则拼接后的s一定包含s1.
	s3 := s2 + s2
	return strings.Contains(s3, s1)

	//利用Builder
	if len(s1) != len(s2) {
		return false
	}

	var build strings.Builder
	build.WriteString(s2)
	build.WriteString(s2)

	return strings.Contains(build.String(), s1)
}