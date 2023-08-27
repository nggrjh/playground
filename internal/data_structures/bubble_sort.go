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
