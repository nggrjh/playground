package playground

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
