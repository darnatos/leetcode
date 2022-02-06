package bitset

type Bitset struct {
	s   []byte
	f   byte
	cnt int
}

func Constructor(size int) Bitset {
	s := make([]byte, size)
	return Bitset{s, 0, 0}
}

func (this *Bitset) Fix(idx int) {
	if this.s[idx] == this.f {
		this.cnt++
	}
	this.s[idx] = 1 ^ this.f
}

func (this *Bitset) Unfix(idx int) {
	if this.s[idx] != this.f {
		this.cnt--
	}
	this.s[idx] = 0 ^ this.f
}

func (this *Bitset) Flip() {
	this.cnt = len(this.s) - this.cnt
	this.f = 1 ^ this.f
}

func (this *Bitset) All() bool {
	return this.cnt == len(this.s)
}

func (this *Bitset) One() bool {
	return this.cnt > 0
}

func (this *Bitset) Count() int {
	return this.cnt
}

func (this *Bitset) ToString() string {
	res := make([]byte, len(this.s))
	for i := range this.s {
		if this.s[i] != this.f {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
