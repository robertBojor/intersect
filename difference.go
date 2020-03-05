package intersect

// SubtractString - subtract a string slice from another
func SubtractString(fromSlice, subtractSlice []string) (result []string) {
	if len(fromSlice) == 0 {
		return
	}
	if len(subtractSlice) == 0 {
		result = fromSlice
		return
	}
	hash := make(map[string]bool)
	for _, e := range subtractSlice {
		hash[e] = true
	}
	for _, e := range fromSlice {
		if hash[e] {
			continue
		}
		result = append(result, e)
	}
	return
}

// SubtractInt - subtract an int slice from another
func SubtractInt(fromSlice, subtractSlice []int) (result []int) {
	if len(fromSlice) == 0 {
		return
	}
	if len(subtractSlice) == 0 {
		result = fromSlice
		return
	}
	hash := make(map[int]bool)
	for _, e := range subtractSlice {
		hash[e] = true
	}
	for _, e := range fromSlice {
		if hash[e] {
			continue
		}
		result = append(result, e)
	}
	return
}
// SubtractInt64 - subtract an int64 slice from another
func SubtractInt64(fromSlice, subtractSlice []int64) (result []int64) {
	if len(fromSlice) == 0 {
		return
	}
	if len(subtractSlice) == 0 {
		result = fromSlice
		return
	}
	hash := make(map[int64]bool)
	for _, e := range subtractSlice {
		hash[e] = true
	}
	for _, e := range fromSlice {
		if hash[e] {
			continue
		}
		result = append(result, e)
	}
	return
}

