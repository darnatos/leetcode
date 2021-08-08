package solution

import (
	"fmt"
	"strings"
)

var Month = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func ReformatDate(date string) string {
	return reformatDate(date)
}

func reformatDate(date string) string {
	split := strings.Split(date, " ")
	day := split[0][0 : len(split[0])-2]
	if len(day) < 2 {
		day = "0" + day
	}
	return fmt.Sprintf("%v-%v-%v", split[2], Month[split[1]], day)
}
