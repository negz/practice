package editdistance

// Determine the (Levenshtein) edit distance between two strings.

// O(nm) where n is len(short) and m is len(long)-len(short)
func distance(i, j string) int {
	// The strings are the same. There's no edit distance.
	if i == j {
		return 0
	}

	short, long := shorterFirst([]rune(i), []rune(j))

	// If they're already the same length the edit distance is len - matches
	if len(short) == len(long) {
		return len(short) - matches(short, long)
	}

	// Edit distance:
	// - Matches in the aligned position
	// - Plus elements stripped from head of longer slice
	// - Plus elements remaining at tail of longer slice
	// i.e. Length of longer slice - matches()

	distance := len(long) // Worst case distance

	// Shift the long down an element until it's the size of short
	for i := 0; i <= len(long)-len(short); i++ {
		// Determine the edit distance at this alignment
		d := len(long) - matches(short, long[i:len(long)])
		// Save it if it's the shortest edit distance we've seen
		if d < distance {
			distance = d
		}
	}
	return distance
}

// O(n) where n is len(short)
func matches(short, long []rune) int {
	m := 0
	for i, c := range short {
		if long[i] == c {
			m++
		}
	}
	return m
}

func shorterFirst(i, j []rune) ([]rune, []rune) {
	if len(i) < len(j) {
		return i, j
	}
	return j, i
}
