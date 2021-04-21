package main

import "fmt"

type Variadic struct {
	status    bool  // false
	score     []int // slice
	statistic Statistic
}
type Statistic struct {
	avg int
	min int
	max int
}

func GetStatistic(num ...int) {
	var total int
	v1 := Variadic{}
	v1.score = (num)

	// max min
	v1.statistic.max = v1.score[0]
	v1.statistic.min = v1.score[0]

	for i := 0; i < len(v1.score); i++ {

		total += v1.score[i] // rata rata

		if v1.score[i] > v1.statistic.max {
			v1.statistic.max = v1.score[i]
		} else if v1.score[i] < v1.statistic.min {
			v1.statistic.min = v1.score[i]
		}

		if v1.statistic.avg > 50 {
			v1.status = true
		} else if v1.statistic.avg <= 50 {
			v1.status = false
		}

		v1.statistic.avg = total / len(v1.score)
	}
	fmt.Printf("Status : %v,\nScore : %v,\nStatistic : {\naverage : %v\nminimum : %v\nmaximum : %v\n}", v1.status, v1.score, v1.statistic.avg, v1.statistic.min, v1.statistic.max)
}
func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
	GetStatistic(50, 50, 50)
	GetStatistic(90, 90, 90, 90)
}
