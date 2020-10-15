package responses

type CountyWeekAverage struct {
	County string `json:"county"`
	Averages []float64 `json:"averages"`
	TrendingDirection string `json:"trending_direction"`
}
