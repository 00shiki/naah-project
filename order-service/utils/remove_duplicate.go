package utils

func RemoveDuplicates[T comparable](nums []T) []T {
	seen := make(map[T]struct{})
	result := []T{}

	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = struct{}{}       // Mark the number as seen
			result = append(result, num) // Add it to the result if it's not a duplicate
		}
	}

	return result
}
