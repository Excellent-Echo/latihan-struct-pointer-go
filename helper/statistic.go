package helper

type Statistic struct {
	avg, min, max int
}

type DataStatistic struct {
	status    bool
	score     []int
	statistic Statistic
}

func GetStatistic(num ...int) DataStatistic {
	min := num[0]
	max := num[0]
	total := 0
	var slice []int
	var check bool

	for i, v := range num {
		slice = append(slice, num[i])

		total += num[i]

		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	avg := total / len(num)

	if avg > 50 {
		check = true
	} else {
		check = false
	}

	data := Statistic{avg, min, max}
	dataStatistic := DataStatistic{check, slice, data}

	return dataStatistic
}
