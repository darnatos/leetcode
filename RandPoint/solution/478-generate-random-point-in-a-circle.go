package solution

import (
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	r, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	ll := this.r * math.Sqrt(rand.Float64())
	deg := 2 * rand.Float64() * math.Pi

	xr := ll * math.Cos(deg)
	yr := ll * math.Sin(deg)
	return []float64{this.x + xr, this.y + yr}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
