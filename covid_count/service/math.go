package service

import "math"

func TrendingDirection(averages []float64) string {
	upCount, downCount := 0, 0
	for i, val := range averages {
		if i == len(averages)-1 {
			break
		}
		if averages[i+1] < val {
			upCount++
		}
		if averages[i+1] > val {
			downCount++
		}
	}
	diff := math.Abs(float64(upCount-downCount))
	if diff < 2  {
		return "Steady"
	}
	if upCount > downCount {
		return "Upwards"
	}
	return "Downwards"
}


