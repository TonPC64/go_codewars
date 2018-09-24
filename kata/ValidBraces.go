package kata

func ValidBraces(str string) bool {
	value := map[string]int{"{": 11, "[": 21, "(": 31, "}": -11, "]": -21, ")": -31}
	result := 0
	for i, v := range str {
		key := string(v)
		if i > 0 && value[key] < 0 && value[string(str[i-1])] > 0 && value[string(str[i-1])]+value[key] != 0 {
			return false
		}
		result += value[key]
	}
	if result == 0 {
		return true
	}
	return false
}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func clever_ValidBraces(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}
