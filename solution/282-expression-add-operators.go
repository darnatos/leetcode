package solution

import "strconv"

func AddOperators(num string, target int) []string {
	res := make([]string, 0)
	if len(num) == 0 {
		return res
	}
	aoHelper(&res, "", num, target, 0, 0, 0)
	return res
}

func aoHelper(res *[]string, path, num string, target, pos, eval, multi int) {
	if pos == len(num) {
		if target == eval {
			*res = append(*res, path)
		}
		return
	}
	for i := pos; i < len(num); i++ {
		if i != pos && num[pos] == '0' {
			break
		}
		curr := num[pos : i+1]
		cur, _ := strconv.Atoi(curr)
		if pos == 0 {
			aoHelper(res, path+curr, num, target, i+1, cur, cur)
		} else {
			aoHelper(res, path+"+"+curr, num, target, i+1, eval+cur, cur)
			aoHelper(res, path+"-"+curr, num, target, i+1, eval-cur, -cur)
			aoHelper(res, path+"*"+curr, num, target, i+1, eval-multi+multi*cur, multi*cur)
		}
	}
}
