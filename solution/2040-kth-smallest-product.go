package solution

func KthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n := len(nums2)
	lo, hi := -int64(1e10)-1, int64(1e10)+1
	for lo < hi {
		//fmt.Println(lo, hi)
		var mid, cnt int64 = lo + (hi-lo)>>1, 0
		for _, i := range nums1 {
			//fmt.Println("  ", i)
			l, r, p := 0, n-1, 0
			if 0 <= i {
				for l <= r {
					c := l + (r-l)>>1
					mul := int64(i) * int64(nums2[c])
					//fmt.Println("    ", l, r, p, c, mul)
					if mul <= mid {
						p = c + 1
						l = c + 1
					} else {
						r = c - 1
					}
				}
			} else {
				for l <= r {
					c := l + (r-l)>>1
					mul := int64(i) * int64(nums2[c])
					//fmt.Println("    ", l, r, p, c, mul)
					if mul <= mid {
						p = n - c
						r = c - 1
					} else {
						l = c + 1
					}
				}
			}
			cnt += int64(p)
			//fmt.Println("      ", l, r, p, cnt)
		}
		//fmt.Println("      ", lo, hi, mid)
		if cnt >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
		//fmt.Println()
	}

	return lo
}
