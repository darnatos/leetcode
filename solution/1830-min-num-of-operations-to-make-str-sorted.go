package solution

var ft = [3001]int{1}
var im = [3001]int{1}

func makeStringSorted(s string) int {
	sz, res, cnt := len(s), 0, [26]int{}

	if ft[1] == 0 {
		for i := 1; i <= 3000; i++ {
			// factorial, n!
			ft[i] = i * ft[i-1] % mod
			// inverse multiplication, 1/(n! % mod)
			// Fermat's little theorem:
			// 	      a^p % p = a % p. Divide a^2
			// => a^(p-2) % p = a/a^2 % p = 1/a % p. Let a to be n!, p to be mod
			// =>  1/n! % mod = n!^(mod-2) % mod, which is modPow(n!, mod-2, mod)
			im[i] = modPow(ft[i], mod-2, mod)
		}
	}
	for i := sz - 1; i >= 0; i-- {
		cnt[s[i]-'a']++
		prems := 0
		for j := 0; j < int(s[i]-'a'); j++ {
			prems += cnt[j]
		}
		prems = prems * ft[sz-i-1] % mod
		for j := range cnt {
			prems = prems * im[cnt[j]] % mod
		}
		res = (res + prems) % mod
	}
	return res
}

func modPow(x, y, m int) int {
	res := 1
	for y > 0 {
		if y&1 > 0 {
			res = res * x % m
		}
		x = x * x % m
		y >>= 1
	}
	return res
}
