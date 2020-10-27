package main

import (
	"./pubsub/subscribers"
	"time"

	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	Pool *redis.Pool
)

func getParemeter(r *http.Request, name string) string {
	queryValues := r.URL.Query()
	res := queryValues.Get(name)
	return res
}

// http://localhost:8001/sensor?
// airportId=NTE
// &sensorType=wind
// &infDate=2020-10-27+00%3A00%3A00+UTC
// &supDate=2020-10-27+23%3A00%3A00+UTC
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

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

// http://localhost:8001/sensorAverage?
// airportId=NTE
// date=2020-10-23
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

func InitializeRouter() *mux.Router {

	// StrictSlash is true => redirect /airports/ to /airports
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/sensor").Name("SensorIndex").HandlerFunc(SensorIndex)
	router.Methods("GET").Path("/sensorAverage").Name("SensorIndexAvg").HandlerFunc(SensorIndexAvg)

	return router
}

func main() {

	// init redis database connection
	Pool = subscribers.RedisConnect()

	// set routes
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe(":8001", router))
}
