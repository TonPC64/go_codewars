package kata

import "strings"

func my_duplicate_count(s1 string) int {
	dup_count := make(map[string]int)
	for _, alphabet := range strings.Split(s1, "") {
		key := strings.ToLower(alphabet)
		if _, ok := dup_count[key]; ok {
			dup_count[key] = dup_count[key] + 1
		} else {
			dup_count[key] = 0
		}
	}
	count := 0
	for _, val := range dup_count {
		if val > 0 {
			count++
		}
	}
	return count
}

func crever_duplicate_count(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}
