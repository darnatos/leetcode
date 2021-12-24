package RangeModule

type RangeModule struct {
	rg []int
}

func Constructor() RangeModule {
	rg := make([]int, 0)
	return RangeModule{rg}
}

func (this *RangeModule) AddRange(left int, right int) {
	start := bisectLeft(this.rg, left)
	end := bisectRight(this.rg, right)

	rg := make([]int, 0, len(this.rg)+2)
	rg = append(rg, this.rg[:start]...)
	if start&1 == 0 {
		rg = append(rg, left)
	}
	if end&1 == 0 {
		rg = append(rg, right)
	}
	rg = append(rg, this.rg[end:]...)
	this.rg = rg
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	start, end := bisectLeft(this.rg, right), bisectRight(this.rg, left)
	return start == end && start&1 == 1
}

func (this *RangeModule) RemoveRange(left int, right int) {
	start := bisectLeft(this.rg, left)
	end := bisectRight(this.rg, right)

	rg := make([]int, 0, len(this.rg)+2)
	rg = append(rg, this.rg[:start]...)
	if start&1 == 1 {
		rg = append(rg, left)
	}
	if end&1 == 1 {
		rg = append(rg, right)
	}
	rg = append(rg, this.rg[end:]...)
	this.rg = rg
}

func bisectLeft(nums []int, tg int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] < tg {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func bisectRight(nums []int, tg int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= tg {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
