package intersect

// Strings: intersect two slices of type string
func Strings(s1, s2 []string) (intersection []string) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			intersection = append(intersection, e)
		}
	}
	intersection = uniqueStrings(intersection)
	return
}

// Ints: Intersect two slices of ints
func Ints(s1, s2 []int) (intersection []int) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			intersection = append(intersection, e)
		}
	}
	intersection = uniqueInts(intersection)
	return
}

// Int64s: Intersect two slices of int64
func Int64s(s1, s2 []int64) (intersection []int64) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	hash := make(map[int64]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			intersection = append(intersection, e)
		}
	}
	intersection = uniqueInt64s(intersection)
	return
}

// Float32s: Intersect two slices of float32
func Float32s(s1, s2 []float32) (intersection []float32) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	hash := make(map[float32]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			intersection = append(intersection, e)
		}
	}
	intersection = uniqueFloat32s(intersection)
	return
}

// Float64s: Intersect two slices of float32
func Float64s(s1, s2 []float64) (intersection []float64) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	hash := make(map[float64]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			intersection = append(intersection, e)
		}
	}
	intersection = uniqueFloat64s(intersection)
	return
}

/* Helpers */
func uniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
func uniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
func uniqueInt64s(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
func uniqueFloat32s(input []float32) []float32 {
	u := make([]float32, 0, len(input))
	m := make(map[float32]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
func uniqueFloat64s(input []float64) []float64 {
	u := make([]float64, 0, len(input))
	m := make(map[float64]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}