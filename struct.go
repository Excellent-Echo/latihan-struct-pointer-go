package main

import "fmt"

type stats struct {
	average, minimum, maximum int
}

type data struct {
	status    bool
	score     []int
	statistic stats
}

func GetStatistic(num ...int) data  {
	minNumber := num[0]
	maxNumber := num[0]
	totalNumber := 0

	var slice []int
	var getAvr bool

	//find min and max score
	for index, value := range num {
		slice = append(slice, num[index])

		totalNumber += num[index]

		if value < minNumber {
			minNumber = value
		} 
		
		if value > maxNumber {
			maxNumber = value
		}
	}

	//get average
	average := totalNumber/len(num)

	//check true or false average is greater than 50
	if average > 50 {
		getAvr = true
	} else {
		getAvr = false
	}

	statsResult := stats{average, minNumber, maxNumber}

	dataResult := data{getAvr,slice, statsResult}

	return dataResult
	
}

func main() {
	result1 := GetStatistic(76, 80, 50, 50, 60, 70)
	result2 := GetStatistic(50, 50, 50)
	result3 := GetStatistic(90, 90, 90, 90)

	fmt.Println(result1,"\n", result2,"\n", result3)
}