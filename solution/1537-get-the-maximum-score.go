package solution

func MaxSum(nums1 []int, nums2 []int) int {
	modulo := 1000000007
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	A, B := 0, 0
	for i < m || j < n {
		if i >= m {
			B += nums2[j]
			j++
		} else if j >= n {
			A += nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				A += nums1[i]
				i++
			} else if nums1[i] > nums2[j] {
				B += nums2[j]
				j++
			} else {
				tmp := max(A, B) + nums1[i]
				A, B = tmp, tmp
				i++
				j++
			}
		}
	}
	return max(A, B) % modulo
}
