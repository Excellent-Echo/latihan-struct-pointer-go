package main

import "fmt"

func GetAverage(data *[]int, avg... int) int{
	var result int

	for _, value := range avg {
		*data = append(*data, value)
	}
	fmt.Println("number : ", *data)
	length := len(*data)
	//fmt.Println("ini length", length)

	newData := 0
	for _, value := range *data {
		newData += value
	}
	fmt.Println("Total : ", newData)
	fmt.Println("------------------------------")

	result = newData / length
	return result
}

func main()  {
	number := []int{80,80,90,90}

	avg := GetAverage(&number , 70,70)
	avg2 := GetAverage(&number, 60,60)
	avg3 := GetAverage(&number, 50,40,80)
	fmt.Println("Rata-rata :", avg)
	fmt.Println("Rata-rata :", avg2)
	fmt.Println("Rata-rata :", avg3)
}

