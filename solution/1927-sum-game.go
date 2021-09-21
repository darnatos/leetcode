package solution

import "sync"

func SumGame(num string) bool {
	m := len(num) / 2
	ls, rs := 0, 0
	lc, rc := 0, 0

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < m; i++ {
			if num[i] == '?' {
				lc++
			} else {
				ls += int(num[i] - '0')
			}
		}
		wg.Done()
	}()
	go func() {
		for i := m; i < len(num); i++ {
			if num[i] == '?' {
				rc++
			} else {
				rs += int(num[i] - '0')
			}
		}
		wg.Done()
	}()
	wg.Wait()

	if (lc+rc)&1 == 1 {
		return true
	}
	ans := float32(ls-rs) + float32((lc/2)*9) - float32((rc/2)*9)

	return ans != 0
}
