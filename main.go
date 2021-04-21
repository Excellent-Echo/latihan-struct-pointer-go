package main

import "fmt"

func main() {
    
    getStatistic(76,70,20,70,10, 30, 90, 5)
    
}

type Data struct {
    status bool
    score []int
    statistic Statistic
}

type Statistic struct {
    average int
    maximum int
    minimum int
}

func getStatistic(value ...int) {
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
        } else if val < data.statistic.minimum{
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
    
    
    
    
    fmt.Println(data)
    
}
