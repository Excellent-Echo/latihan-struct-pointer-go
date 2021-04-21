package structminmax

type statistics struct {
	min, max, avg int
}

type data struct {
	status bool
	score  []int
	stats  statistics
}

func GetStatistics(number ...int) data {
	minScore := number[0]
	maxScore := number[0]
	resScore := 0

	var slice []int
	var avg bool

	// get min and max
	for index, i := range number {
		slice = append(slice, number[index])
		resScore += number[index]

		if i < minScore {
			minScore = i
		}
		if i > maxScore {
			maxScore = i
		}
	}
	average := resScore / len(number)

	// get status data
	if average > 50 {
		avg = true
	} else {
		avg = false
	}

	statsRes := statistics{average, minScore, maxScore}

	dataRes := data{avg, slice, statsRes}

	return dataRes
}
