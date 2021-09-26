package solution

func MovesToChessboard(b [][]int) int {
	n := len(b)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[0][0]^b[i][0]^b[0][j]^b[i][j] == 1 {
				return -1
			}
		}
	}

	rc, cc, rs, cs := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		rc += b[i][0]
		cc += b[0][i]
		rs += b[i][0] ^ (i & 1)
		cs += b[0][i] ^ (i & 1)
	}

	if rc != n/2 && rc != (n+1)/2 {
		return -1
	}
	if cc != n/2 && cc != (n+1)/2 {
		return -1
	}

	if n&1 == 1 {
		if rs&1 == 1 {
			rs = n - rs
		}
		if cs&1 == 1 {
			cs = n - cs
		}
	} else {
		rs = min(rs, n-rs)
		cs = min(cs, n-cs)
	}

	return (rs + cs) / 2
}
