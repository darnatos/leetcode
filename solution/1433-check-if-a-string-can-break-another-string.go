package solution

import "sort"

type bArray []byte

func (a bArray) Len() int {
	return len(a)
}

func (a bArray) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a bArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func CheckIfCanBreak(s1 string, s2 string) bool {
	b1 := bArray(s1)
	b2 := bArray(s2)

	sort.Sort(b1)
	sort.Sort(b2)

	var diff int
	for i := 0; i < len(b1); i++ {
		if diff > 0 {
			if int(b1[i])-int(b2[i]) < 0 {
				return false
			}
		} else if diff < 0 {
			if int(b1[i])-int(b2[i]) > 0 {
				return false
			}
		} else {
			diff = int(b1[i]) - int(b2[i])
		}
	}

	return true
}
