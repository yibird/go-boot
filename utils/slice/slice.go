package slice

import "strconv"

func ToIntSlice(s []string) []int {
	if s == nil {
		return nil
	}
	intSlice := make([]int, len(s))
	for i, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil
		}
		intSlice[i] = num
	}
	return intSlice
}

func ToStrSlice(s []int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = strconv.Itoa(v)
	}
	return r
}
