package responses

type CountyWeekAverage struct {
	County string `json:"county"`
	Averages map[string]float64 `json:"averages"`
	TrendingDirection string `json:"trending_direction"`
	TrendingRatio float64 `json:"trending_ratio"`
}
