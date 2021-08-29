package solution

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)

	for i := range strs {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		s := string(b)
		if _, ok := m[s]; !ok {
			m[s] = []string{}
		}
		m[s] = append(m[s], strs[i])
	}
	for i := range m {
		res = append(res, m[i])
	}

	return res
}
