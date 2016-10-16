package quicksort

// Implement Quicksort

func qsort(data []int) {
	if len(data) < 2 {
		return
	}

	futurePivot := 0
	pivot := len(data) - 1 // This is not a great choice (i.e. O(n^2)) for sorted slices.

	for i := range data {
		// If the element is less than the pivot
		if data[i] < data[pivot] {
			// Move it before the future pivot
			data[futurePivot], data[i] = data[i], data[futurePivot]
			futurePivot++ // futurePivot is always at or before i
		}
	}

	// We now know everything ahead of the future pivot is smaller than or equal
	// to the current pivot, and everything between it and the current pivot is
	// larger. So let's put the current pivot where the future pivot was.
	data[futurePivot], data[pivot] = data[pivot], data[futurePivot]

	// Further sort everything below the pivot (which is still called futurePivot)
	qsort(data[:futurePivot])

	// Further sort everything above the pivot
	qsort(data[futurePivot+1:])
}
