package main

import (
	"fmt"
)

type Data struct {
	status bool
	score []int
	statistic Statistic
}

type Statistic struct {
	average int
	minimum int
	maximum int
}

func GetStatistic(numbers... int) {
	length := len(numbers)
	avarage := 0
	total := 0
	for _, value := range numbers {
		total += value
		division := total / length
		avarage = division
	}

	var checkBool bool

	if avarage > 50 {
		checkBool = true
	} else if avarage < 50 {
		checkBool = false
	}

	minimum := numbers[0]
	maximum := numbers[0]
	for _, element := range numbers {
		if element < minimum {
			minimum = element
		}
		if element > maximum {
			maximum = element
		}
	}

	addData := Data{}

	addData = Data{
		status: checkBool,
		score: numbers,
	}
	addData.statistic = Statistic{
		average: avarage,
		minimum: minimum,
		maximum: maximum,
	}
	fmt.Println(addData)
}


func main(){
	GetStatistic(76,80,50,50,60,70)
	fmt.Println("----------------------")

	GetStatistic(50,50,50)
	fmt.Println("----------------------")

	GetStatistic(90,90,90,90)
	fmt.Println("----------------------")
}
