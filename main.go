package main

import "fmt"

type Statistic struct {
	average float32
	minimum int
	maximum int
}

type Respon struct {
	status bool
	score  []int
	Statistic
}

func GetStatistic(numbers ...int) {
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
	fmt.Println(avg, max, min)
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
}
