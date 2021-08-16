package solution

import "sort"

func NumTriplets(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	return foo(nums1, nums2) + foo(nums2, nums1)
}

func foo(nums1, nums2 []int) int {
	res := 0
	m, n := len(nums1), len(nums2)
	for i := 0; i < m; i++ {
		s := nums1[i] * nums1[i]
		j, k := 0, n-1
		for j < k {
			p := nums2[j] * nums2[k]
			if s < p {
				k--
			} else if s > p {
				j++
			} else if nums2[j] != nums2[k] {
				j0 := j
				k0 := k
				for nums2[j] == nums2[j0] {
					j++
				}
				for nums2[k] == nums2[k0] {
					k--
				}
				res += (j - j0) * (k0 - k)

			} else {
				res += (k - j) * (k - j + 1) / 2
				break
			}
		}
	}
	return res
}
