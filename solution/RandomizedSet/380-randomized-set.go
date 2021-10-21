package RandomizedSet

import "math/rand"

type RandomizedSet struct {
	nums []int
	nti  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		nti:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.nti[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.nti[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.nti[val]
	if !ok {
		return false
	}

	n := len(this.nums)
	this.nums[i] = this.nums[n-1]
	this.nti[this.nums[i]] = i
	this.nums = this.nums[:n-1]
	delete(this.nti, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Int()%len(this.nums)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
