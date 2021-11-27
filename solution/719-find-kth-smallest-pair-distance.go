package solution

import "sort"

func SmallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	n:=len(nums)
	l:=0
	r:=nums[n-1]-nums[0]

	for l<r{
		m:=l+(r-l)/2
		cnt:=0
		for i,j:=0,0;i<n;i++{
			for j<n && nums[j]-nums[i]<=m{
				j++
			}
			cnt+=j-i-1
		}
		if cnt < k{
			l=m+1
		} else {
			r=m
		}
	}
	return l
}
