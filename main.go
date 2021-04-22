package main

type Statistic struct {
	average int
	minimum int
	maximum int
}

type Respon struct {
	status bool
	score  []int
	Statistic
}

func GetStatistic(numbers ...int) map[string]Respon {
	var (
		sum, max, min, avg int
	)
	max = 0
	min = 1000

	for _, number := range numbers {
		sum += number
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}
	avg = sum / len(numbers)
	res := Respon{}
	res.status = true
	res.score = numbers
	res.average = avg
	res.maximum = max
	res.minimum = min

	result := map[string]Respon{
		"response": res,
	}

	return result
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
}
