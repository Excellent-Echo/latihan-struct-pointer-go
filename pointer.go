package main

import (
	"fmt"
)

func main() {
	number := []int{80, 80, 90, 90}

	avg := GetAverage(&number, 70, 70)
	avg2 := GetAverage(&number, 60, 60)
	avg3 := GetAverage(&number, 50, 40, 80)

	fmt.Println(avg)
	fmt.Println(avg2)
	fmt.Println(avg3)
}

func GetAverage(oldNumber *[]int, newNumber ...int) int {
	var sum, avg int

	for _, val := range newNumber {
		*oldNumber = append(*oldNumber, val)
	}

	for _, n := range *oldNumber {
		sum += n
	}

	avg = sum / len(*oldNumber)

	return avg
}
