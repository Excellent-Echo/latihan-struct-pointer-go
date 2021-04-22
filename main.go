package main

import (
	"fmt"
	"functionGo/function"
)

func main() {

	var number = []int{80, 80, 90, 90}
	// fmt.Println(GetAverage(number, 70))

	fmt.Println("Latihan struct")
	function.GetStatistic(76, 70, 20, 70, 10, 30, 90, 5)

	fmt.Println("Latihan pointer")
	avg := function.GetAverage(&number, 70, 70)
	avg2 := function.GetAverage(&number, 60, 60)
	avg3 := function.GetAverage(&number, 50, 40, 80)
	fmt.Println(avg)
	fmt.Println(avg2)
	fmt.Println(avg3)

}

// package main

// import "fmt"

// func GetAverage(data *[]int, avg ...int) int {
// 	var result int

// 	for _, value := range avg {
// 		*data = append(*data, value)
// 	}
// 	fmt.Println("ini newAvg", *data)
// 	length := len(*data)
// 	fmt.Println("ini length", length)

// 	newData := 0
// 	for _, value := range *data {
// 		newData += value
// 	}
// 	fmt.Println("ini newData", newData)

// 	result = newData / length
// 	return result
// }

// func main() {
// 	number := []int{80, 80, 90, 90}

// 	avg := GetAverage(&number, 70, 70)
// 	avg2 := GetAverage(&number, 60, 60)
// 	avg3 := GetAverage(&number, 50, 40, 80)
// 	fmt.Println(avg)
// 	fmt.Println(avg2)
// 	fmt.Println(avg3)
// }
