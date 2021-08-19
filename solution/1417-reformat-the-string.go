package solution

func reformat(s string) string {
	d := make([]byte, 0, 16)
	a := make([]byte, 0, 16)
	for i := range s {
		if isDigit(s[i]) {
			d = append(d, s[i])
			continue
		}
		a = append(a, s[i])
	}

	dl := len(d)
	al := len(a)
	if dl-al > 1 || al-dl > 1 {
		return ""
	}

	res := make([]byte, 0, len(s))
	if dl > al {
		for i := 0; i < dl; i++ {
			res = append(res, d[i])
			if i < al {
				res = append(res, a[i])
			}
		}
	} else {
		for i := 0; i < al; i++ {
			res = append(res, a[i])
			if i < dl {
				res = append(res, d[i])
			}
		}
	}

	return string(res)
}

func isDigit(v byte) bool {
	return v >= '0' && v <= '9'
}
