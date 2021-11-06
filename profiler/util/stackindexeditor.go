package util

func GetPreviousIndex(j int) int {
	j--

	if j < 0 {
		return 0
	} else {
		return j
	}
}

func GetNextIndex(j int, maxValue int) int {
	j++

	if j < maxValue {
		return j
	} else {
		return maxValue
	}
}

func GetCurrentStackIndex(j int) int {
	return GetPreviousIndex(j)
}
