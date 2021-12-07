package BrowserHistroy

type BrowserHistory struct {
	ht  []string
	cur int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.cur++
	bh.ht = bh.ht[:bh.cur]
	bh.ht = append(bh.ht, url)
}

func (bh *BrowserHistory) Back(steps int) string {
	bh.cur -= steps
	if bh.cur < 0 {
		bh.cur = 0
	}
	return bh.ht[bh.cur]
}

func (bh *BrowserHistory) Forward(steps int) string {
	bh.cur += steps
	if bh.cur >= len(bh.ht) {
		bh.cur = len(bh.ht) - 1
	}
	return bh.ht[bh.cur]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
