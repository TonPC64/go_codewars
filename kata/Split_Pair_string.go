package kata

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += "_"
	}
	result := make([]string, len(str)/2)
	for i, s := range str {
		result[i/2] = result[i/2] + string(s)
	}
	return result
}

func clever_Solution(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
