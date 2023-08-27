package datastructures

func selectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		smallestIndex := i

		for j := i+1; j < len(array); j++ {
			if array[j] < array[smallestIndex] {
				smallestIndex = j
			}
		}

		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}

	return array
}
