package models

import (
	"fmt"
	"strconv"
	"time"
)

type CaptorValue struct {
	Id        int       `example:"326071054919270689"`
	CaptorId  int       `example:"2"`
	AirportId string    `example:"BDX"`
	Type      string    `example:"wind"`
	Date      time.Time `example:"2006-01-02 03:04:05"`
	Value     float64   `example:"24.05"`
}

func (o CaptorValue) String() string {

	return fmt.Sprintf("{ Id: %d, CaptorId %d, AirportId: %s, Type: %s, Date: %s, Value: %.2f}",
		o.Id, o.CaptorId, o.AirportId, o.Type, o.Date, o.Value)

}

func MakeCaptorValueFromMap(m map[string]interface{}) CaptorValue {

	res := CaptorValue{
		Id:        int(m["Id"].(uint64)),
		CaptorId:  int(m["CaptorId"].(float64)),
		AirportId: m["AirportId"].(string),
		Type:      m["Type"].(string),
		Date:      m["Date"].(time.Time),
		Value:     m["Value"].(float64),
	}
	return res

}

func MakeCaptorValueFromArray(a []string, uniqueId int) CaptorValue {

	layout := "2006-01-02 03:04:05"

	date, _ := time.Parse(layout, a[0])

	airportId := a[1]
	captorType := a[2]
	captorId, _ := strconv.Atoi(a[3])
	value, _ := strconv.ParseFloat(a[4], 64)

	res := CaptorValue{
		Id:        uniqueId,
		CaptorId:  captorId,
		AirportId: airportId,
		Type:      captorType,
		Date:      date,
		Value:     value,
	}
	return res

}

func MakeCaptorValueFromRedisArray(a []string) CaptorValue {

	layout := "2006-01-02 03:04:05 +0000 UTC"

	uniqueId, _ := strconv.Atoi(a[0])
	captorId, _ := strconv.Atoi(a[1])
	airportId := a[2]
	captorType := a[3]
	date, _ := time.Parse(layout, a[4])
	value, _ := strconv.ParseFloat(a[5], 64)

	res := CaptorValue{
		Id:        uniqueId,
		CaptorId:  captorId,
		AirportId: airportId,
		Type:      captorType,
		Date:      date,
		Value:     value,
	}
	return res

}
