package main

import (
	"fmt"
	"structPointer/structminmax"
)

func main() {
	fmt.Println("Min & Max Struct")
	res := structminmax.GetStatistics(76, 80, 50, 50, 60, 70)
	fmt.Println(res)
}
