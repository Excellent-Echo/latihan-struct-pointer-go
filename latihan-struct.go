package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Data struct {
	Status    bool
	Score     []int
	Statistic Stats
}

type Stats struct {
	Average, Minimum, Maximum int
}

func GetStatistic(scores ...int) string {
	var avg, max, sum int
	min := math.MaxInt64
	var check bool
	var scoreSlice []int

	// compute average, minimum, maximum score, and append score to slice
	for _, n := range scores {
		sum += n

		if n > max {
			max = n
		}
		if n < min {
			min = n
		}

		scoreSlice = append(scoreSlice, n)
	}
	avg = sum / len(scores)

	// set status(check) rely on average score
	if avg > 50 {
		check = true
	} else {
		check = false
	}

	// assign struct value and convert to JSON
	stats := Stats{Average: avg, Minimum: min, Maximum: max}
	data := Data{Status: check, Score: scoreSlice, Statistic: stats}
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(result))
	return ""
}

func main() {
	GetStatistic(76, 80, 50, 50, 60, 70)
	GetStatistic(50, 50, 50)
	GetStatistic(90, 90, 90, 90)
}
