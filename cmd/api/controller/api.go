package controller

import (
	"../../pubsub/subscribers"
	"encoding/json"
	"net/http"
	"time"

	_ "../../models"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	Pool *redis.Pool
)

func getParemeter(r *http.Request, name string) string {
	queryValues := r.URL.Query()
	res := queryValues.Get(name)
	return res
}

// AirportList godoc
// @Summary Get list of aiports
// @Description Retrieve aiports list
// @ID get-aiport-list
// @Accept json
// @Success 200 {array} models.Airport
// @Produce  json
// @Router /airports [get]
func AirportList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// init redis database connection
	Pool = subscribers.RedisConnect()

	airports, _ := subscribers.ScanAirports(Pool)
	json.NewEncoder(w).Encode(airports)
}

// CaptorsValues godoc
// @Summary Get captors values
// @Description Retrieve a list of captors values
// @ID get-captors--value-list
// @Param airportId query string false "filter by airport identifier" Enums(NTE, BDX)
// @Param sensorType query string false "filter by sensor type" Enums(temperature, wind, pressure)
// @Param requestStartDate query string false "start date of search"
// @Param requestEndDate query string false "end date of search"
// @Accept json
// @Success 200 {array} models.CaptorValue
// @Produce  json
// @Router /captors [get]
func CaptorsValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// init redis database connection
	Pool = subscribers.RedisConnect()

	airportId := getParemeter(r, "airportId")
	sensorType := getParemeter(r, "sensorType")
	requestStartDate := getParemeter(r, "startDate")
	requestEndDate := getParemeter(r, "endDate")

	startDate, _ := time.Parse("2006-01-02", requestStartDate)
	endDate, _ := time.Parse("2006-01-02", requestEndDate)
	endDate = endDate.Add(time.Hour * 24)

	airports := subscribers.ScanByAirportAndTypeAndDate(Pool, airportId, sensorType, startDate, endDate)
	json.NewEncoder(w).Encode(airports)
}

func Captor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	// init redis database connection
	Pool = subscribers.RedisConnect()

	airports, _ := subscribers.ScanAirports(Pool)
	fmt.Println(airports)
	json.NewEncoder(w).Encode(airports)
}

// SensorIndex godoc
// @Summary Get sensor values
// @Description Retrieve sensor values
// @ID get-sensor-value
// @Param airportId query string false "airport identifier" Enums(NTE, BDX)
// @Param sensorType query string false "sensor type" Enums(temperature, wind, pressure)
// @Param infDate query string false "low date of the search interval : '2009-11-10 23:00:00 UTC'"
// @Param supDate query string false "high date of the search interval : '2009-11-10 23:00:00 UTC'"
// @Accept json
// @Success 200 {array} swaggerModel.CaptorValue
// @Produce  json
// @Router /sensor [get]
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	airportId := getParemeter(r, "airportId")
	sensorType := getParemeter(r, "sensorType")
	sensorId := getParemeter(r, "sensorId")
	fmt.Println(sensorId)

	// format "2009-11-10 23:00:00 UTC"
	infDate := getParemeter(r, "infDate")
	// fmt.Println(infDate)

	t1, err := time.Parse("2006-01-02 15:04:05 UTC", infDate)
	if err != nil {
		fmt.Println("parse error", err.Error())
	}

	// format "2020-10-27 10:04:35 UTC"
	supDate := getParemeter(r, "supDate")
	// fmt.Println(supDate)

	t2, err := time.Parse("2006-01-02 15:04:05 UTC", supDate)
	if err != nil {
		fmt.Println("parse error", err.Error())
	}

	// init redis database connection
	Pool = subscribers.RedisConnect()

	if infDate == "" || infDate == "" {
		sensorValues, _ := subscribers.ScanByAirportAndType(Pool, airportId, sensorType)
		json.NewEncoder(w).Encode(sensorValues)
	} else {
		start := time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, t1.Location())
		end := time.Date(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), 0, t2.Location())
		fmt.Println(start)
		sensorValues := subscribers.ScanByAirportAndTypeAndDate(Pool, airportId, sensorType, start, end)
		json.NewEncoder(w).Encode(sensorValues)
	}
}

// SensorIndexAvg godoc
// @Summary Get sensor values average
// @Description Retrieve the average values of the sensors of each type
// @ID get-average-sensor-value
// @Param airportId query string false "airport identifier" Enums(NTE, BDX)
// @Param date query string false "low date of the search interval : '2009-11-10'"
// @Accept  json
// @Produce  json
// @Success 200 {array} swaggerModel.AverageCaptorValue
// @Router /sensorAverage [get]
func SensorIndexAvg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	airportId := getParemeter(r, "airportId")

	date := getParemeter(r, "date")

	t1, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("parse error", err.Error())
	}

	start := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	end := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 0, t1.Location())

	Pool = subscribers.RedisConnect()

	sensorValues := subscribers.ScanByAirportAndTypeAndDate(Pool, airportId, "", start, end)

	sensorValuesCum := map[string]float64{}
	sensorValuesNb := map[string]float64{}
	sensorValuesAvg := map[string]float64{}

	for _, v := range sensorValues {
		sensorValuesCum[v.Type] += v.Value
		sensorValuesNb[v.Type]++
	}

	for k, v := range sensorValuesCum {
		sensorValuesAvg[k] = v / sensorValuesNb[k]
	}

	w.Write([]byte(fmt.Sprintf(`{"airportId": "%s", "date": "%s", "temperature": "%f", "pressure": "%f", "wind": "%f"}`,
		airportId,
		date,
		sensorValuesAvg["temperature"],
		sensorValuesAvg["pressure"],
		sensorValuesAvg["wind"])))

}
