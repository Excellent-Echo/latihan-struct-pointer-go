package main

import "fmt"

//Task 1
type Statistic struct {
	average int
	minimum int
	maximum int
}

type Respon struct {
	status bool
	score  []int
	Statistic
}

func GetStatistic(numbers ...int) Respon {
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
	res := Respon{}
	res.status = true
	res.score = numbers
	res.average = avg
	res.maximum = max
	res.minimum = min

	return res
}

func main() {
	// Task 1
	stat := GetStatistic(76, 80, 50, 50, 60, 70)

	//Task 2
	var number = []int{80, 80, 90, 90}

	avg := GetAverage(&number, 70, 70)

	fmt.Println(stat, avg)
}

//Task 2
func GetAverage(oldNums *[]int, newNums ...int) int {
	tmp := make([]int, 0)
	var sum int
	for _, num := range *oldNums {
		tmp = append(tmp, num)
	}
	for _, num := range newNums {
		tmp = append(tmp, num)
	}
	for _, num := range tmp {
		//fmt.Println(num)
		sum += num
	}

	avg := sum / len(tmp)
	return avg
}
