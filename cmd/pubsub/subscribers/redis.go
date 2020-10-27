package subscribers

import (
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

func HGet(id, key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("HGET", id, key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s for %s: %v", key, id, err)
	}
	return data, err

}

func HGetAll(pool *redis.Pool, hash string) (CaptorValue, error) {

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

func ScanByAirportAndType(pool *redis.Pool, airportId, captorType string) ([]CaptorValue, error) {

	conn := pool.Get()
	defer conn.Close()

	patern := "*" + airportId + "*" + captorType + "*"

	var data []interface{}
	data, err := redis.Values(conn.Do("SCAN", 0, "MATCH", patern, "COUNT", "1000000000"))

	var captorValues []CaptorValue
	if err != nil {
		return captorValues, fmt.Errorf("error scanning for %s: %v", patern, err)
	}

	keys, _ := redis.Strings(data[1], nil)

	for _, id := range keys {
		captorValue, _ := HGetAll(pool, id)
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
