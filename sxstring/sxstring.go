package sxstring

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func Indentical(a, b *string) bool {
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

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := make(map[string]int, len(a))
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

func Unique(s []string) []string {
	seen := make(map[string]struct{}, len(s))
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
func Missing(a, b []string) []string {
	// create map with length of the 'a' slice
	ma := make(map[string]struct{}, len(a))
	var diffs []string
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
func Unmatched(a, b []string) []string {
	var diff []string
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