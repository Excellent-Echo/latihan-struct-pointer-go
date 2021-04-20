package main

import (
	"fmt"
	"main/helper"
)

func main() {
	//Latihan Struct

	result := helper.GetStatistic(76, 80, 50, 50, 60, 70)
	// result := GetStatistic(50, 50, 50)
	// result := GetStatistic(90, 90, 90, 90)

	fmt.Println(result)
	fmt.Printf("========================================")
}
