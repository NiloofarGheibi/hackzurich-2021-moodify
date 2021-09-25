package util

import "strconv"

// GetLimit gets limit as a string. It returns as an integer and in case of receiving bad size, sets to 10
func GetLimit(l string) int {
	limit, err := strconv.Atoi(l)
	if err != nil || limit < 0 {
		limit = 10
	}

	return limit
}

// GetOffset gets offset as a string. It returns as an integer and in case of receiving bad offset, sets to 0
func GetOffset(o string) int {
	offset, err := strconv.Atoi(o)
	if err != nil {
		offset = 0
	}
	if offset < 0 {
		offset = 0
	}

	return offset
}
