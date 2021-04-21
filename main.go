package main

import (
	"fmt"
	"structPointer/latihanpointer"
	"structPointer/structminmax"
)

func main() {
	fmt.Println("Min & Max Struct")
	res := structminmax.GetStatistics(76, 80, 50, 50, 60, 70)
	fmt.Println(res)

	fmt.Println("================\nPointer")

	number := []int{80, 80, 90, 90}
	avg := latihanpointer.GetAverage(&number, 70, 70)  // 80+80+90+90+70+70 / 6 = 80
	avg2 := latihanpointer.GetAverage(&number, 60, 60) // 80+80+90+90+70+70+60+60 / 8 = 75
	avg3 := latihanpointer.GetAverage(&number, 50, 40, 80)

	fmt.Println(avg, "\n", avg2, "\n", avg3)
}
