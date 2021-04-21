package helper

func Change(original *[]int, value []int) {
	*original = value
}

func GetAverage(arr []int, num ...int) ([]int, int) {
	slice := arr
	total := 0

	for i := range num {
		slice = append(slice, num[i])
	}

	for i := range slice {
		total += slice[i]
	}

	avg := total / len(slice)

	return slice, avg
}
