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

	s1 := Statistic{}

	for i := 0; i < len(v1.score); i++ {
		s1.avg += v1.score[i] / 5
	}

	fmt.Println(s1.avg)
}
