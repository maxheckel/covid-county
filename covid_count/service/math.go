package service

import (
	"math"
	"time"
)

func TrendingDirection(averages map[string]float64, county string) (string, float64) {
	upCount, downCount := 0, 0
	for k, v := range averages{
		date,_ := time.Parse("2006-01-02", k)
		if nextDayVal, ok := averages[date.AddDate(0, 0, 1).Format("2006-01-02")]; ok {
			if nextDayVal > v {
				upCount++
				continue
			}
			if nextDayVal < v {
				downCount++
				continue
			}
		}
	}

	diff := math.Abs(float64(upCount-downCount))
	var resString string


	if diff < float64(2) {
		return "Steady", diff
	}
	if upCount > downCount {
		resString = "Upwards"
	}
	if upCount < downCount {
		resString = "Downwards"
	}



	return resString, diff
}


