package main

import "fmt"

func GetAverage(data []int, avg... int) {

	lengthavg := len(avg)
	lengthdata := len(data)
	finalLength := lengthavg + lengthdata

	newAvg := 0
	for _, value := range avg {
		newAvg += value
	}

	newData := 0
	for _, value := range data {
		newData += value
	}

	result := (newAvg + newData) / finalLength
	fmt.Println(result)

}

func main() {
	var number = []int{80,80,90,90}
	GetAverage(number, 70,70)
	GetAverage(number, 60,60)
	GetAverage(number, 50,40,80)
}
