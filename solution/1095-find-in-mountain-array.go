package solution

/**
* // This is the MountainArray's API interface.
* // You should not implement it, or speculate about its implementation
* type MountainArray struct {
* }
*
* func (this *MountainArray) get(index int) int {}
* func (this *MountainArray) length() int {}
 */

type MountainArray []int

func (ma *MountainArray) get(index int) int {
	return (*ma)[index]
}
func (ma *MountainArray) length() int {
	return len(*ma)
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	p := peakIndexInMountainArray(mountainArr)
	ll := bs(target, mountainArr, 0, p)
	if ll != -1 {
		return ll
	}
	return bsl(target, mountainArr, p+1, mountainArr.length()-1)
}

func peakIndexInMountainArray(arr *MountainArray) int {
	l, r := 1, arr.length()-2
	for l <= r {
		m := l + (r-l)/2
		if arr.get(m) <= arr.get(m-1) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func bs(target int, arr *MountainArray, l, r int) int {
	for l <= r {
		m := l + (r-l)/2
		if arr.get(m) == target {
			return m
		} else if arr.get(m) > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func bsl(target int, arr *MountainArray, l, r int) int {
	for l <= r {
		m := l + (r-l)/2
		if arr.get(m) == target {
			return m
		} else if arr.get(m) < target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
