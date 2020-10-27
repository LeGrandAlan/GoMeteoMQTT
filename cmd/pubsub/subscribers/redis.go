package subscribers

import (
	"../../models"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"time"
)

func RedisConnect() *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}

func Ping(pool *redis.Pool) error {

	conn := pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func HGetAllCaptors(pool *redis.Pool, hash string) (CaptorValue, error) {

	conn := pool.Get()
	defer conn.Close()

	var data CaptorValue
	datas, err := redis.Values(conn.Do("HGETALL", hash))
	if err != nil {
		return data, fmt.Errorf("error getting all keys/values for %s: %v", hash, err)
	}

	var array []string

	for i, s := range datas {
		data := string(s.([]byte))
		switch i {
		case 1, 3, 5, 7, 9, 11, 13:
			array = append(array, data)
			break
		}
	}

	fetchedCaptor := MakeFromRedisArray(array)

	return fetchedCaptor, err

}

func HGetAllAirports(pool *redis.Pool, hash string) (models.Airport, error) {

	conn := pool.Get()
	defer conn.Close()

	var data models.Airport
	datas, err := redis.Values(conn.Do("HGETALL", hash))
	if err != nil {
		return data, fmt.Errorf("error getting all keys/values for %s: %v", hash, err)
	}

	airportMap := make(map[string]interface{})
	for i, s := range datas {
		data := string(s.([]byte))
		switch i {
		case 1:
			airportMap["Id"] = data
			break
		case 3:
			airportMap["Name"] = data
			break
		case 5:
			airportMap["City"] = data
			break
		}
	}

	fetchedAirport := models.AirportMapper(airportMap)

	fmt.Println(fetchedAirport)

	return fetchedAirport, err

}

func HSetAirportIfNoExists(pool *redis.Pool, airport models.Airport) error {

	conn := pool.Get()
	defer conn.Close()

	var keysValues []interface{}
	keysValues = append(keysValues, "goMeteoMQTT:airport:"+airport.Id)

	values := reflect.ValueOf(airport)
	num := values.NumField()

	for i := 0; i < num; i++ {
		key := values.Type().Field(i).Name
		value := values.Field(i)

		keysValues = append(keysValues, key)
		keysValues = append(keysValues, value)
	}

	data, err1 := conn.Do("EXISTS", airport.Id)

	if err1 != nil {
		return fmt.Errorf("error setting hash keys %v", err1)
	}

	var err2 error
	if int(data.(int64)) == 0 {
		_, err2 := conn.Do("HMSET", keysValues...)

		if err2 != nil {
			return fmt.Errorf("error setting hash keys %v", err1)
		}
	}

	return err2

}

func HSetCaptorValue(captorValue CaptorValue, idPrefix, id string) error {

	conn := Pool.Get()
	defer conn.Close()

	var keysValues []interface{}
	keysValues = append(keysValues, idPrefix+id)

	values := reflect.ValueOf(captorValue)
	num := values.NumField()

	for i := 0; i < num; i++ {
		key := values.Type().Field(i).Name
		value := values.Field(i)

		keysValues = append(keysValues, key)
		keysValues = append(keysValues, value)
	}

	_, err := conn.Do("HMSET", keysValues...)

	// _, err2 := conn.Do("SADD", "goMeteoMQTT:all-captorValues", id)

	// _, err3 := conn.Do("ZADD", "goMeteoMQTT:dateIndex", captorValue.StringDate, id)

	if err != nil {
		return fmt.Errorf("error setting hash keys %v", err)
	}

	return err

}

func ScanAirports(pool *redis.Pool) ([]models.Airport, error) {

	conn := pool.Get()
	defer conn.Close()

	patern := "goMeteoMQTT:airport:*"

	var data []interface{}
	data, err := redis.Values(conn.Do("SCAN", 0, "MATCH", patern, "COUNT", "1000000000"))

	var airports []models.Airport
	if err != nil {
		return airports, fmt.Errorf("error scanning for %s: %v", patern, err)
	}

	keys, _ := redis.Strings(data[1], nil)

	for _, id := range keys {
		airport, _ := HGetAllAirports(pool, id)
		airports = append(airports, airport)
	}

	return airports, err

}

func ScanByAirportAndType(pool *redis.Pool, airportId, captorType string) ([]CaptorValue, error) {

	conn := pool.Get()
	defer conn.Close()

	patern := "goMeteoMQTT:captorValues:*" + airportId + "*" + captorType + "*"

	var data []interface{}
	data, err := redis.Values(conn.Do("SCAN", 0, "MATCH", patern, "COUNT", "1000000000"))

	var captorValues []CaptorValue
	if err != nil {
		return captorValues, fmt.Errorf("error scanning for %s: %v", patern, err)
	}

	keys, _ := redis.Strings(data[1], nil)

	for _, id := range keys {
		captorValue, _ := HGetAllCaptors(pool, id)
		captorValues = append(captorValues, captorValue)
	}

	return captorValues, err

}

func ScanByAirportAndTypeAndDate(pool *redis.Pool, airportId, captorType string, dateMin, dateMax time.Time) []CaptorValue {

	captorValues, _ := ScanByAirportAndType(pool, airportId, captorType)

	var result []CaptorValue
	for i := range captorValues {
		date := captorValues[i].Date
		if date.After(dateMin) && date.Before(dateMax) {
			result = append(result, captorValues[i])
		}
	}

	return result

}

// ZRANGEBYSCORE goMeteoMQTT:dateIndex 20201024 20201025     ==> donne la liste de toutes les valeurs entre les dates
// SCAN 0 MATCH *BDX*wind* COUNT 1000						 ==> donne la liste des valeurs de capteur de vents de bordeaux aéroport
