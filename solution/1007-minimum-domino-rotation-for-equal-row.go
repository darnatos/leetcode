package solution

import (
	"math"
	"sync"
)

func MinDominoRotations(tops []int, bottoms []int) int {
	var wg = sync.WaitGroup{}
	r1 := 0
	r2 := 0
	r3 := 0
	r4 := 0

	wg.Add(4)
	go func() {
		r1 = dominoHelper(tops[0], tops, bottoms)
		wg.Done()
	}()
	go func() {
		r2 = dominoHelper(bottoms[0], bottoms, tops)
		wg.Done()
	}()
	go func() {
		r3 = dominoHelper(bottoms[0], tops, bottoms)
		wg.Done()
	}()
	go func() {
		r4 = dominoHelper(tops[0], bottoms, tops)
		wg.Done()
	}()
	wg.Wait()

	res := minIn(r1, r2, r3, r4)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func dominoHelper(target int, tops, bottoms []int) int {
	c1 := 0
	for i := 0; i < len(tops); i++ {
		if target != tops[i] {
			c1++
			if target != bottoms[i] {
				return math.MaxInt32
			}
		}
	}

	return c1
}

func minIn(a ...int) int {
	res := math.MaxInt32

	for i := 0; i < 4; i++ {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}
