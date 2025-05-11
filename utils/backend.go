package utils

import "strconv"

func Pagination(pagestr string, lengthstr string) (int, int) {
	page, _ := strconv.Atoi(pagestr)
	length, _ := strconv.Atoi(lengthstr)
	if page < 1 {
		page = 1
	}
	if length < 1 {
		length = 4
	}
	return page, length
}
