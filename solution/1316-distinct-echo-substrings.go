package solution

func distinctEchoSubstrings(text string) int {
	res := make(map[int]struct{})
	var sh_l, sh_r, p int
	for m := 1; m < len(text); m++ {
		sh_l, sh_r, p = 0, 0, 1
		for n := 1; m-n >= 0 && m+n-1 < len(text); n++ {
			sh_l = (sh_l*26 + int(text[m-n])) % mod
			sh_r = (sh_r + int(text[m+n-1])*p) % mod
			p = p * 26 % mod
			if sh_l == sh_r {
				res[sh_l] = struct{}{}
			}
		}
	}
	return len(res)
}
