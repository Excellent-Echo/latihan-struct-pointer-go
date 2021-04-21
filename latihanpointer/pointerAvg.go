package latihanpointer

func GetAverage(prevNum *[]int, nextNum ...int) int {
	var result int

	for _, value := range nextNum {
		*prevNum = append(*prevNum, value)
	}

	for _, val := range *prevNum {
		result += val
	}

	average := result / len(*prevNum)

	return average
}
