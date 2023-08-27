package datastructures

func bubbleSort(array []int) []int {
	var isSorted bool

	for !isSorted {
		isSorted = true

		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
	}

	return array
}

func selectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		smallestIndex := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[smallestIndex] {
				smallestIndex = j
			}
		}

		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}

	return array
}

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]

		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}

	return array
}
