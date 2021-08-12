package solution

import (
	"sort"
)

type Friend struct {
	id     int
	arrive int
	leave  int
	chair  int
}

type FriendList []Friend

func (fl FriendList) Len() int {
	return len(fl)
}

func (fl FriendList) Less(i, j int) bool {
	return fl[i].arrive < fl[j].arrive
}

func (fl FriendList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

func SmallestChair(times [][]int, targetFriend int) int {
	fl := make(FriendList, len(times))

	for i, time := range times {
		fl[i] = Friend{i, time[0], time[1], -1}
	}
	sort.Sort(fl)

	chairs := make([]int, len(times))
	for i := range fl {
		j := 0
		for ; j <= i; j++ {
			if chairs[j] <= fl[i].arrive {
				break
			}
		}
		chairs[j] = fl[i].leave
		fl[i].chair = j
		if targetFriend == fl[i].id {
			return fl[i].chair
		}
	}
	return -1
}
