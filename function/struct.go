package function

import "fmt"

func GetStatistic(value ...int) {
	data := Data{}
	average := 0
	data.statistic.maximum = value[0]
	data.statistic.minimum = value[0]
	// max := data.score[0]
	// min := data.score[0]

	//looping untuk append value ke data score
	for _, val := range value {
		data.score = append(data.score, val)
	}

	//looping average and max min
	for _, val := range data.score {
		average += val
		if val > data.statistic.maximum {
			data.statistic.maximum = val
		} else if val < data.statistic.minimum {
			data.statistic.minimum = val
		}
	}
	data.statistic.average = average / len(data.score)

	//condition status
	if data.statistic.average > 50 {
		data.status = true
	} else {
		data.status = false
	}

	fmt.Printf("status : %v\n score : %v\n Statistic : {\n average : %v,\n minimum : %v,\n maximum : %v\n }\n\n", data.status, data.score, data.statistic.average, data.statistic.minimum, data.statistic.maximum)
}

type Data struct {
	status    bool
	score     []int
	statistic Statistic
}

type Statistic struct {
	average int
	maximum int
	minimum int
}
