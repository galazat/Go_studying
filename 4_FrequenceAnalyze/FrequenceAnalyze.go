package frequenceanalyze

import (
	"sort"
)

// func main() {
// 	s := "shgudihgdfhbiueghkasvbrudfjgsduighkdfbukahiusrahfgsdfygisrhidsughibhkguuukhukgjhgjggjgjgjgugujgjgjgjgjugjugjgjgjhhghjgjguhugue"
// 	fmt.Println(len(s))
// 	FrequenceTopTen(s)
// }

func FrequenceStringTopTen(s string) []string {
	letterMap := make(map[string]int)
	for _, element := range s {
		letterMap[string(element)] += 1
	}

	keys := make([]string, 0, len(letterMap))

	for key := range letterMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return letterMap[keys[i]] > letterMap[keys[j]]
	})

	res := keys[:10]
	return res
}
