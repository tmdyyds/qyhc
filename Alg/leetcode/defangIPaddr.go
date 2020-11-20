func TestAlg1010(t *testing.T) {
	s := "1.1.1.1"
	fmt.Println(defangIPaddr(s))
}


//1108. IP 地址无效化
func defangIPaddr(address string) string {
	m := map[string]string{
		".": "[.]",
	}
	var r string

	for _, v := range address {
		if _, ok := m[string(v)]; ok {
			r += m[string(v)]
			continue
		}

		r += string(v)
	}

	return r

	// return strings.ReplaceAll(address, ".", "[.]")
}