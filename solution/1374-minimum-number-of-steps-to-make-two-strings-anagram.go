package solution

import (
	"leetcode/util"
	"sync"
)

var wg = sync.WaitGroup{}

func MinSteps(s string, t string) int {
	m := make([]int, 26)
	n := make([]int, 26)

	wg.Add(2)
	go func() {
		for i := 0; i < len(s); i++ {
			m[s[i]-'a'] += 1
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < len(s); i++ {
			n[t[i]-'a'] += 1
		}
		wg.Done()
	}()
	wg.Wait()

	var c int
	ch := make(chan int)

	go func() {
		for v := range ch {
			c += v
		}
	}()

	wg.Add(2)
	go bar(m[:14], n[:14], ch)
	go bar(m[14:26], n[14:26], ch)
	wg.Wait()

	close(ch)

	return c >> 1
}

func bar(m, n []int, ch chan<- int) {
	for i := 0; i < len(m); i++ {
		ch <- util.Abs(m[i] - n[i])
	}
	wg.Done()
}
