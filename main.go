package main

import "fmt"

type Struct struct {
	status    bool
	score     []int
	statistic Statistic
}

type Statistic struct {
	avg int
	min int
	max int
}

func getStatistic(value ...int) {
	s := Struct{}
	s.score = (value)

	s.statistic.max = s.score[0]
	s.statistic.min = s.score[0]

	for i := 0; i < len(s.score); i++ {
		var hasil int
		hasil += s.score[i]
		s.statistic.avg = hasil / len(s.score)

		if s.statistic.max > 50 {
			s.status = true
		} else if s.statistic.min < 50 {
			s.status = false
		}

		if s.score[i] > s.statistic.max {
			s.statistic.max = s.score[i]
		} else if s.score[i] < s.statistic.min {
			s.statistic.min = s.score[i]
		}

	}
	fmt.Printf("Status : %v, \n Score : %v, \n Statistic : { \n Average : %v \n Maximum : %v \n Minimum : %v } \n", s.status, s.score, s.statistic.avg, s.statistic.max, s.statistic.min)

}
func main() {
	getStatistic(76, 80, 50, 50, 60, 70)
	getStatistic(50, 50, 50)
	getStatistic(90, 90, 90, 90)
}
