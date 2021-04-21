package main

import "fmt"

type Data struct {
	status    bool
	score     []int
	statistic Statistic
}

type Statistic struct {
	avg int
	min int
	max int
}

func GetStatistic(num ...int) {
	var total int
	dt := Data{}
	dt.score = (num)

	dt.statistic.max = dt.score[0]
	dt.statistic.max = dt.score[0]

	for i := 0; i < len(dt.score); i++ {
		total += dt.score[i]

		if dt.score[i] > dt.statistic.max {
			dt.statistic.max = dt.score[i]
		} else if dt.score[i] < dt.statistic.min {
			dt.statistic.min = dt.score[i]
		}

		if dt.statistic.avg > 50 {
			dt.status = true
		} else if dt.statistic.avg <= 50 {
			dt.status = false
		}

		dt.statistic.avg = total / len(dt.score)
	}

	fmt.Printf("\n{\nstatus : %v,\nscore : %v,\nstatistic : {\naverage : %v,\nminimum : %v,\nmaximum : %v,\n}\n}", dt.status, dt.score, dt.statistic.avg, dt.statistic.min, dt.statistic.max)
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
	GetStatistic(50, 50, 50)
	GetStatistic(90, 90, 90, 90)
}
