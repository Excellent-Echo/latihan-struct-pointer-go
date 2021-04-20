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

// func createStruct()  {

// }
func main() {
	v1 := Variadic{}
	v1.score = append(v1.score, 76, 80, 50, 50, 60, 70)

	// v1.statistic.max = v1.score[0]
	// v1.statistic.min = v1.score[0]

	for i := 0; i < len(v1.score); i++ {
		v1.statistic.avg += v1.score[i] / 5
		// if v1.score[i] > max {
		// 	v1.statistic.max = v1.score[i]
		// } else if v1.score[i] < min {
		// 	v1.statistic.min = v1.score[i]
		// }
		if v1.statistic.avg > 50 {
			v1.status = true
		} else if v1.statistic.avg <= 50 {
			v1.status = false
		}
	}

	fmt.Println(v1.status, v1.score, v1.statistic.avg)
}
