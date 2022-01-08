package solution

func maxJumps(arr []int, d int) int {
	res := 0
	dp := make([]int, len(arr))
	for i := range arr {
		res = max(res, mjDfs(arr, i, d, dp))
	}
	return res
}

func mjDfs(arr []int, i, d int, dp []int) int {
	if dp[i] != 0 {
		return dp[i]
	}
	res := 1
	for j := i + 1; j <= min(i+d, len(arr)-1) && arr[j] < arr[i]; j++ {
		res = max(res, mjDfs(arr, j, d, dp)+1)
	}
	for j := i - 1; j >= max(i-d, 0) && arr[j] < arr[i]; j-- {
		res = max(res, mjDfs(arr, j, d, dp)+1)
	}
	dp[i] = res
	return dp[i]
}
