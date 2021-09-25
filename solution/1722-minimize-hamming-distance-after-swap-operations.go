package solution

func MinimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	ds := make([]int, len(source))
	for i := range ds {
		ds[i] = i
	}
	for _, swap := range allowedSwaps {
		union(ds, swap[0], swap[1])
	}
	g := getGroups(ds)

	res := 0
	for _, v := range g {
		m := make(map[int]int, len(v))
		for _, j := range v {
			m[source[j]]++
		}
		for _, j := range v {
			if m[target[j]] > 0 {
				m[target[j]]--
			} else {
				res++
			}
		}
	}

	return res
}

func find(ds []int, k int) int {
	if ds[k] == k {
		return k
	}
	ds[k] = find(ds, ds[k])
	return ds[k]
}

func union(ds []int, u, v int) {
	pu, pv := find(ds, u), find(ds, v)
	if pu != pv {
		ds[pu] = pv
	}
}

func getGroups(ds []int) map[int][]int {
	groups := make(map[int][]int)
	for u := range ds {
		fu := find(ds, u)
		groups[fu] = append(groups[fu], u)
	}
	return groups
}
