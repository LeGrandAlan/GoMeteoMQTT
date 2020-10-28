package models

import (
	"fmt"
)

type Publisher struct {
	Id        int
	AirportId string
	Type      string
	Unit      string
	Min       float64
	Max       float64
}

func (o Publisher) String() string {

	return fmt.Sprintf("{ Id: %d, AirportId: %s, Type: %s, Unit: %s, Min: %.2f, Max: %.2f }",
		o.Id, o.AirportId, o.Type, o.Unit, o.Min, o.Max)

}

func MakePublisherFromMap(m map[string]interface{}) Publisher {

	res := Publisher{
		Id:        int(m["Id"].(float64)),
		AirportId: m["AirportId"].(string),
		Type:      m["Type"].(string),
		Unit:      m["Unit"].(string),
		Min:       m["Min"].(float64),
		Max:       m["Max"].(float64),
	}
	return res

}
