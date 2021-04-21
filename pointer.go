package main

import (
	"fmt"
)

func GetAverage(defaultNum *[]int, newNum ...int)  {
	var total, avg int

	for _, value := range newNum {
		*defaultNum = append(*defaultNum, value)
	}

	for _, value := range *defaultNum {
		total += value
	}

	lenNum := len(*defaultNum)

	avg = total / lenNum

	fmt.Println(defaultNum, "/", lenNum, "=", avg)
}

func main() {
	number := []int{80, 80, 90, 90}

	GetAverage(&number, 70, 70) 
	GetAverage(&number, 60, 60)
	GetAverage(&number, 50, 40, 80)


}