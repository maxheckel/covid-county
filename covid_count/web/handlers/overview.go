package handlers

import (
	"fmt"
	"github.com/maxheckel/covid_county/covid_count/domain"
	"github.com/maxheckel/covid_county/covid_count/repository"
	"github.com/maxheckel/covid_county/covid_count/service"
	"github.com/maxheckel/covid_county/covid_count/web"
	"github.com/maxheckel/covid_county/covid_count/web/responses"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var counties = [88]string{
	"Ross",
	"Summit",
	"Brown",
	"Warren",
	"Clermont",
	"Miami",
	"Licking",
	"Jefferson",
	"Crawford",
	"Stark",
	"Monroe",
	"Ashtabula",
	"Jackson",
	"Huron",
	"Knox",
	"Auglaize",
	"Pike",
	"Union",
	"Sandusky",
	"Vinton",
	"Erie",
	"Lorain",
	"Franklin",
	"Ottawa",
	"Wyandot",
	"Harrison",
	"Pickaway",
	"Wood",
	"Fulton",
	"Medina",
	"Adams",
	"Columbiana",
	"Seneca",
	"Washington",
	"Wayne",
	"Perry",
	"Logan",
	"Defiance",
	"Greene",
	"Van Wert",
	"Allen",
	"Henry",
	"Hancock",
	"Lawrence",
	"Meigs",
	"Ashland",
	"Geauga",
	"Lake",
	"Darke",
	"Montgomery",
	"Tuscarawas",
	"Butler",
	"Hardin",
	"Cuyahoga",
	"Delaware",
	"Morgan",
	"Noble",
	"Highland",
	"Putnam",
	"Hocking",
	"Marion",
	"Richland",
	"Shelby",
	"Mahoning",
	"Athens",
	"Muskingum",
	"Belmont",
	"Clark",
	"Champaign",
	"Holmes",
	"Carroll",
	"Coshocton",
	"Preble",
	"Morrow",
	"Gallia",
	"Fairfield",
	"Madison",
	"Fayette",
	"Williams",
	"Trumbull",
	"Clinton",
	"Portage",
	"Mercer",
	"Hamilton",
	"Paulding",
	"Guernsey",
	"Lucas",
	"Scioto",
}

const isUpdatingCacheKey = "writing"
const daysBack = 21
const averagesKey = "averages"
const interval = 7


type Overview struct {
	Manager *repository.Manager
	Cache   *service.Cache
}

func (o Overview) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isUpdating, err := o.isUpdating()
	if isUpdating {
		web.WriteJSON(w, 503, responses.IsUpdating{IsUpdating: true})
		return
	}

	if output, found := o.Cache.Get(averagesKey); found {
		fmt.Println("exists in cache")
		web.WriteJSON(w, 200, output)
		return
	}
	output, err := o.getSevenDayAverages()
	fmt.Println("adding to cache")
	o.Cache.Set(averagesKey, output, 1)
	if err != nil {
		web.WriteJSONError(w, r, web.UnexpectedError(err.Error()))
		return
	}

	web.WriteJSON(w, 200, output)
}


func (o Overview) getSevenDayAverages() ([]responses.CountyWeekAverage, error){
	caseResponse, err := o.Manager.Cases().GetAllCasesForDays(daysBack)
	if err != nil {
		return nil, err
	}
	caseResponse = populateZeroDays(caseResponse)

	var output []responses.CountyWeekAverage
	for _, county := range counties {

		res := responses.CountyWeekAverage{
			County:   county,
			Averages: []float64{},
		}
		countyCases := []*domain.DailyCases{}
		for _, cases := range caseResponse {
			if cases.County == county {
				countyCases = append(countyCases, cases)
			}
		}
		sort.SliceStable(countyCases, func(i, j int) bool {
			return countyCases[i].Date.After(countyCases[j].Date)
		})

		// calculate the 7 day rolling average
		for i := 0; i < len(countyCases)-interval; i++ {
			var sum = 0
			for j := 0; j < interval; j++ {
				sum += countyCases[i+j].Count
			}
			res.Averages = append(res.Averages, float64(sum/interval))
		}
		res.TrendingDirection = service.TrendingDirection(res.Averages)
		output = append(output, res)
	}
	return output, nil
}



func populateZeroDays(caseResponse []*domain.DailyCases) []*domain.DailyCases {
	var buckets = make(map[string][]*domain.DailyCases)
	dateIterator := time.Now().AddDate(0,0,(daysBack+6)*-1)
	for i := 0; i < daysBack; i++ {
		dateString := DateToString(dateIterator)
		buckets[dateString] = []*domain.DailyCases{}
		for _, cases := range caseResponse {
			if DateToString(cases.Date) == dateString {
				buckets[dateString] = append(buckets[dateString], cases)
			}
		}
		dateIterator = dateIterator.AddDate(0,0,1)
	}
	populateEmptyDates(buckets)
	var resp []*domain.DailyCases
	for _, day := range buckets{
		resp = append(resp, day...)
	}
	return resp
}

func populateEmptyDates(buckets map[string][]*domain.DailyCases) {
	for date, cases := range buckets {
		if len(cases) == 0 {
			continue
		}
		if len(cases) == len(counties) {
			continue
		}
		var dailyCounties []string
		for _, day := range cases {
			dailyCounties = append(dailyCounties, day.County)
		}
		missingCounties := countyDifference(counties, dailyCounties)
		for _, missingCounty := range missingCounties {
			buckets[date] = append(buckets[date], &domain.DailyCases{
				County: missingCounty,
				Date:   cases[0].Date,
				Count:  0,
			})
		}
	}
}

// countyDifference returns the elements in `a` that aren't in `b`.
func countyDifference(a [88]string, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func DateToString(date time.Time) string {
	y,m,d := date.Date()
	return strconv.Itoa(y)+"-"+m.String()+"-"+strconv.Itoa(d)
}


func (o Overview) isUpdating() (bool, error) {
	var isUpdating bool
	var err error
	res, found := o.Cache.Get(isUpdatingCacheKey)
	if !found {
		isUpdating, err = o.Manager.IsUpdating().IsUpdating()
		o.Cache.Set(isUpdatingCacheKey, isUpdating, 60)
	}
	if res, set := res.(bool); set {
		isUpdating = res
	}
	return isUpdating, err
}
