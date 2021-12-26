package solution

func isIsomorphic(s string, t string) bool {
	m := [256]byte{} // ascii code
	for i := range m {
		m[i] = 255
	}
	u := [256]bool{}
	for i := range s {
		if m[s[i]] != 255 {
			if m[s[i]] != t[i] {
				return false
			}
		} else {
			if u[t[i]] {
				return false
			}
			m[s[i]] = t[i]
			u[t[i]] = true
		}
	}
	return true
}
