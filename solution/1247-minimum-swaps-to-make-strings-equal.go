package solution

func MinimumSwap(s1 string, s2 string) int {
	xy, yx := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				xy++
			} else {
				yx++
			}
		}
	}

	if (xy+yx)%2 > 0 {
		return -1
	}
	// 1  , 1  , 2
	//"xx","yy","xy"
	//"yy","xx","yx"

	res := xy/2 + yx/2 + xy%2*2

	return res
}
