package publishers

import (
	"fmt"
)

type Publisher struct {
	Id        string
	AirportId string
	Name      string
	Type      string
	Unit      string
	Min       float64
	Max       float64
	Topic     string
}

func (o Publisher) String() string {
	return fmt.Sprintf("{ Id: %s, AirportId: %s, Name: %s, Type: %s, Unit: %s, Min: %.2f, Max: %.2f, Topic: %s }",
		o.Id, o.AirportId, o.Name, o.Type, o.Unit, o.Min, o.Max, o.Topic)
}

func MakeFromMap(m map[string]interface{}) Publisher {
	res := Publisher{
		Id:        m["Id"].(string),
		AirportId: m["AirportId"].(string),
		Name:      m["Name"].(string),
		Type:      m["Type"].(string),
		Unit:      m["Unit"].(string),
		Min:       m["Min"].(float64),
		Max:       m["Max"].(float64),
		Topic:     m["Topic"].(string),
	}
	return res
}
