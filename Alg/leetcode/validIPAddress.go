func TestAlg1102(t *testing.T) {
	nums := "192.0.0.1"
	fmt.Println(validIPAddress(nums))
}

//468. 验证IP地址
func validIPAddress(IP string) string {
	if len(strings.Split(IP, ".")) == 4 {
		s := strings.Split(IP, ".")
		return validateIpv4(s)
	}

	if len(strings.Split(IP, ":")) == 8 {
		return validateIpv6(strings.Split(IP, ":"))
	}
	return "Neither"
}

func validateIpv4(ip []string) string {
	for _, v := range ip {
		if len(v) == 0 || (len(v) > 1 && v[0] == '0') {
			return "Neither"
		}

		if v[0] < '0' || v[0] > '9' {
			return "Neither"
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			return "Neither"
		}

		if n < 0 || n > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

func validateIpv6(ip []string) string {
	for _, v := range ip {
		if len(v) <= 0 || len(v) > 4 {
			return "Neither"
		}

		for i := 0; i < len(v); i++ {
			if v[i] >= '0' && v[i] <= '9' {
				continue
			}
			if v[i] >= 'A' && v[i] <= 'F' {
				continue
			}
			if v[i] >= 'a' && v[i] <= 'f' {
				continue
			}

			return "Neither"
		}
	}
	return "IPv6"
}