package models

import (
	"fmt"
)

type Airport struct {
	Id   string `example:"NTE"`
	Name string `example:"Nantes Atlantique"`
	City string `example:"Nantes"`
}

func (o Airport) String() string {

	return fmt.Sprintf("{ Id: %s, Name: %s, City: %s }",
		o.Id, o.Name, o.City)

}

func AirportMapper(m map[string]interface{}) Airport {
	res := Airport{
		Id:   m["Id"].(string),
		Name: m["Name"].(string),
		City: m["City"].(string),
	}
	return res
}
