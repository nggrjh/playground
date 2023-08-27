package datastructures

import "errors"

func binarySearch(array []int, value int) (int, error) {
	lower := 0
	upper := len(array) - 1

	for lower <= upper {
		midPoint := (lower + upper) / 2

		if midValue := array[midPoint]; midValue > value {
			upper = midPoint - 1
		} else if midValue < value {
			lower = midPoint + 1
		} else {
			return midPoint, nil
		}

	}

	return 0, errors.New("not found")
}
