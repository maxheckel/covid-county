package responses

type CountyWeekAverage struct {
	County string `json:"county"`
	Averages []float64 `json:"averages"`
}
