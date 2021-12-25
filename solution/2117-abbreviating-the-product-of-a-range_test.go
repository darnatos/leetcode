package solution

import "testing"

func TestAbbreviateProduct(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{1, 4}, "24e0"},
		{"test2", args{999998, 1000000}, "99999...00002e6"},
		{"test3", args{256, 65535}, "23510...78528e16317"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbbreviateProduct(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("AbbreviateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
