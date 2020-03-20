package sxint32

func Contains(slice []int32, item int32) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func Indentical(a, b *int32) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		if *a == *b {
			return true
		}
	}
	return false
}

func Equal(a,b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	diff := make(map[int32]int, len(a))
	for _, _x := range a {
		diff[_x]++
	}
	for _, _y := range b {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

func Unique(s []int32) []int32 {
	seen := make(map[int32]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

//returns missing elements in b compared to a
func Missing(a, b []int32) []int32 {
	// create map with length of the 'a' slice
	ma := make(map[int32]struct{}, len(a))
	var diffs []int32
	// Convert first slice to map with empty struct (0 bytes)
	for _, ka := range a {
		ma[ka] = struct{}{}
	}
	// find missing values in a
	for _, kb := range b {
		if _, ok := ma[kb]; !ok {
			diffs = append(diffs, kb)
		}
	}
	return diffs
}

//returns uncommon elements
func Unmatched(a, b []int32) []int32 {
	var diff []int32
	// Loop two times, first to find a strings not in b,
	// second loop to find b strings not in a
	for i := 0; i < 2; i++ {
		for _, s1 := range a {
			found := false
			for _, s2 := range b {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			a, b = b, a
		}
	}

	return diff
}