package solution

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k1, k2 := (m+n+1)/2, (m+n+2)/2
	if k1 == k2 {
		return float64(findKth(nums1, 0, nums2, 0, k1))
	}
	return float64(findKth(nums1, 0, nums2, 0, k1)+findKth(nums1, 0, nums2, 0, k2)) / 2
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	m, n := len(nums1), len(nums2)
	if i >= m {
		return nums2[j+k-1]
	}
	if j >= n {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	var midVal1, midVal2 int
	if i+k/2-1 < m {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = math.MaxInt32
	}
	if j+k/2-1 < n {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = math.MaxInt32
	}
	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
