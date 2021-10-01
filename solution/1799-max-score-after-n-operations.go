package solution

import (
	"leetcode/util"
)

func MaxScore(nums []int) int {
	n := 1 << len(nums)
	tmp := make([]int, (len(nums)/2+1)*n)
	dp := make([][]int, len(nums)/2+1)
	for i := range dp {
		dp[i] = tmp[i*n : (i+1)*n]
	}

	res := msDfs(nums, dp, 1, 0)
	return res
}

func msDfs(nums []int, dp [][]int, i, mask int) int {
	if i > len(nums)/2 {
		return 0
	}
	if dp[i][mask] == 0 {
		for j := 0; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				newMask := 1<<j + 1<<k
				if mask&newMask == 0 {
					dp[i][mask] = util.Max(dp[i][mask], i*gcd(nums[j], nums[k])+msDfs(nums, dp, i+1, mask+newMask))
				}
				//fmt.Printf("%d %d %d %8b %8b %d\n", i, j, k, mask, newMask, dp[i][mask])
			}
		}
	}
	return dp[i][mask]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
