package solution

func compareVersion(version1 string, version2 string) int {
	i0, j0, m, n := 0, 0, len(version1), len(version2)
	for cnt1, cnt2 := 0, 0; i0 < m && j0 < n; {
		i := i0
		for ; i <= m; i++ {
			if i == m || version1[i] == '.' {
				cnt1++
				break
			}
		}
		j := j0
		for ; j <= n; j++ {
			if j == n || version2[j] == '.' {
				cnt2++
				break
			}
		}
		if cnt1 != cnt2 {
			break
		}
		for i0 < i && version1[i0] == '0' {
			i0++
		}
		for j0 < j && version2[j0] == '0' {
			j0++
		}
		if i-i0 > j-j0 {
			return 1
		} else if i-i0 < j-j0 {
			return -1
		}
		for i1, j1 := i-1, j-1; i1 >= i0 && j1 >= j0; i1, j1 = i1-1, j1-1 {
			if version1[i1] < version2[j1] {
				return -1
			} else if version1[i1] > version2[j1] {
				return 1
			}
		}
		i0 = i + 1
		j0 = j + 1
		if i0 > m {
			i0 = m
		}
		if j0 > n {
			j0 = n
		}
	}
	for i0 < m {
		if version1[i0] != '.' && version1[i0] != '0' {
			return 1
		}
		i0++
	}
	for j0 < n {
		if version2[j0] != '.' && version2[j0] != '0' {
			return -1
		}
		j0++
	}
	return 0
}
