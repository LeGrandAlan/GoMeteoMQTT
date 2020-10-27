package main

import (
	"./pubsub/subscribers"

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

// http://localhost:8001/sensor?airportId=NTE&sensorType=wind
func SensorIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	airportId := getParemeter(r, "airportId")
	fmt.Println(airportId)

	sensorType := getParemeter(r, "sensorType")
	fmt.Println(sensorType)

	sensorId := getParemeter(r, "sensorId")
	fmt.Println(sensorId)

	infDate := getParemeter(r, "lowDate")
	fmt.Println(infDate)

	supDate := getParemeter(r, "higDate")
	fmt.Println(supDate)

	sensorValues, _ := subscribers.ScanByAirportAndType(Pool, airportId, sensorType)

	/*
	    pathParams := mux.Vars(r)

	    var sensorType = -1
		var err error
		if val, ok := pathParams["type"];
		ok {
			fmt.Println(val)
			sensorType, err = strconv.Atoi(val)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	*/

	json.NewEncoder(w).Encode(sensorValues)
}

// http://localhost:8001/sensorAverage?airportId=NTE
func SensorIndexAvg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	airportId := getParemeter(r, "airportId")
	fmt.Println(airportId)

	date := getParemeter(r, "date")
	fmt.Println(date)

	sensorValues, _ := subscribers.ScanByAirportAndType(Pool, airportId, "")

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

	w.Write([]byte(fmt.Sprintf(`{"temperature": "%f", "pressure": "%f", "wind": "%f"}`,
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
