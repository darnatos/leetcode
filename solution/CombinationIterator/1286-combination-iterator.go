package CombinationIterator

type CombinationIterator struct {
	chars []byte
	clen  int
	curr  int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	chars := []byte(characters)
	l, r := 0, len(characters)-1
	for l < r {
		chars[l], chars[r] = chars[r], chars[l]
		l, r = l+1, r-1
	}
	return CombinationIterator{
		chars: chars,
		clen:  combinationLength,
		curr:  (1 << len(characters)) - 1,
	}
}

func countOne(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func (cbi *CombinationIterator) Next() string {
	for cbi.curr >= 0 && countOne(cbi.curr) != cbi.clen {
		cbi.curr--
	}
	res := make([]byte, cbi.clen)
	for i, j := 0, cbi.clen-1; i < len(cbi.chars); i++ {
		// fmt.Printf("%08b %08b %08b\n",cbi.curr,cbi.curr&(1<<i),(cbi.curr&(1<<i))>>i)
		if cbi.curr&(1<<i) > 0 {
			res[j] = cbi.chars[i]
			j--
		}
	}
	//fmt.Printf("%s %v %08b %s\n", string(cbi.chars), cbi.curr, cbi.curr, string(res))
	cbi.curr--
	return string(res)
}

func (cbi *CombinationIterator) HasNext() bool {
	for cbi.curr >= 0 && countOne(cbi.curr) != cbi.clen {
		cbi.curr--
	}
	return cbi.curr >= 0
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
