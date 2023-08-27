package playground

import "fmt"

// Source: https://www.youtube.com/live/_fkdalNPcw0
func BoredEngineerWorkingHours(tasks []int, gap int) int {
	lastIdx := make(map[int]int)
	workingHours := len(tasks)

	for idx, task := range tasks {

		if _, ok := lastIdx[task]; ok {
			if diff := idx - lastIdx[task]; diff <= gap {
				workingHours += gap - diff + 1
			}
		}

		lastIdx[task] = idx
	}

	return workingHours
}

// FIXME
func BoredEngineerRearrangeWorkingHours(tasks []int, gap int) int {
	taskCount := make(map[int]int)
	for _, task := range tasks {
		taskCount[task] += 1
	}

	uniqueCount := len(taskCount)
	if uniqueCount == 1 {
		return len(tasks) + ((len(tasks) - 1) * gap)
	}

	fmt.Println("UNIQUE", uniqueCount)
	
	var sequence []int
	firstIdx := make(map[int]int)
	for task := range taskCount {
		if _, ok := firstIdx[task]; !ok {
			firstIdx[task] = len(sequence)
		}

		sequence = append(sequence, task)

		fmt.Println("SEQUENCE", sequence)
		fmt.Println("INDEX", firstIdx)
	}

	
	fmt.Println("COUNT", taskCount)

	return 0
}
