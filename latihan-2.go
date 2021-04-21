package main

import "fmt"

func GetAverage(data *[]int, avg... int) int{
	var result int

	lengthavg := len(avg)
	newAvg := 0
	for _, value := range avg {
		newAvg += value
	}


	lengthdata := len(*data)
	newData := 0
	for _, value := range *data {
		newData += value
	}
	finalLength := lengthavg + lengthdata


	result = (newAvg + newData) / finalLength
	return result
}

func main() {
	number := []int{80,80,90,90}
	avg := GetAverage(&number , 70,70)
	//avg2 := GetAverage(&number, 60,60)
	//avg3 := GetAverage(&number, 50,40,80)

	fmt.Println(avg)
	//fmt.Println(avg2)
	//fmt.Println(avg3)
}
