package function

func GetAverage(number *[]int, new ...int) int {

	for _, val := range new {
		*number = append(*number, val)
	}
	len := len(*number)

	average := 0
	for _, val := range *number {
		average += val
	}

	result := average / len
	return result
}
