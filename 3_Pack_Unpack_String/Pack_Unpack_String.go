package packunpackstring

import (
	"strconv"
)

func PackString(s string) string {
	var res string
	var sum int

	curLet := rune(s[0])
	lenStr := len(s)
	for leterInd, leter := range s {
		if leter == curLet {
			sum += 1
		} else {
			res += strconv.Itoa(sum) + string(curLet)
			curLet = leter
			sum = 1
		}
		if leterInd == lenStr-1 {
			res += strconv.Itoa(sum) + string(curLet)
			curLet = leter
		}
	}
	return res
}

func UnpackString(s string) string {
	var count int
	var res string

	for _, leter := range s {
		if num, err := strconv.Atoi(string(leter)); err != nil {
			for j := 0; j < count; j++ {
				res += string(leter)
			}
			count = 1
		} else {
			count = num
		}
	}
	return res
}
