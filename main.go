package main

import (
	"fmt"
	"main/helper"
)

func main() {
	//Latihan Struct

	statistic := helper.GetStatistic(76, 80, 50, 50, 60, 70)
	// result := GetStatistic(50, 50, 50)
	// result := GetStatistic(90, 90, 90, 90)

	fmt.Println(statistic)
	fmt.Printf("========================================\n")

	//Latihan Pointer

	number := []int{80, 80, 90, 90}

	avgSlice, avg := helper.GetAverage(number, 70, 70)
	fmt.Println(avg)

	helper.Change(&number, avgSlice)

	avgSlice, avg2 := helper.GetAverage(number, 60, 60)
	fmt.Println(avg2)

	helper.Change(&number, avgSlice)

	avgSlice, avg3 := helper.GetAverage(number, 50, 40, 80)
	fmt.Println(avg3)

}
