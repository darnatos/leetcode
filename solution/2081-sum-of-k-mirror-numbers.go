package solution

func KMirror(k int, n int) int64 {
	res := int64(0)
	i := int64(0)
	for {
		i = nextMir(i)
		if iskMir(i, int64(k)) {
			res += i
			n--
			if n == 0 {
				break
			}
		}
	}
	return res
}

func nextMir(a int64) int64 {
	a++
	arr := make([]int64, 0, 64)
	for tmp := a; tmp > 0; tmp /= 10 {
		arr = append(arr, tmp%10)
	}
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	mid := len(arr) / 2
	for {
		for i := 0; i < mid; i++ {
			arr[len(arr)-1-i] = arr[i]
		}
		tmp := int64(0)
		for i := range arr {
			tmp = tmp*10 + arr[i]%10
		}
		if tmp >= a {
			return tmp
		} else {
			j := mid - 1
			if len(arr)%2 == 1 {
				j = mid
			}
			for arr[j] == 9 {
				arr[j] = 0
				j--
			}
			arr[j]++
		}
	}
}

func iskMir(a, k int64) bool {
	mir := int64(0)
	tmp := a
	for tmp > 0 {
		mir = mir*k + tmp%k
		tmp /= k
	}
	return a == mir
}
