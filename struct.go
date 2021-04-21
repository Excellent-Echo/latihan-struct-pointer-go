package main

import (
	"fmt"
)

type Statistic struct {
	average int
	minimum int
	maximum int
}
type Report struct {
	status    bool
	score     []int
	statistic Statistic
}

func getStat(data ...int) (avg, min, max int) {
	min = data[0]
	max = data[0]
	sum := 0
	for _, nilai := range data {
		if min > nilai {
			min = nilai
		}
		if max < nilai {
			max = nilai
		}
		sum += nilai
	}

	avg = sum / len(data)

	return
}

func GetStatistic(data ...int) Report {
	report := Report{}
	avg, min, max := getStat(data...)
	stat := Statistic{avg, max, min}

	var status bool
	if avg > 50 {
		status = true
	} else {
		status = false
	}

	report.score = data
	report.status = status
	report.statistic = stat

	fmt.Println(report)
	return report
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
	GetStatistic(50, 50, 50)
	GetStatistic(90, 90, 90, 90)
}
