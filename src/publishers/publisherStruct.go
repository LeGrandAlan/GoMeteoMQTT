package publishers

import (
	"fmt"
)

type Publisher struct {
	Name  string
	Type  string
	Unit  string
	Min   float64
	Max   float64
	Topic string
}

func (o Publisher) String() string {
	return fmt.Sprintf("{ Name: %s, Type: %s, Unit: %s, Min: %.2f, Max: %.2f, Topic: %s }",
		o.Name, o.Type, o.Unit, o.Min, o.Max, o.Topic)
}

func makeFromMap(m map[string]interface{}) Publisher {
	res := Publisher{
		Name:  m["Name"].(string),
		Type:  m["Type"].(string),
		Unit:  m["Unit"].(string),
		Min:   m["Min"].(float64),
		Max:   m["Max"].(float64),
		Topic: m["Topic"].(string),
	}
	return res
}
