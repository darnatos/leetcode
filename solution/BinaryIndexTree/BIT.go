package BinaryIndexTree

type BIT struct {
	t []int
	n int
}

func Constructor(n int) BIT {
	return BIT{t: make([]int, n+1), n: n + 1}
}
func (b *BIT) Add(i int, delta int) {
	for i++; i < b.n; i += i & (-i) {
		b.t[i] += delta
	}
}
func (b BIT) Query(i int) int {
	sum := 0
	for i++; i > 0; i -= i & (-i) {
		sum += b.t[i]
	}
	return sum
}
