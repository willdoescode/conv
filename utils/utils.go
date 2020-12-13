package utils

import "strconv"

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ThrowError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReverseSlice(s interface{}) interface{} {
	switch s.(type) {
	case []int:
		for i, j := 0, len(s.([]int))-1; i < j; i, j = i+1, j-1 {
			s.([]int)[i], s.([]int)[j] = s.([]int)[j], s.([]int)[i]
		}
		return s
	case []string:
		for i, j := 0, len(s.([]string))-1; i < j; i, j = i+1, j-1 {
			s.([]string)[i], s.([]string)[j] = s.([]string)[j], s.([]string)[i]
		}
		return s
	}
	return s
}
