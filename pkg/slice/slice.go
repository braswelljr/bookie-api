package slice

import "strings"

// Contains return true/false if an element is in a slice or not
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

// Remove element by value from slice
func Remove[T comparable](l []T, item T) []T {
	for i, ele := range l {
		if ele == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// Filter - filter slice by conditions
//
//	@param slice - slice to filter
//	@param conditions - conditions to filter by (array of functions)
//	@return []T - filtered slice
func Filter[T comparable](slice []T, conditions ...func(T) bool) []T {
	// results for the condition
	var result []T

	// if the slice or the conditions are empty
	if len(slice) == 0 || len(conditions) == 0 {
		// return the original slice
		return slice
	}

	// loop through the slice
	for _, item := range slice {
		// match flag
		match := true
		// loop through the conditions
		for _, condition := range conditions {
			// if the condition is not met
			if !condition(item) {
				// set the match flag to false and
				match = false

				// break the loop
				break
			}
		}

		// if the match flag is true
		if match {
			// append the item to the results
			result = append(result, item)
		}
	}
	// if the result is empty
	if len(result) == 0 {
		// return the original slice
		result = slice
	}

	// return the result
	return result
}
