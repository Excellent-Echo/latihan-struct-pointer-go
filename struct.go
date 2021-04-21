package main

import "fmt"

type dataVariadic struct {
	status    bool
	score     []int
	statistic Statistic
}
type Statistic struct {
	average int
	min     int
	max     int
}

func GetStatistic(num ...int) {
	var total int
	data := dataVariadic{}
	data.score = (num)

	data.statistic.max = data.score[0]
	data.statistic.min = data.score[0]

	for i := 0; i < len(data.score); i++ {

		total += data.score[i]

		if data.score[i] > data.statistic.max {
			data.statistic.max = data.score[i]
		} else if data.score[i] < data.statistic.min {
			data.statistic.min = data.score[i]
		}

		if data.statistic.average > 50 {
			data.status = true
		} else if data.statistic.average <= 50 {
			data.status = false
		}

		data.statistic.average = total / len(data.score)
	}
	fmt.Printf("\nStatus : %v,\nScore : %v,\nStatistic : {\naverage : %v\nminimum : %v\nmaximum : %v\n}\n\n", data.status, data.score, data.statistic.average, data.statistic.min, data.statistic.max)
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
	GetStatistic(50, 50, 50)
	GetStatistic(90, 90, 90, 90)
}
