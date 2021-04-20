package main

import "fmt"

type Variadic struct {
	score     []int // slice
	statistic Statistic
}
type Statistic struct {
	avg int
}

func main() {
	v1 := Variadic{}
	v1.score = append(v1.score, 80, 80, 90, 90)
	v1.score = append(v1.score, 70, 70)
	for i := 0; i < len(v1.score); i++ {
		v1.statistic.avg += v1.score[i] / 5
	}

	fmt.Printf("Score : %v,\nRata-rata : %v}", v1.score, v1.statistic.avg)
}
