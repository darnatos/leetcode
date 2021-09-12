package solution

import "strconv"

func DayOfYear(date string) int {
	y, _ := strconv.Atoi(date[:4])
	m, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:])
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	res := d
	if m > 2 && y%4 == 0 {
		res += 1
	}
	for i := 0; i < m-1; i++ {
		res += days[i]
	}
	return res
}
