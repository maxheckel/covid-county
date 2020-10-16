package service

import (
	"fmt"
	"math"
	"time"
)

func TrendingDirection(averages map[string]float64, county string) (string, float64) {
	upCount, downCount := float64(0), float64(0)
	total := float64(0)
	for k, v := range averages{
		date,_ := time.Parse("2006-01-02", k)
		if nextDayVal, ok := averages[date.AddDate(0, 0, 1).Format("2006-01-02")]; ok {
			total+=v
			if nextDayVal > v {
				upCount+=nextDayVal-v
				continue
			}
			if nextDayVal < v {
				downCount+=v-nextDayVal
				continue
			}
		}
	}

	diff := math.Abs(upCount - downCount)
	var resString string

	avg := total/float64(len(averages))
	if county == "Monroe"{
		fmt.Println(diff, avg)
	}
	// Deals with counties that are still trending up but have a lower coefficient
	if (diff < avg && avg < 1) || (diff < 1 && avg > 1){
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


